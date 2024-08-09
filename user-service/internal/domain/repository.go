package domain

import (
	"context"
)

type ExtRepo interface {
	Repository
	ExecWithTx(ctx context.Context, cb func(Repository) (interface{}, error)) func() (interface{}, error)
}

type Repository interface {
	CreateUser(ctx context.Context, arg NewUser) (UserMetadata, error)
	GetUser(ctx context.Context, arg string) (UserMetadata, error)
	UpdateUser(ctx context.Context, arg UserMetadata) error
	CreateContact(ctx context.Context, arg NewContact) error
	GetContacts(ctx context.Context, arg string) ([]Contact, error)
	CreateUserToChannelAssociation(ctx context.Context, arg NewChannel) error
	CreateGroupChannel(ctx context.Context, arg NewChannel) error
	GetUsersAssociatedToChannel(ctx context.Context, arg string) ([]string, error)
	GetChannelsAssociatedToUser(ctx context.Context, arg string) ([]Channel, error)
	GetUsersAssociatedToTargetUser(ctx context.Context, arg string) ([]string, error)
}