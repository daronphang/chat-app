package domain

import "context"

type Repository interface {
	GetLatestMessages(ctx context.Context, channelID string) ([]Message, error)
	GetPreviousMessages(ctx context.Context, arg PrevMessageRequest) ([]Message, error)
	CreateMessage(ctx context.Context, arg Message) error
	UpdateMessageStatus(ctx context.Context, arg Message) error
}