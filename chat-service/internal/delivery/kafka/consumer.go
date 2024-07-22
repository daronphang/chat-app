package kafka

import (
	"chat-service/internal/config"
	"chat-service/internal/domain"
	"chat-service/internal/usecase"
	cv "chat-service/internal/validator"
	"context"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type KafkaConsumer struct {
	reader *kafka.Reader
	client string 
}

// Consumer group ID for each user's topic should not be changed
// as this would affect what messages would be pushed to the client upon
// connecting with the chat server.

// Important to call Close() on a Reader when a process exits
// as Kafka server needs a graceful disconnect to stop it from
// continuing to attempt to send messages on connected clients.
func NewConsumer(cfg *config.Config, consumerGroupID string, topic string) *KafkaConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: strings.Split(cfg.Kafka.BrokerAddresses, ","),
		// Consumers in the same consumer group will always read a unique partition.
		// Different consumer groups can read from the same partition.
		GroupID: consumerGroupID, 
		Topic: topic,
		MaxWait: 1 * time.Second,
	})
	return &KafkaConsumer{reader: r, client: topic}
}

// To execute the consumer function as a goroutine.
// One consumer per thread/goroutine is the rule.
// Creating more consumers than the number of partitions will result in unused consumers.
func (c *KafkaConsumer) ConsumeMsgFromUserTopic(ctx context.Context, uc *usecase.UseCaseService) {
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			logger.Error(
				"error reading from messages topic",
				zap.String("trace", err.Error()),
			)
			if err := c.reader.Close(); err != nil {
				logger.Error(
					"unable to close kafka reader",
					zap.String("trace", err.Error()),
				)
			}
			return
		}
	
		msg := new(domain.Message)
		if err := cv.UnmarshalAndValidate(m.Value, msg); err != nil {
			logger.Error(
				"error unmarshaling message from topic into JSON",
				zap.String("payload", string(m.Value)),
				zap.String("trace", err.Error()),
			)
			continue
		}
	
		if err := uc.ForwardMsgToClient(ctx, c.client, *msg); err != nil {
			logger.Error(
				"error forwarding message to client",
				zap.String("payload", string(m.Value)),
				zap.String("trace", err.Error()),
			)
		}
	}
}