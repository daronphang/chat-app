package usecase

import (
	"context"
	"fmt"
	"message-service/internal"
	"message-service/internal/domain"
	"protobuf/proto/common"
	"slices"
	"time"

	"go.uber.org/zap"
)

var (
	logger, _ = internal.WireLogger()
)

func (uc *UseCaseService) GetLatestMessages(ctx context.Context, channelID string) ([]domain.Message, error) {
	// Very rarely users fetch old messages.
	// to return messages in ascending order.
	rv, err := uc.Repository.GetLatestMessages(ctx, channelID)
	if err != nil {
		return nil, err
	}
	// Messages are returned in descending order, to reverse.
	slices.Reverse(rv)
	return rv, nil
}

func (uc *UseCaseService) GetPreviousMessages(ctx context.Context, arg domain.PrevMessageRequest) ([]domain.Message, error) {
	rv, err := uc.Repository.GetPreviousMessages(ctx, arg)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *UseCaseService) UpdateMessageStatus(ctx context.Context, arg domain.Message) error {
	if err := uc.Repository.UpdateMessageStatus(ctx, arg); err != nil {
		return err
	}
	return nil
}

func (uc *UseCaseService) SaveMessageAndNotifyRecipients(ctx context.Context, arg domain.Message) error {
	// Save message in db.
	arg.MessageStatus = domain.Received
	if err := uc.Repository.CreateMessage(ctx, arg); err != nil {
		return err
	}

	// Broadcast message event.
	msg := &common.Message{
		MessageId: arg.MessageID,
		ChannelId: arg.ChannelID,
		SenderId: arg.SenderID,
		Content: arg.Content,
		CreatedAt: arg.CreatedAt,
		MessageType: arg.MessageType,
		MessageStatus: int32(arg.MessageStatus),
	}
	_, err := uc.NotificationClient.BroadcastMessageEvent(ctx, msg)
	if err != nil {
		return err
	}

	// Update status of delivered message in db and notify sender.
	arg.MessageStatus = domain.Delivered
	if err := uc.UpdateMessageStatus(ctx, arg); err != nil {
		return err
	}

	event := domain.BaseEvent{
		Event: domain.EventMessage,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data: arg,
	}
	if err := uc.EventBroker.PublishEventToUserQueue(ctx, arg.SenderID, event); err != nil {
		logger.Error(
			fmt.Sprintf("failed to notify delivery to sender %v for message %v", arg.SenderID, arg.MessageID),
			zap.String("trace", err.Error()),
		)
	}

	return nil
}