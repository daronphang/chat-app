package kafka

import (
	"context"
	"errors"
	"io"
	"time"

	"message-service/internal/domain"
	"message-service/internal/usecase"
	cv "message-service/internal/validator"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

// Important to call Close() on a Reader when a process exits
// as Kafka server needs a graceful disconnect to stop it from
// continuing to attempt to send messages on connected clients.
// This can result in a delay when a new reader on the same topic connects.
func NewConsumer(brokers []string, consumerGroupID string, topic string) *KafkaConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		// Consumers in the same consumer group will always read a unique partition.
		// Different consumer groups can read from the same partition.
		GroupID: consumerGroupID, 
		Topic: topic,
		// Workaround for bug where reader hangs on process exit.
		MaxWait: 1 * time.Second, 
	})
	return &KafkaConsumer{reader: r}
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
func (c *KafkaConsumer) ConsumeFromMessageTopic(ctx context.Context, uc *usecase.UseCaseService) {
	for {
		m, err := c.reader.FetchMessage(ctx)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			logger.Error(
				"error reading message from messages topic",
				zap.String("trace", err.Error()),
			)
			continue
		}
	
		msg := new(domain.Message)
		if err := cv.UnmarshalAndValidate(m.Value, msg); err != nil {
			logger.Error(
				"error unmarshaling Kafka message into JSON",
				zap.String("payload", string(m.Value)),
				zap.String("trace", err.Error()),
			)
			continue
		}
	
		if err := uc.SaveMessageAndNotifyRecipients(ctx, *msg); err != nil {
			logger.Error(
				"error saving and delivering message",
				zap.String("payload", string(m.Value)),
				zap.String("trace", err.Error()),
			)
			continue
		}

		if err := c.reader.CommitMessages(ctx, m); err != nil {
			logger.Error(
				"failed to commit message",
				zap.String("trace", err.Error()),
			)
		} else {
			logger.Info(
				"message consumption success",
				zap.String("payload", string(m.Value)),
			)
		}
	}
}