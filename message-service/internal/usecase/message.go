package usecase

import (
	"context"
	"message-service/internal/domain"

	"github.com/google/uuid"
)

func (uc *UseCaseService) GetMessages(ctx context.Context, channelID string) error {
	return nil
}

func (uc *UseCaseService) SaveMessageAndRoute(ctx context.Context, arg domain.Message) error {
	// Save message in Cassandra.

	// Get list of receiver users to send message to.

	// If user is online, send message to respective chat server.
	// If user is offine, to send push notification via queue.
	return nil
}

func (uc *UseCaseService) AddUsersToChat(ctx context.Context, channelID string, users []string) {
	if channelID == "" {
		channelID = uuid.NewString()

		if len(users) == 2 {
			// Check if a channel has already been created by the other user.
			
		}
	}
}

func (uc *UseCaseService) JoinGroup(ctx context.Context) {}

func (uc *UseCaseService) LeaveGroup(ctx context.Context) {}

func (uc *UseCaseService) DeleteUserChat(ctx context.Context, client string, channelID string) {}