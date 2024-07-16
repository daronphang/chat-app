package kafka

import (
	"context"
	"io"
	"strings"
	"time"

	"message-service/internal/config"
	"message-service/internal/domain"
	"message-service/internal/usecase"
	cv "message-service/internal/validator"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Kafka struct {
	Reader *kafka.Reader
}

// Important to call Close() on a Reader when a process exits
// as Kafka server needs a graceful disconnect to stop it from
// continuing to attempt to send messages on connected clients.
func New(cfg *config.Config, consumerGroupID string, topic string) Kafka {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: strings.Split(cfg.Kafka.BrokerAddresses, ","),
		// Consumers in the same consumer group will always read a unique partition.
		// Different consumer groups can read from the same partition.
		GroupID: consumerGroupID, 
		Topic: topic,
		MaxWait: 1 * time.Second,
	})
	return Kafka{Reader: r}
}

// To execute the consumer function as a goroutine.
// One consumer per thread/goroutine is the rule.
// Creating more consumers than the number of partitions will result in unused consumers.
func (k Kafka) ConsumeMsg(ctx context.Context, uc *usecase.UseCaseService) bool {
	m, err := k.Reader.ReadMessage(ctx)
	if err != nil {
		if err != io.EOF {
			logger.Error(
				"error reading message from messages topic",
				zap.String("trace", err.Error()),
			)
		}
		return false
	}

	msg := new(domain.Message)
	if err := cv.UnmarshalAndValidate(m.Value, msg); err != nil {
		logger.Error(
			"error unmarshaling Kafka message into JSON",
			zap.String("payload", string(m.Value)),
			zap.String("trace", err.Error()),
		)
		return true
	}

	if err := uc.SaveMessageAndRoute(ctx, *msg); err != nil {
		logger.Error(
			"error saving message and routing",
			zap.String("payload", string(m.Value)),
			zap.String("trace", err.Error()),
		)
	}
	return true
}