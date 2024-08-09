package usecase

import (
	"chat-service/internal/domain"
	"context"
)

func (uc *UseCaseService) SendEventToClient(ctx context.Context, clientId string, event domain.Event, payload interface{}) error {
	if err := uc.ServerClienter.SendEventToClient(ctx, clientId, event, payload); err != nil {
		return err
	}
	return nil
}