package kafka

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
	"user-service/internal"
	"user-service/internal/config"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

const (
	reconnectDelay = 2 * time.Second
	exponentialBackoff = 2
	pingInterval = 3 * time.Second
)

var (
	logger, _ = internal.WireLogger()
)

type KafkaClient struct {
	controllerConn 	*kafka.Conn
	done 			chan bool
}

func New(cfg *config.Config) *KafkaClient {
	kc := &KafkaClient{
		done: make(chan bool),
	}
	brokerAddresses := strings.Split(cfg.Kafka.BrokerAddresses, ",")
	go kc.handleControllerReconnect(brokerAddresses)
	return kc
}

func (kc *KafkaClient) Close() {
	kc.done <- true
	if err := kc.controllerConn.Close(); err != nil {
		logger.Error(
			"failed to close Kafka conn",
			zap.String("trace", err.Error()),
		)
	}
}

func (kc *KafkaClient) connectToController(brokerAddresses []string) error {
	var conn *kafka.Conn
	var err error

	for _, address := range brokerAddresses {
		conn, err = kafka.Dial("tcp", address)
		if err == nil {
			break
		} 
	}

	if conn == nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	kc.controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}
	return nil
}

func (kc *KafkaClient) handleControllerReconnect(brokerAddresses []string) {
	delay := reconnectDelay
	for {
		if err := kc.connectToController(brokerAddresses); err != nil {
			logger.Warn(
				fmt.Sprintf("failed to init kafka controller conn, retrying after %v...", delay),
				zap.String("trace", err.Error()),
			)
			select {
			case <-kc.done:
				return
			case <-time.After(delay):
				delay *= time.Duration(exponentialBackoff)
			}
			continue
		}
		delay = reconnectDelay

		// There is no native ping/health check mechanism to check if conn is closed.
		// Workaround is to call Brokers at periodic intervals.
		ping:
		for {
			select {
			case <-kc.done:
				return
			case <- time.After(pingInterval):
				_, err := kc.controllerConn.Brokers()
				if err != nil {
					break ping
				}
			}
		}
	}
}

/*
Topics are explicitly configured for the following reasons:
- You cannot decrease the number of partitions 
- Increasing the partitions will force a re-balance
- ReplicationFactor cannot be greater than the number of brokers available
- Having different consumer groups will read from the same partition and result in duplication
*/

func (kc *KafkaClient) CreateUserTopic(ctx context.Context, topic string) error {
	tc := kafka.TopicConfig{
		Topic: topic,
		NumPartitions: 1,
		ReplicationFactor: 1,
	}
	if err := kc.controllerConn.CreateTopics(tc); err != nil {
		return err
	}
	
	return nil
}