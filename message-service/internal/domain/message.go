package domain

type MessageStatus string

const (
	Received 		MessageStatus 	= "received"
	Delivered 		MessageStatus 	= "delivered"
	Read 			MessageStatus 	= "read"
)

type Message struct {
	MessageID 			uint64 			`json:"messageId"`
	ChannelID 			string 			`json:"channelId" validate:"required"` 
	SenderID 			string 			`json:"senderId" validate:"required"`
	MessageType 		string 			`json:"messageType" validate:"required"`
	Content 			string 			`json:"content" validate:"required"`
	CreatedAt 			string 			`json:"createdAt"`
	MessageStatus		MessageStatus	`json:"messageStatus"` 
}

type UserChannelRequest struct {
	ChannelID 	string 		`validate:"required"`
	UserIDs		[]string 	`validate:"required"`
}

type PrevMessageRequest struct {
	ChannelID 		string	`json:"channelId" validate:"required"` 
	LastMessageID 	string	`json:"lastMessageId" validate:"required"` 
}
