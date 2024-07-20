package usecase

import (
	"chat-service/internal/domain"
	"context"
	"time"

	snowflake "github.com/godruoyi/go-snowflake"
)

func (uc *UseCaseService) AckSenderMsg(ctx context.Context, arg domain.Message) (domain.Message, error) {
	arg.MessageID = snowflake.ID()
	arg.CreatedAt = time.Now().String()
	err := uc.EventBroker.PublishMessage(ctx, arg.ChannelID, domain.MessageTopicConfig.Topic, arg)
	if err != nil {
		return domain.Message{}, err
	}
	return arg, nil
}

func (uc *UseCaseService) SendMsgToClientDevices(ctx context.Context, clientId string, arg domain.Message) error {
	if err := uc.ServerClienter.SendMsgToClientDevices(ctx, clientId, arg); err != nil {
		return err
	}
	return nil
}