package domain

var (
	MessageTopicConfig = BrokerTopicConfig{
		Topic: 				"message",
		Partitions: 		10,
		ReplicationFactor: 	1,
		ConsumerGroupID: 	"message1",
	}
	UserTopicConfig = BrokerTopicConfig{
		Topic: 				"", // UserId.
		Partitions: 		1,
		ReplicationFactor: 	1,
		ConsumerGroupID: 	"", // Machine id of host.
	}
)

/*
Topics are explicitly configured for the following reasons:

- You cannot decrease the number of partitions 

- Increasing the partitions will force a re-balance

- ReplicationFactor cannot be greater than the number of brokers available

- Having different consumer groups will read from the same partition and result in duplication
*/
type BrokerTopicConfig struct {
	Topic 				string
	Partitions 			int
	ReplicationFactor 	int
	ConsumerGroupID 	string
}
type Message struct {
	MessageID 			uint64 `json:"messageId"`
	ChannelID 			string `json:"channelId" validate:"required"` 
	SenderID 			string `json:"senderId" validate:"required"`
	MessageType 		string `json:"messageType" validate:"required"`
	Content 			string `json:"content" validate:"required"`
	CreatedAt 			string `json:"createdAt"`
	TempID				int64  `json:"tempId,omitempty"` // For client (sender) to be aware of msg ack.
}

type ReceiverMessage struct {
	ReceiverID 	string 	`json:"receiverId" validate:"required"`
	Message 	Message `json:"message" validate:"required"`
}
