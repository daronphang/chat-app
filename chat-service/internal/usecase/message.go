package usecase

import (
	"chat-service/internal/domain"
	"context"

	snowflake "github.com/godruoyi/go-snowflake"
)

func (uc *UseCaseService) SaveNewMessage(ctx context.Context, arg domain.Message) (domain.Message, error) {
	arg.MessageID = snowflake.ID()
	arg.MessageStatus = domain.Delivered
	// arg.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	err := uc.EventBroker.PublishMessage(ctx, arg.ChannelID, domain.MessageTopic, arg)
	if err != nil {
		return domain.Message{}, err
	}
	return arg, nil
}