package domain

const (
	Received = iota + 1
	Delivered 
)

type Message struct {
	MessageID 			uint64 			`json:"messageId"`
	ChannelID 			string 			`json:"channelId" validate:"required"` 
	SenderID 			string 			`json:"senderId" validate:"required"`
	MessageType 		string 			`json:"messageType" validate:"required"`
	Content 			string 			`json:"content" validate:"required"`
	CreatedAt 			string 			`json:"createdAt"`
	MessageStatus		int				`json:"messageStatus"` 
}

type UserChannelRequest struct {
	ChannelID 	string 		`validate:"required"`
	UserIDs		[]string 	`validate:"required"`
}

type MessageRequest struct {
	ChannelID 		string	`json:"channelId" validate:"required"` 
	LastMessageID 	*uint64	`json:"lastMessageId" validate:"required"` 
}