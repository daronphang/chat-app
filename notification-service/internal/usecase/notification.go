package usecase

import (
	"context"
	"encoding/json"
	"notification-service/internal/domain"
	"notification-service/internal/util"
	"time"

	"go.uber.org/zap"
)

// Pre-determined logic for notifying events to user if he/she is online/offline.
// Regardless of whether the status is online/offline, to push event to the user's queue.
// If user is offline, to send push notification of the event via message queue.
//
// To be executed in a goroutine.
func (uc * UseCaseService) HandleEventRoutingByUserStatus(ctx context.Context, user domain.UserContact, event domain.BaseEvent, isOffline bool) {
	// Push event to user's queue, regardless of online/offline status.
	err := util.ExpBackoff(
		500 * time.Millisecond,
		2,
		3,
		func() error {
			return uc.EventBroker.PublishEventToUserQueue(ctx, user.UserID, event)
		},
	)

	b, _ := json.Marshal(event)
	if err != nil {
		logger.Error(
			"failed to push event to user queue",
			zap.String("payload", string(b)),
			zap.String("trace", err.Error()),
		)
		return
	}

	// If user is offline, to send push notification via message queue.
	// Delivery is not guaranteed.
	if isOffline {
		if err := uc.MessageBroker.PublishPushNotificationEvent(ctx, event); err != nil {
			logger.Error(
				"failed to notify user",
				zap.String("event", string(b)),
				zap.String("trace", err.Error()),
			)
		}
	}
}