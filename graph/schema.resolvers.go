package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/rafimuhammad01/learn-go-graphql/graph/generated"
	"github.com/rafimuhammad01/learn-go-graphql/graph/model"
	"github.com/rafimuhammad01/learn-go-graphql/internal/core"
)

// CreateProduct is the resolver for the CreateProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, product *model.CreateProductParams) (*model.Response, error) {
	productEntity := core.Product{
		Name:        SafeDeref(product.Name),
		Description: SafeDeref(product.Description),
		Price:       SafeDeref(product.Price),
	}

	err := r.ProductService.Create(ctx, productEntity)
	if err != nil {
		return nil, err
	}

	msg := "success"
	return &model.Response{Message: &msg}, nil
}

// SendMessage is the resolver for the SendMessage field.
func (r *mutationResolver) SendMessage(ctx context.Context, msg model.MessageInput) (*model.Response, error) {
	msgWrapped := core.Message{
		ID: msg.ID,
		From: core.User{
			ID:   msg.From.ID,
			Name: msg.From.Name,
		},
		To: core.User{
			ID:   msg.To.ID,
			Name: msg.To.Name,
		},
		Msg: msg.Msg,
	}

	err := r.MessageService.SendMessage(ctx, msgWrapped)
	if err != nil {
		return nil, err
	}

	msgResp := "success"
	return &model.Response{Message: &msgResp}, nil
}

// ListProduct is the resolver for the ListProduct field.
func (r *queryResolver) ListProduct(ctx context.Context) ([]*model.Product, error) {
	//panic(fmt.Errorf("not implemented: ListProduct - ListProduct"))
	products, err := r.ProductService.List(ctx)
	if err != nil {
		return nil, err
	}

	var resp []*model.Product
	for _, v := range products {
		resp = append(resp, &model.Product{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
		})
	}

	return resp, nil
}

// GetProduct is the resolver for the GetProduct field.
func (r *queryResolver) GetProduct(ctx context.Context, id string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: GetProduct - GetProduct"))
}

// ListenMessage is the resolver for the ListenMessage field.
func (r *subscriptionResolver) ListenMessage(ctx context.Context) (<-chan *model.Message, error) {
	// Create new channel for request
	messages := make(chan *model.Message, 1)
	//var errChan chan error

	go func() {
		for {
			msg, _ := r.MessageService.ListenMessage(ctx)
			msgParse := model.Message{
				ID: msg.ID,
				From: &model.User{
					ID:   msg.From.ID,
					Name: msg.From.Name,
				},
				To: &model.User{
					ID:   msg.To.ID,
					Name: msg.To.Name,
				},
				Msg: msg.Msg,
			}
			messages <- &msgParse
		}
	}()

	return messages, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
