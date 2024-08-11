package usecase

import (
	"context"
	"notification-service/internal/domain"
	"sync"

	"google.golang.org/protobuf/types/known/wrapperspb"
)


func (uc * UseCaseService) BroadcastMessageEvent(ctx context.Context, arg domain.Message) error {
	// Get all users associated to channel.
	users, err := uc.UserClient.GetUsersAssociatedToChannel(ctx, &wrapperspb.StringValue{Value: arg.ChannelID})
	if err != nil {
		return err
	}

	userIds := make([]string, 0)
	userContacts := make([]domain.UserContact, 0)
	for _, user := range users.UserContacts {
		userIds = append(userIds, user.UserId)
		userContacts = append(
			userContacts, 
			domain.UserContact{
				UserID: user.UserId,
				Email: user.Email,
			},
		)
	}

	// Get offline users.
	offlineUsers := uc.GetOfflineUsers(ctx, userIds)

	// Instead of sending messages directly to the sender's devices, to push messages 
	// as events into the recipients' queues. This ensures events will never be lost.
	event := domain.BaseEvent{
		Event: domain.EventMessage,
		Data: arg,
	}
	maxGoroutines := 10
	guard := make(chan bool, maxGoroutines)
	wg := sync.WaitGroup{}
	for _, userContact := range userContacts {
		guard <- true
		wg.Add(1)
		go func(u domain.UserContact) {
			defer wg.Done()
			isOffline := false
			for _, userSession := range offlineUsers {
				if userSession.UserID == u.UserID {
					isOffline = true
					break
				}
			}
			uc.HandleEventRoutingByUserStatus(ctx, u, event, isOffline)
			<- guard
		}(userContact)
	}
	wg.Wait()
	return nil
}