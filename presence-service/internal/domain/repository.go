package domain

import (
	"context"
)

type Repository interface {
	UpdateUser(ctx context.Context, arg HeartbeatRequest) error
	GetUser(ctx context.Context, userID string) (User, error)
}