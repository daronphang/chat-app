package domain

import (
	"context"
)

type Repository interface {
	UpdateUserSession(ctx context.Context, arg HeartbeatRequest) error
	GetUserSession(ctx context.Context, userID string) (UserSession, error)
}