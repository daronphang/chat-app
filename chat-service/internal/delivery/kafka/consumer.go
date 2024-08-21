package kafka

import (
	"chat-service/internal/domain"
	"chat-service/internal/usecase"
	cv "chat-service/internal/validator"
	"context"
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)
type KafkaConsumer struct {
	reader *kafka.Reader
	client string 
}

var (
	errInvalidEvent = errors.New("invalid event")
)

// Consumer group ID for each user's topic should not be changed
// as this would affect what messages would be pushed to the client upon
// connecting with the chat server.
//
// Important to call Close() on a Reader when a process exits
// as Kafka server needs a graceful disconnect to stop it from
// continuing to attempt to send messages on connected clients.
func NewConsumer(brokers []string, consumerGroupID string, topic string) *KafkaConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		// Consumers in the same consumer group will always read a unique partition.
		// Different consumer groups can read from the same partition.
		GroupID: consumerGroupID, 
		Topic: topic,
		MaxWait: 10 * time.Millisecond,
	})
	return &KafkaConsumer{reader: r, client: topic}
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
		// To consume events for messages, channels, etc.
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

		// Unmarshal and validate base event.
		event := new(domain.BaseEvent)
		if err := cv.UnmarshalAndValidate(m.Value, event); err != nil {
			logger.Error(
				"error validating base event from user topic",
				zap.String("payload", string(m.Value)),
				zap.String("trace", err.Error()),
			)
			continue
		}

		// Validate nested data.
		if (event.Event == domain.EventMessage) {
			p, _ := json.Marshal(event.Data)
			v := new(domain.Message)
			err = cv.UnmarshalAndValidate(p, v)
		} else if (event.Event == domain.EventChannel) {
			p, _ := json.Marshal(event.Data)
			v := new(domain.Channel)
			err = cv.UnmarshalAndValidate(p, v)
		} else {
			err = errInvalidEvent
		}

		if err != nil {
			logger.Error(
				"error validating event from user topic",
				zap.String("payload", string(m.Value)),
				zap.String("trace", err.Error()),
			)
			continue
		}
			
		if err := uc.SendEventToClient(ctx, c.client, *event); err != nil {
			logger.Error(
				"error sending inbound data to client",
				zap.String("payload", string(m.Value)),
				zap.String("trace", err.Error()),
			)
		}
	}
}