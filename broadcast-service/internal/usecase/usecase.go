package usecase

import (
	"broadcast-service/internal/domain"
	"context"
	pb "protobuf/proto/user"
)

type UseCaseService struct {
	Repository 		domain.Repository
	UserClient   	pb.UserClient
	MessageBroker 	MessageBroker
	EventBroker		EventBroker
}

type EventBroker interface {
	PublishEventToUserQueue(ctx context.Context, userID string, arg domain.BaseEvent) error
}

type MessageBroker interface {
	PublishNotificationEvent(ctx context.Context, arg domain.BaseEvent) error
}


func NewUseCaseService(repo domain.Repository, u pb.UserClient, mb MessageBroker, eb EventBroker) *UseCaseService {
	return &UseCaseService{Repository: repo, UserClient: u, MessageBroker: mb, EventBroker: eb}
}
