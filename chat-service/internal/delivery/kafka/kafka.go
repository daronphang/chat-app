package kafka

import (
	"chat-service/internal"
	"chat-service/internal/config"
	"chat-service/internal/domain"
	"net"
	"strconv"
	"strings"

	"github.com/segmentio/kafka-go"
)

var (
	logger, _ = internal.WireLogger()
)

type KafkaClient struct {
	Writer *kafka.Writer
}

func New(cfg *config.Config) *KafkaClient {
	return &KafkaClient{Writer: NewWriter(cfg)}
}

func CreateKafkaTopics(cfg *config.Config, topicCfgs ...domain.BrokerTopicConfig) error {
	// Connect to cluster.
	conn, err := kafka.Dial("tcp", strings.Split(cfg.Kafka.BrokerAddresses, ",")[0])
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	// Get leader controller.
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()

	// Create topics.
	var tcfgs []kafka.TopicConfig
	for _, c := range topicCfgs {
		temp := kafka.TopicConfig{
			Topic: c.Topic,
			NumPartitions: c.Partitions,
			ReplicationFactor: c.ReplicationFactor,
		}
		tcfgs = append(tcfgs, temp)
	} 
	if err := controllerConn.CreateTopics(tcfgs...); err != nil {
		return err
	}

	return nil
}