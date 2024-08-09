package kafka

import (
	"chat-service/internal/domain"
	"chat-service/internal/usecase"
	cv "chat-service/internal/validator"
	"context"
	"errors"
	"io"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type KafkaConsumerConfig struct {
	Brokers 		[]string
	ConsumerGroupID string
	Topic 			string
}
type KafkaConsumer struct {
	reader *kafka.Reader
	client string 
}

// Consumer group ID for each user's topic should not be changed
// as this would affect what messages would be pushed to the client upon
// connecting with the chat server.
//
// Important to call Close() on a Reader when a process exits
// as Kafka server needs a graceful disconnect to stop it from
// continuing to attempt to send messages on connected clients.
func NewConsumer(cfg KafkaConsumerConfig) *KafkaConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Brokers,
		// Consumers in the same consumer group will always read a unique partition.
		// Different consumer groups can read from the same partition.
		GroupID: cfg.ConsumerGroupID, 
		Topic: cfg.Topic,
		MaxWait: 1 * time.Second,
	})
	return &KafkaConsumer{reader: r, client: cfg.Topic}
}

func (c *KafkaConsumer) Close() {
	if err := c.reader.Close(); err != nil {
		logger.Error(
			"error closing kafka reader",
			zap.String("trace", err.Error()),
		)
	}
}

// To execute the consumer function as a goroutine.
// One consumer per thread/goroutine is the rule.
// Creating more consumers than the number of partitions will result in unused consumers.

func (c *KafkaConsumer) ConsumeFromUserTopic(ctx context.Context, uc *usecase.UseCaseService) {
	for {
		// To consume both messages and channel events.
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			logger.Error(
				"error reading from messages topic",
				zap.String("trace", err.Error()),
			)
			continue
		}

		// Validate dynamic struct.
		msg := new(domain.Message)
		newChEvent := new(domain.NewChannelEvent)
		var v interface{}
		var event domain.Event

		if err := cv.UnmarshalAndValidate(m.Value, msg); err == nil {
			v = *msg
			event = domain.NewMessage
		} else if err := cv.UnmarshalAndValidate(m.Value, newChEvent); err == nil {
			v = *newChEvent
			event = domain.NewChannel
		} else {
			logger.Error(
				"error validating message from user topic",
				zap.String("payload", string(m.Value)),
				zap.String("trace", err.Error()),
			)
			continue
		}
	
		if err := uc.SendEventToClient(ctx, c.client, event, v); err != nil {
			logger.Error(
				"error sending inbound data to client",
				zap.String("payload", string(m.Value)),
				zap.String("trace", err.Error()),
			)
		}
	}
}