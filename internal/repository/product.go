package repository

import (
	"context"
	"github.com/rafimuhammad01/learn-go-graphql/internal/core"
)

type product struct {
	products []*core.Product
}

func (p *product) List(ctx context.Context) ([]*core.Product, *core.Error) {
	return p.products, nil
}

func (p *product) Create(ctx context.Context, product core.Product) *core.Error {
	p.products = append(p.products, &product)
	return nil
}

func NewProduct() *product {
	return &product{}
}
