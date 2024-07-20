package main

import (
	"context"
	"fmt"
	"message-service/internal"
	"message-service/internal/config"
	"message-service/internal/delivery/kafka"
	rmq "message-service/internal/delivery/rabbitmq"
	"message-service/internal/domain"
	"message-service/internal/usecase"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	ctx := context.Background()

	// Create config.
	cfg, err := config.ProvideConfig()
	if err != nil {
		panic(fmt.Sprintf("error reading config file: %v", err))
    }

	// Create logger.
	logger, err = internal.WireLogger()
	if err != nil {
		logger.Fatal("error setting up logger", zap.String("trace", err.Error()))
    }

	// Create Kafka topics.
	if err := kafka.CreateKafkaTopics(cfg); err != nil {
		logger.Fatal("error creating kafka topics", zap.String("trace", err.Error()))
	}

	// Create hub for rabbitmq.
	hub := rmq.NewHub(logger)

	// Spin up goroutines equivalent to the number of partitions per topic.
	// One consumer per goroutine.
	fmt.Println("running goroutines for reading Kafka topics...")
	for range domain.MessageTopicConfig.Partitions {
		go func(){
			k := kafka.New(cfg, domain.MessageTopicConfig.ConsumerGroupID, domain.MessageTopicConfig.Topic)

			// Create usecase and dependencies.
			mb := rmq.NewClient(logger)
			hub.AddClient(mb)
			uc := usecase.NewUseCaseService(mb)

			for {
				if ok := k.ConsumeMsgFromMessageTopic(ctx, uc); !ok {
					break
				}
			}
		}()
	}

	// Create a single TCP rabbitmq connection for all goroutines to use.
	go hub.Run(cfg)

	// Create ctx for listening to SIGINT and SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	gracefulShutdown(hub)
}

func gracefulShutdown(hub *rmq.RabbitMQHub) {
	fmt.Println("performing graceful shutdown...")

	// close RabbitMQ.
	hub.Close()
}