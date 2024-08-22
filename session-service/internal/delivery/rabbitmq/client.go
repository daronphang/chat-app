package rmq

import (
	"fmt"
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

var (
	notificationQueueConfig = BrokerQueueConfig{
		Queue: 			"notificationExchange",
		RoutingKeys: 	[]string{"email"},
	}
)

type BrokerQueueConfig struct {
	Queue 		string
	RoutingKeys []string
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
		notificationQueueConfig.Queue,
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
		notificationQueueConfig.RoutingKeys[0], // queue name
		false, // durable
		false, // auto-deleted
		false, // exclusive
		false, // no wait
		nil, // arguments
	); err != nil {
		return err
	}

	if err := ch.QueueBind(
		notificationQueueConfig.RoutingKeys[0], // queue name
		notificationQueueConfig.RoutingKeys[0], // routing key
		notificationQueueConfig.Queue, // exchange
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
	<-client.notifyConnOpen 
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
				delay *= time.Duration(exponentialBackoff)
			}
			continue
		}

		delay = reInitDelay
		client.updateChannel(ch)
		client.m.Lock()
		client.isReady = true
		client.m.Unlock()

		select {
		case <-client.done:
			return true
		case <-client.notifyConnClose:
			client.m.Lock()
			client.isReady = false
			client.m.Unlock()
			return false
		case <-client.notifyChanClose:
		}
	}
}

func NewClient(logger *zap.Logger) *RabbitMQClient {
	client := &RabbitMQClient{
		m: &sync.Mutex{},
		logger: logger,
		done: make(chan bool),
		notifyConnOpen: make(chan bool),
		notifyConnClose: make(chan bool),
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