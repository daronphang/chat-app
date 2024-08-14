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
	CreateUserToChannelAssociation(ctx context.Context, arg Channel) error
	CreateGroupChannel(ctx context.Context, arg Channel) error
	GetUsersAssociatedToChannel(ctx context.Context, arg string) ([]UserContact, error)
	GetChannelsAssociatedToUser(ctx context.Context, arg string) ([]Channel, error)
	GetUsersAssociatedToTargetUser(ctx context.Context, arg string) ([]string, error)
	GetUsersContactsMetadata(ctx context.Context, arg []string) ([]UserContact, error)
	GetGroupChannel(ctx context.Context, arg string) (Channel, error)
	RemoveGroupMembers(ctx context.Context, arg GroupMembers) error
	RemoveGroup(ctx context.Context, arg string) error
}