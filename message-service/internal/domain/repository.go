package domain

import "context"

type Repository interface {
	GetUnreadMessages(ctx context.Context, arg MessageRequest) ([]Message, error)
	GetPreviousMessages(ctx context.Context, arg MessageRequest) ([]Message, error)
	CreateMessage(ctx context.Context, arg Message) error
	UpdateMessageStatus(ctx context.Context, arg Message) error
}