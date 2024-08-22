package usecase

import (
	"chat-service/internal/domain"
	"context"
)

func (uc *UseCaseService) SendEventToClient(ctx context.Context, clientId string, event domain.BaseEvent) error {
	if err := uc.ServerClienter.SendEventToClient(ctx, clientId, event); err != nil {
		return err
	}
	return nil
}