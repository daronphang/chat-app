package kafka

import (
	"message-service/internal"
	"message-service/internal/config"
	"message-service/internal/domain"
	"net"
	"strconv"
	"strings"

	"github.com/segmentio/kafka-go"
)

var (
	logger, _ = internal.WireLogger()
)

type KafkaClient struct {
	Reader *kafka.Reader
	Writer *kafka.Writer
}

func New(r *kafka.Reader, w *kafka.Writer) *KafkaClient {
	return &KafkaClient{Reader: r, Writer: w}
}

func CreateKafkaTopics(cfg *config.Config) error {
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
	topicConfigs := []kafka.TopicConfig{
		{
			Topic: domain.MessageTopicConfig.Topic,
			NumPartitions: domain.MessageTopicConfig.Partitions,
			ReplicationFactor: domain.MessageTopicConfig.ReplicationFactor,
		},
	}
	if err := controllerConn.CreateTopics(topicConfigs...); err != nil {
		return err
	}

	return nil
}