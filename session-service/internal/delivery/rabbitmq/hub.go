package rmq

import (
	"fmt"
	"session-service/internal/config"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type RabbitMQHub struct {
	clients			[]*RabbitMQClient
	connection		*amqp.Connection
	done			chan bool
	notifyConnClose chan *amqp.Error
	logger 			*zap.Logger
}

func NewHub(logger *zap.Logger) *RabbitMQHub {
	hub := &RabbitMQHub{
		logger: logger,
		done: make(chan bool),
	}
	return hub
}

func (hub *RabbitMQHub) AddClient(client *RabbitMQClient) {
	hub.clients = append(hub.clients, client)
}

// Goroutines will maintain a separate channel, but share the same TCP connection.
func (hub *RabbitMQHub) connect(addr string) (*amqp.Connection, error) {
	// amqp://guest:guest@localhost:5672/
	conn, err := amqp.Dial(addr)

	if err != nil {
		return nil, err
	}
	fmt.Println("connected to rabbitmq!")
	return conn, nil
}

func (hub *RabbitMQHub) updateConnection(conn *amqp.Connection) {
	hub.connection = conn
	hub.notifyConnClose = make(chan *amqp.Error, 1)
	hub.connection.NotifyClose(hub.notifyConnClose)
}

// To be executed in a goroutine.
// Handles reconnection with exponential backoff.
// Notifies all clients of new connection if required.
func (hub *RabbitMQHub) Run(cfg *config.Config) {
	addr := fmt.Sprintf("amqp://%s", cfg.RabbitMQ.HostAddress)
	delay := reconnectDelay
	for {
		conn, err := hub.connect(addr)
		if err != nil {
			hub.logger.Warn(
				fmt.Sprintf("rabbitmq connection failed, retrying after %v...", delay),
				zap.String("trace", err.Error()),
			)
			select {
			case <-hub.done:
				return
			case <-time.After(delay):
				delay *= time.Duration(exponentialBackoff)
			}
			continue
		}
		// Reset delay after connection succeeded.
		delay = reconnectDelay

		hub.updateConnection(conn)
		
		// Notify clients of new connection.
		for _, client := range hub.clients {
			client.updateConnection(conn)
		}
		
		select {
		case <-hub.done:
			return
		case <-hub.notifyConnClose:
			for _, client := range hub.clients {
				client.notifyConnClose <- true
			}
		}
	}
}

func (hub *RabbitMQHub) Close() {
	for _, client := range hub.clients {
		if !client.isReady {
			return
		}
		close(client.done)
		if err := client.channel.Close(); err != nil {
			client.logger.Error(
				"unable to close rabbitmq channel",
				zap.String("trace", err.Error()),
			)
		}
	}
	close(hub.done)
	if err := hub.connection.Close(); err != nil {
		hub.logger.Error(
			"unable to close rabbitmq connection",
			zap.String("trace", err.Error()),
		)
	}
}