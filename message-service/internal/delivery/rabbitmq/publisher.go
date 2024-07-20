package rmq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// UnsafePush will push to the queue without checking for
// confirmation. It returns an error if it fails to connect.
// No guarantees are provided for whether the server will
// receive the message.
func (client *RabbitMQClient) unsafePush(ctx context.Context, exchange string, routingKey string, body []byte) error {
	if err := client.channel.PublishWithContext(
		ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing {
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body: body,
		},
	); err != nil {
		return err
	}
	return nil
}

// Push will push data onto the queue, and wait for a confirmation.
// This will block until the server sends a confirmation with exponential backoff. 
// Errors are only returned if the push action itself fails.
func (client *RabbitMQClient) PublishMessage(ctx context.Context, queue string, routingKey string, arg interface{}) error {
	client.m.Lock()
	if !client.isReady {
		client.m.Unlock()
		return PublishError{message: "failed to push: disconnected from rabbitmq"}
	}
	client.m.Unlock()

	// Convert msg to byte array.
	v, err := json.Marshal(arg)
	if err != nil {
		return err
	}

	delay := resendDelay
	errCount := 0
	negAckCount := 0
	maxRetries := 3
	for {
		if negAckCount == maxRetries {
			return PublishError{message: "failed to push: max retries for publishing ack exceeded"}
		}

		// Queue argument refers to the exchange for RabbitMQ.
		err := client.unsafePush(ctx, queue, routingKey, v)
		if err != nil {
			if errCount == maxRetries {
				return PublishError{message: fmt.Sprintf("failed to push as maximum retries exceeded: %v", err)}
			}
			select {
			case <- client.done:
				return PublishError{message: fmt.Sprintf("failed to push as server is shutting down: %v", err)}
			case <- time.After(delay):
				delay *= exponentialBackoff
				errCount += 1
			}
			continue
		}
		confirm := <- client.notifyConfirm
		if confirm.Ack {
			return nil
		}
		negAckCount += 1
	}
}