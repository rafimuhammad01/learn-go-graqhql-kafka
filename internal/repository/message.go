package repository

import (
	"context"
	"github.com/rafimuhammad01/learn-go-graphql/internal/core"
)

type Broker interface {
	Consume(ctx context.Context) (*core.Message, *core.Error)
	Produce(ctx context.Context, msg interface{}) *core.Error
}

type message struct {
	broker Broker
}

func (m *message) ListenMessage(ctx context.Context) (*core.Message, *core.Error) {
	msg, err := m.broker.Consume(ctx)
	if err != nil {
		return nil, err
	}

	return msg, err
}

func (m *message) SendMessage(ctx context.Context, message core.Message) *core.Error {
	err := m.broker.Produce(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

func NewMessage(broker Broker) *message {
	return &message{
		broker: broker,
	}
}
