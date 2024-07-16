package domain

type Message struct {
	MessageID uint64 `json:"messageId"`
	PreviousMessageID uint64 `json:"previousMessageId" validate:"required"`
	ChannelID string `json:"channelId" validate:"required"` 
	SenderID string `json:"senderId" validate:"required"`
	Type string `json:"type" validate:"required"`
	Content string `json:"content" validate:"required"`
	CreatedAt string `json:"createdAt" validate:"required"`
}

type ReceiverMessage struct {
	ReceiverID string `json:"receiverId" validate:"required"`
	Message Message `json:"message" validate:"required"`
}
