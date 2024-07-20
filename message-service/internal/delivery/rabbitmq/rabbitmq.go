package rmq

import (
	"fmt"
	"message-service/internal/config"
	"message-service/internal/domain"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

const (
	reconnectDelay = 2 * time.Second
	reInitDelay = 3 * time.Second
	resendDelay = 2 * time.Second
	exponentialBackoff = 2
)

type PublishError struct {
	message string
}
func (e PublishError) Error() string {
	return e.message
}

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
				delay *= exponentialBackoff
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
		case <- hub.notifyConnClose:
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

	if err := hub.connection.Close(); err != nil {
		hub.logger.Error(
			"unable to close rabbitmq connection",
			zap.String("trace", err.Error()),
		)
	}
}

type RabbitMQClient struct {
	m               *sync.Mutex
	logger          *zap.Logger
	connection      *amqp.Connection
	channel         *amqp.Channel
	done            chan bool
	notifyConnOpen 	chan bool
	notifyConnClose chan bool
	notifyChanClose chan *amqp.Error
	notifyConfirm   chan amqp.Confirmation
	isReady         bool
}

func (client *RabbitMQClient) declareAndBindExchangesAndQueues(ch *amqp.Channel) error {
	// Errors returned will close the channel.
	// Once an exchange is declared, its type cannot be changed.
	// Declaring a queue is idempotent.

	if err := ch.ExchangeDeclare(
		domain.NotificationQueueConfig.Queue,
		amqp.ExchangeDirect,
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	); err != nil {
		return err
	}

	if _, err := ch.QueueDeclare(
		domain.NotificationQueueConfig.RoutingKeys[0], // queue name
		false, // durable
		false, // auto-deleted
		false, // exclusive
		false, // no wait
		nil, // arguments
	); err != nil {
		return err
	}

	if err := ch.QueueBind(
		domain.NotificationQueueConfig.RoutingKeys[0], // queue name
		domain.NotificationQueueConfig.RoutingKeys[0], // routing key
		domain.NotificationQueueConfig.Queue, // exchange
		false, // no-wait
		nil, // arguments
	); err != nil {
		return err
	}

	return nil
}

func (client *RabbitMQClient) initChannel() (*amqp.Channel, error) {
	ch, err := client.connection.Channel()
	if err != nil {
		return nil, err
	}

	// Puts channel into confirm mode to ensure all publishings have ack.
	err = ch.Confirm(false)
	if err != nil {
		return nil, err
	}

	// Declare and bind exchanges and queues.
	if err := client.declareAndBindExchangesAndQueues(ch); err != nil {
		return nil, err
	}
	return ch, nil
}

func (client *RabbitMQClient) updateConnection(conn *amqp.Connection) {
	client.connection = conn
	client.notifyConnOpen <- true
}

func (client *RabbitMQClient) updateChannel(ch *amqp.Channel) {
	client.channel = ch
	client.notifyChanClose = make(chan *amqp.Error, 1)
	client.channel.NotifyClose(client.notifyChanClose)
	client.notifyConfirm = make(chan amqp.Confirmation, 1)
	client.channel.NotifyPublish(client.notifyConfirm)
}

func (client *RabbitMQClient) handleReInitChannel() bool {
	<- client.notifyConnOpen 
	delay := reInitDelay
	for {
		client.m.Lock()
		client.isReady = false
		client.m.Unlock()

		ch, err := client.initChannel()
		if err != nil {
			client.logger.Warn(
				fmt.Sprintf("failed to init rabbitmq channel, retrying after %v...", delay),
				zap.String("trace", err.Error()),
			)
			select {
			case <-client.done:
				return true
			case <-client.notifyConnClose:
				return false
			case <-time.After(delay):
				delay *= exponentialBackoff
			}
			continue
		}

		delay = reInitDelay
		client.updateChannel(ch)
		client.m.Lock()
		client.isReady = true
		client.m.Unlock()

		select {
		case <- client.done:
			return true
		case <-client.notifyConnClose:
			client.m.Lock()
			client.isReady = false
			client.m.Unlock()
			return false
		case <- client.notifyChanClose:
		}
	}
}

func NewClient(logger *zap.Logger) *RabbitMQClient {
	client := &RabbitMQClient{
		m: &sync.Mutex{},
		logger: logger,
		done: make(chan bool),
	}

	go func() {
		for {
			if done := client.handleReInitChannel(); done {
				break
			}
		}
	}()
	return client
}