package usecase

import (
	"context"
	"message-service/internal/domain"
)

type EventBroker interface {
	PublishMessage(ctx context.Context, partitionKey string, topic string, arg interface{}) error
}

type MessageBroker interface {
	PublishMessage(ctx context.Context, queue string, routingKey string, arg interface{}) error
}

type UseCaseService struct {
	Repository 		domain.Repository
	MessageBroker 	MessageBroker
	EventBroker		EventBroker
}

func NewUseCaseService(mb MessageBroker, eb EventBroker, repo domain.Repository) *UseCaseService {
	return &UseCaseService{MessageBroker: mb, EventBroker: eb, Repository: repo}
}
