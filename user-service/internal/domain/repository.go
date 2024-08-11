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
	AddFriend(ctx context.Context, arg NewFriend) error
	GetFriends(ctx context.Context, arg string) ([]Friend, error)
	CreateUserToChannelAssociation(ctx context.Context, arg NewChannel) error
	CreateGroupChannel(ctx context.Context, arg NewChannel) error
	GetUsersAssociatedToChannel(ctx context.Context, arg string) ([]UserContact, error)
	GetChannelsAssociatedToUser(ctx context.Context, arg string) ([]Channel, error)
	GetUsersAssociatedToTargetUser(ctx context.Context, arg string) ([]string, error)
}