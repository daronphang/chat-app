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
	MessageStatus		int  			`json:"messageStatus"`
	CreatedAt 			string 			`json:"createdAt" validate:"required"`
}