package usecase

import (
	"context"
	"message-service/internal/domain"
)

func (uc *UseCaseService) GetMessagesForChannel(ctx context.Context, channelID string) error {
	return nil
}

func (uc *UseCaseService) SaveMessageAndRoute(ctx context.Context, arg domain.Message) error {
	// Save message in Cassandra.

	// Get list of receiver users to send message to.

	// If user is online, send message to respective chat server.
	// If user is offine, to send push notification via queue.
	return nil
}