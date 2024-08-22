package kafka

import (
	"message-service/internal"
	"message-service/internal/config"
	"net"
	"strconv"
	"strings"

	"github.com/segmentio/kafka-go"
)

var (
	logger, _ = internal.WireLogger()
)

const (
	MessageTopic 		string 	= "message"
	MessagePartitions 	int 	= 10
)

type KafkaClient struct {
	writer *kafka.Writer
}

func New(cfg *config.Config) *KafkaClient {
	return &KafkaClient{writer: newWriter(cfg)}
}

func (kc *KafkaClient) Close() {
	kc.writer.Close()
}

/*
Topics are explicitly configured for the following reasons:

- You cannot decrease the number of partitions 

- Increasing the partitions will force a re-balance

- ReplicationFactor cannot be greater than the number of brokers available

- Having different consumer groups will read from the same partition and result in duplication
*/
func CreateKafkaTopics(cfg *config.Config) error {
	// Connect to cluster.
	var conn *kafka.Conn
	var err error

	for _, address := range strings.Split(cfg.Kafka.BrokerAddresses, ",") {
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

	// Get leader controller.
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()

	// Create topics.
	if err := controllerConn.CreateTopics(
		kafka.TopicConfig{
			Topic: MessageTopic,
			NumPartitions: MessagePartitions,
			ReplicationFactor: 1,
		},
	); err != nil {
		return err
	}
	
	return nil
}

