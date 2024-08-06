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
	UserTopicConfig = BrokerTopicConfig{
		Topic: 				"", // UserId.
		Partitions: 		1,
		ReplicationFactor: 	1,
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
	ChannelID 			string `json:"channelId" validate:"required"` 
	SenderID 			string `json:"senderId" validate:"required"`
	MessageType 		string `json:"messageType" validate:"required"`
	Content 			string `json:"content" validate:"required"`
	CreatedAt 			string `json:"createdAt"`
}

type UserChannelRequest struct {
	ChannelID 	string 		`validate:"required"`
	UserIDs		[]string 	`validate:"required"`
}

type PrevMessageRequest struct {
	ChannelID 		string	`json:"channelId" validate:"required"` 
	LastMessageID 	string	`json:"lastMessageId" validate:"required"` 
}
