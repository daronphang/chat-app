package kafka

import (
	"chat-service/internal"
	"chat-service/internal/config"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var (
	logger, _ = internal.WireLogger()
)

type KafkaClient struct {
	writer *kafka.Writer
}

func New(cfg *config.Config) *KafkaClient {
	return &KafkaClient{writer: newWriter(cfg)}
}

func (kc *KafkaClient) Close() {
	if err := kc.writer.Close(); err != nil {
		logger.Error(
			"failed to close Kafka writer",
			zap.String("trace", err.Error()),
		)
	}
}