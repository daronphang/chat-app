package domain

type MessageStatus string

const (
	Received 		MessageStatus 	= "received"
	Delivered		MessageStatus 	= "delivered"
	Read		 	MessageStatus 	= "read"
)
type Message struct {
	MessageID 			uint64 			`json:"messageId"`
	ChannelID 			string 			`json:"channelId" validate:"required"` 
	SenderID 			string 			`json:"senderId" validate:"required"`
	MessageType 		string 			`json:"messageType" validate:"required"`
	Content 			string 			`json:"content" validate:"required"`
	MessageStatus		MessageStatus  	`json:"messageStatus"`
	CreatedAt 			string 			`json:"createdAt" validate:"required"`
}