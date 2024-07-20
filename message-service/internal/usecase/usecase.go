package usecase

import (
	"context"
	"message-service/internal/repository"
)

type MessageBroker interface {
	PublishMessage(ctx context.Context, queue string, routingKey string, arg interface{}) error
}

type UseCaseService struct {
	Repository 		repository.Repository
	MessageBroker 	MessageBroker
}

func NewUseCaseService(mb MessageBroker) *UseCaseService {
	return &UseCaseService{MessageBroker: mb}
}
