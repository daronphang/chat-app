package kafka

import (
	"broadcast-service/internal/config"

	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	writer *kafka.Writer
}

func New(cfg *config.Config) *KafkaClient {
	return &KafkaClient{writer: newWriter(cfg)}
}

func (kc *KafkaClient) Close() {
	kc.writer.Close()
}