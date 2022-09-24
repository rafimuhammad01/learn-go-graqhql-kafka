package service

import (
	"context"
	"github.com/rafimuhammad01/learn-go-graphql/internal/core"
)

type MessageRepository interface {
	ListenMessage(ctx context.Context) (*core.Message, *core.Error)
	SendMessage(ctx context.Context, message core.Message) *core.Error
}

type message struct {
	repository MessageRepository
}

func (m *message) ListenMessage(ctx context.Context) (*core.Message, *core.Error) {
	return m.repository.ListenMessage(ctx)
}

func (m *message) SendMessage(ctx context.Context, message core.Message) *core.Error {
	return m.repository.SendMessage(ctx, message)
}

func NewMessage(repository MessageRepository) *message {
	return &message{repository: repository}
}
