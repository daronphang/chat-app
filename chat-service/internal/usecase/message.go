package usecase

import (
	"chat-service/internal/domain"
	"context"
	"time"

	snowflake "github.com/godruoyi/go-snowflake"
)

func (uc *UseCaseService) SendMessage(ctx context.Context, arg domain.Message) (domain.Message, error) {
	arg.MessageID = snowflake.ID()
	arg.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	err := uc.EventBroker.PublishMessage(ctx, arg.ChannelID, domain.MessageTopicConfig.Topic, arg)
	if err != nil {
		return domain.Message{}, err
	}
	return arg, nil
}

func (uc *UseCaseService) ForwardMsgToClient(ctx context.Context, clientId string, arg domain.Message) error {
	if err := uc.ServerClienter.DeliverOutboundMsg(ctx, clientId, arg); err != nil {
		return err
	}
	return nil
}