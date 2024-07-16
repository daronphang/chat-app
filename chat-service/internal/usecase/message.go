package usecase

import (
	"chat-service/internal/domain"
	"context"

	snowflake "github.com/godruoyi/go-snowflake"
)

func (uc *UseCaseService) PushSenderMsgToQueue(ctx context.Context, arg domain.Message) (domain.Message, error) {
	arg.MessageID = snowflake.ID()
	err := uc.EventBroker.PublishMessage(ctx, arg)
	if err != nil {
		return domain.Message{}, err
	}
	return arg, nil
}

func (uc *UseCaseService) SendMsgToClient(ctx context.Context, arg domain.ReceiverMessage) error {
	if err := uc.ServerClienter.SendMsgToClient(ctx, arg); err != nil {
		return err
	}
	return nil
}