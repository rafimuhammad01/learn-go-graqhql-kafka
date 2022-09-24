package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/rafimuhammad01/learn-go-graphql/internal/core"
	"github.com/rafimuhammad01/learn-go-graphql/internal/repository"
	"github.com/rafimuhammad01/learn-go-graphql/internal/service"
	"github.com/rafimuhammad01/learn-go-graphql/pkg/broker"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rafimuhammad01/learn-go-graphql/graph"
	"github.com/rafimuhammad01/learn-go-graphql/graph/generated"
)

const defaultPort = "8080"
const defaultKafkaAddress = "localhost:29092"
const defaultKafkaTopic = "test"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		ProductService: service.NewProduct(repository.NewProduct()),
		MessageService: service.NewMessage(repository.NewMessage(broker.NewKafka(defaultKafkaAddress, defaultKafkaTopic))),
	}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		KeepAlivePingInterval: 1000 * time.Second,
	})

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)

		var myErr *core.Error
		if errors.As(e, &myErr) {
			err.Message = myErr.Cause
			logrus.Errorln(myErr)
		}

		return err
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
