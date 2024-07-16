package kafka

import (
	"message-service/internal"
	"message-service/internal/config"
	"net"
	"strconv"
	"strings"

	"github.com/segmentio/kafka-go"
)

var logger, _ = internal.WireLogger()

type Topic int32

const (
	Messages Topic = iota
)

func (t Topic) String() string {
	switch t {
	case Messages:
		return "messages"
	}
	return "unknown"
}

func (k Kafka) CreateKafkaTopics(cfg *config.Config) error {
	conn, err := kafka.Dial("tcp", strings.Split(cfg.Kafka.BrokerAddresses, ",")[0])
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()

	/* 
	Config for each topic must be explicitly set for the following reasons:
	- You cannot decrease the number of partitions 
	- Increasing the partitions will force a rebalance
	- ReplicationFactor cannot be greater than the number of brokers available
	*/
	topicConfigs := []kafka.TopicConfig{
		{
			Topic:  Messages.String(),
			NumPartitions: 10,
			ReplicationFactor: 1,
		},
	}

	if err := controllerConn.CreateTopics(topicConfigs...); err != nil {
		return err
	}

	return nil
}

