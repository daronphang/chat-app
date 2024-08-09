package usecase

import (
	"broadcast-service/internal/domain"
	"context"
	"sync"

	"google.golang.org/protobuf/types/known/wrapperspb"
)


func (uc * UseCaseService) BroadcastMessageEvent(ctx context.Context, arg domain.Message) error {
	// Get all users associated to channel.
	users, err := uc.UserClient.GetUsersAssociatedToChannel(ctx, &wrapperspb.StringValue{Value: arg.ChannelID})
	if err != nil {
		return err
	}

	// Get offline users.
	offlineUsers := uc.GetOfflineUsers(ctx, users.UserIds)

	// Instead of sending messages directly to the sender's devices, to push messages 
	// as events into the recipients' queues. This ensures events will never be lost.
	event := domain.BaseEvent{
		Event: domain.EventMessage,
		Data: arg,
	}
	maxGoroutines := 10
	guard := make(chan bool, maxGoroutines)
	wg := sync.WaitGroup{}
	for _, userID := range users.UserIds {
		guard <- true
		wg.Add(1)
		isOffline := false
		for _, userSession := range offlineUsers {
			if userSession.UserID == userID {
				isOffline = true
				break
			}
		}
		go uc.HandleEventRoutingByUserStatus(ctx, userID, event, guard, isOffline)
	}
	wg.Wait()
	return nil
}