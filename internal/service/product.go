package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/rafimuhammad01/learn-go-graphql/internal/core"
)

type ProductRepository interface {
	List(ctx context.Context) ([]*core.Product, *core.Error)
	Create(ctx context.Context, product core.Product) *core.Error
}

func (p *product) List(ctx context.Context) ([]*core.Product, *core.Error) {
	data, err := p.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *product) Create(ctx context.Context, product core.Product) *core.Error {
	product.ID = uuid.New().String()

	err := product.Validate()
	if err != nil {
		return err.WrapTrace("validate error")
	}

	err = p.repository.Create(ctx, product)
	if err != nil {
		return err.WrapTrace("create error")
	}

	return nil
}

type product struct {
	repository ProductRepository
}

func NewProduct(repository ProductRepository) *product {
	return &product{repository: repository}
}
