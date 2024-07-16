package kafka

import (
	"chat-service/internal/config"
	"chat-service/internal/domain"
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	Writer *kafka.Writer
}

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
func New(cfg *config.Config) Kafka {
	w := &kafka.Writer{
		Addr: kafka.TCP(strings.Split(cfg.Kafka.BrokerAddresses, ",")...),
		RequiredAcks: kafka.RequireOne,
		MaxAttempts: 5,
		BatchSize: 100,
		BatchTimeout: 1 * time.Second,
	}
	return Kafka{Writer: w}
}

func (k Kafka) PublishMessage(ctx context.Context, arg domain.Message) error {
	v, _ := json.Marshal(arg)
	msg := kafka.Message{
		Key: []byte(arg.ChannelID),
		Value: v,
		Topic: Messages.String(),
	}
	if err := k.Writer.WriteMessages(ctx, msg); err != nil {
		return err
	}
	return nil
}