package usecase

import (
	"context"
	"message-service/internal/domain"

	pb "protobuf/proto/user"
)

type EventBroker interface {
	PublishEventToUserQueue(ctx context.Context, userID string, arg domain.BaseEvent) error
}

type UseCaseService struct {
	Repository 		domain.Repository
	EventBroker		EventBroker
	UserClient   	pb.UserClient
}

func NewUseCaseService(eb EventBroker, repo domain.Repository, uc pb.UserClient) *UseCaseService {
	return &UseCaseService{EventBroker: eb, Repository: repo, UserClient: uc}
}
