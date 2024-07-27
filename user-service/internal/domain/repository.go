package domain

import (
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, arg NewUser) (UserMetadata, error)
	UpdateUser(ctx context.Context, arg UserMetadata) error
}