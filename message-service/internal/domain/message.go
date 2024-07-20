package domain

var (
	MessageTopicConfig = BrokerTopicConfig{
		Topic: 				"message",
		Partitions: 		10,
		ReplicationFactor: 	1,
		ConsumerGroupID: 	"message1",
	}
	NotificationQueueConfig = BrokerQueueConfig{
		Queue: 			"notificationExchange",
		RoutingKeys: 	[]string{"email"},
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

type BrokerQueueConfig struct {
	Queue 		string
	RoutingKeys []string
}

type Message struct {
	MessageID 			uint64 `json:"messageId"`
	PreviousMessageID 	uint64 `json:"previousMessageId" validate:"required"`
	ChannelID 			string `json:"channelId" validate:"required"` 
	SenderID 			string `json:"senderId" validate:"required"`
	Type 				string `json:"type" validate:"required"`
	Content 			string `json:"content" validate:"required"`
	CreatedAt 			string `json:"createdAt" validate:"required"`
}

type ReceiverMessage struct {
	ReceiverID 	string 	`json:"receiverId" validate:"required"`
	Message 	Message `json:"message" validate:"required"`
}
