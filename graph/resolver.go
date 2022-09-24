package graph

import (
	"context"
	"github.com/rafimuhammad01/learn-go-graphql/internal/core"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductService ProductService
	MessageService MessageService
}

type ProductService interface {
	List(ctx context.Context) ([]*core.Product, *core.Error)
	Create(ctx context.Context, product core.Product) *core.Error
}

type MessageService interface {
	ListenMessage(ctx context.Context) (*core.Message, *core.Error)
	SendMessage(ctx context.Context, message core.Message) *core.Error
}
