package kafka

import (
	"net"
	"session-service/internal"
	"session-service/internal/config"
	"session-service/internal/domain"
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

func CreateKafkaTopic(cfg *config.Config, topicCfg domain.BrokerTopicConfig) error {
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

	// Create topic.
	if err := controllerConn.CreateTopics(
		kafka.TopicConfig{
			Topic: topicCfg.Topic,
			NumPartitions: topicCfg.Partitions,
			ReplicationFactor: topicCfg.ReplicationFactor,
		},
	); err != nil {
		return err
	}
	
	return nil
}
