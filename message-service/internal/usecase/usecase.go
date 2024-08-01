package usecase

import (
	"context"
	"message-service/internal/domain"

	pb "protobuf/user"
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
	UserClient   	pb.UserClient
}

func NewUseCaseService(mb MessageBroker, eb EventBroker, repo domain.Repository, uc pb.UserClient) *UseCaseService {
	return &UseCaseService{MessageBroker: mb, EventBroker: eb, Repository: repo, UserClient: uc}
}
