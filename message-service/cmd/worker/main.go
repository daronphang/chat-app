package main

import (
	"context"
	"fmt"
	"message-service/internal"
	"message-service/internal/config"
	k "message-service/internal/delivery/kafka"
	rmq "message-service/internal/delivery/rabbitmq"
	"message-service/internal/domain"
	"message-service/internal/repository"
	"message-service/internal/usecase"
	"os"
	"os/signal"
	"time"

	"github.com/segmentio/kafka-go"
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

	// Create Kafka topic.
	if err := k.CreateKafkaTopic(cfg, domain.MessageTopicConfig); err != nil {
		logger.Fatal("error creating kafka topic", zap.String("trace", err.Error()))
	}

	// Create Kafka writer.
	// Safe to use across goroutines.
	kw := k.NewWriter(cfg)

	// Create hub for rabbitmq.
	hub := rmq.NewHub(logger)

	// Setup DB.
	if err := repository.SetupDB(ctx, cfg); err != nil {
		logger.Fatal("error setting up DB", zap.String("trace", err.Error()))
	}

	// Create db instance.
	db, err := repository.New(cfg)
	if err != nil {
		logger.Fatal("error setting up DB instance", zap.String("trace", err.Error()))
	}

	// Spin up goroutines equivalent to the number of partitions per topic.
	// One consumer per goroutine.
	fmt.Println("running goroutines for reading Kafka topics...")
	for range domain.MessageTopicConfig.Partitions {
		// For each goroutine, will have separate RabbitMQ channel and Kafka reader.
		kr := k.NewReader(cfg, domain.MessageTopicConfig.ConsumerGroupID, domain.MessageTopicConfig.Topic)
		k := k.New(kr, kw)

		mb := rmq.NewClient(logger)
		hub.AddClient(mb)
		uc := usecase.NewUseCaseService(mb, k, db)

		go k.ConsumeMsgFromMessageTopic(ctx, uc)
	}

	// Create a single TCP rabbitmq connection for all goroutines to use.
	go hub.Run(cfg)

	// Create ctx for listening to SIGINT and SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	<-ctx.Done()
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	gracefulShutdown(hub, kw, db)
}

func gracefulShutdown(hub *rmq.RabbitMQHub, kw *kafka.Writer, db *repository.Querier) {
	fmt.Println("performing graceful shutdown...")

	hub.Close()

	if err := kw.Close(); err != nil {
		logger.Error(
			"unable to close kafka writer",
			zap.String("trace", err.Error()),
		)
	}

	db.Close()
}