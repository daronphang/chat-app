package repository

import (
	"context"
	"message-service/internal/domain"

	"github.com/gocql/gocql"
)

type Repository interface {
	GetMessages(ctx context.Context, channelID string) ([]domain.Message, error)
	CreateMessage(ctx context.Context, arg domain.Message) error
	AddUsersToChannel(ctx context.Context, channelID string, users []string) error
	GetUsersByChannel(ctx context.Context, channelID string) ([]string, error)
	GetChannelsByUser(ctx context.Context, userId string) ([]string, error)
	GetChannelOfPairUsers(ctx context.Context, userId string, user2Id string) (string, error)
}

type Querier struct {
	session *gocql.Session
}

func New() *Querier {
	return &Querier{}
}