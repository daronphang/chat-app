package domain

type Event string

var (
	NewChannel Event = "event/newChannel"
	NewMessage Event = "event/Message"
	UserPresence Event = "event/presence"
)

type WebSocketEvent struct {
	Event Event 
	Data interface{}
}

type NewChannelEvent struct {
	ClientID 		string 		`json:"targetClientId" validate:"required"`
	ChannelID 		string		`json:"channelId" validate:"required"` 
	ChannelName 	string		`json:"channelName" validate:"required"`
	CreatedAt 		string 		`json:"createdAt" validate:"required"`
}

type ChannelUserEvent struct {
	ClientID 		string 	`json:"targetClientId" validate:"required"`
	ChannelID 		string	`json:"channelId" validate:"required"`
	UserID 			string	`json:"userId" validate:"required"`
	Action 			string	`json:"action" validate:"required"`
	Email 			string	`json:"email" validate:"required"`
	DisplayName 	string	`json:"displayName" validate:"required"`
}

type UserPresenceEvent struct {
	TargetID	string 	`json:"targetId" validate:"required"` // UserID to notify.
	ClientID 	string 	`json:"clientId" validate:"required"`
	Status		string	`json:"status" validate:"required,oneof=online offline"` 
}