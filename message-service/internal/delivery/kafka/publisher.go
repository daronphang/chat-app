package kafka

import (
	"context"
	"encoding/json"
	"message-service/internal/config"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

/*
Methods of Writer are safe to use concurrently from multiple goroutines.
However the writer configuration should not be modified after first use.

When sending synchronously and the writer's batch size is configured to
be greater than 1, this method blocks until either a full batch can be
assembled or the batch timeout is reached. The best way to achieve good
batching behavior is to share one Writer amongst multiple go routines.

One writer is maintained throughout the application, and you have the
the ability to define the topic on a per-message basis by setting Message.Topic.
*/
func NewWriter(cfg *config.Config) *kafka.Writer {
	w := &kafka.Writer{
		Addr: kafka.TCP(strings.Split(cfg.Kafka.BrokerAddresses, ",")...),
		RequiredAcks: kafka.RequireOne,
		MaxAttempts: 5,
		BatchSize: 100,
		BatchTimeout: 1 * time.Second,
	}
	return w
}

func (k *KafkaClient) PublishMessage(ctx context.Context, key string, topic string, arg interface{}) error {
	v, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Key: []byte(key),
		Value: v,
		Topic: topic,
	}
	if err := k.Writer.WriteMessages(ctx, msg); err != nil {
		return err
	}
	return nil
}