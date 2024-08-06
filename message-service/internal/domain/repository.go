package domain

import "context"

type Repository interface {
	GetLatestMessages(ctx context.Context, arg string) ([]Message, error)
	GetPreviousMessages(ctx context.Context, arg PrevMessageRequest) ([]Message, error)
	CreateMessage(ctx context.Context, arg Message) error
	// GetUserIdsAssociatedToChannels(ctx context.Context, channelID string) ([]string, error) 
	// GetChannelsAssociatedToUserID(ctx context.Context, userID string) ([]string, error) 
	// AddUserIDsToChannel(ctx context.Context, channelID string, userIDs []string) error
	// GetUserRelations(ctx context.Context, userID string) ([]string, error)
}