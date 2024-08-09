package main

import (
	"context"
	"fmt"
	"message-service/internal"
	"message-service/internal/config"
	g "message-service/internal/delivery/grpc"
	k "message-service/internal/delivery/kafka"
	"message-service/internal/repository"
	"message-service/internal/usecase"
	"os"
	"os/signal"
	"strings"
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
	if err := k.CreateKafkaTopics(cfg); err != nil {
		logger.Fatal("error creating kafka topics", zap.String("trace", err.Error()))
	}

	// Setup DB.
	if err := repository.SetupDB(ctx, cfg); err != nil {
		logger.Fatal("error setting up DB", zap.String("trace", err.Error()))
	}

	// Create kafka dependency.
	kc := k.New(cfg)

	// Create db dependency.
	db, err := repository.New(cfg)
	if err != nil {
		logger.Fatal("error setting up DB instance", zap.String("trace", err.Error()))
	}

	// Create gRPC client dependency.
	client, err := g.NewClient(cfg)
	if err != nil {
		logger.Fatal("error setting up grpc user client", zap.String("trace", err.Error()))
	}	

	// Create usecase.
	uc := usecase.NewUseCaseService(kc, db, client)

	// Spin up goroutines equivalent to the number of partitions per topic.
	// All goroutines share the same consumer group.
	// One consumer per goroutine.
	fmt.Println("running goroutines for reading Kafka message partitions...")
	consumers := make([]*k.KafkaConsumer, 0)
	brokers := strings.Split(cfg.Kafka.BrokerAddresses, ",")
	consumerGroup := "messageConsumer"
	for range k.MessagePartitions {
		// For each goroutine, will have a separate Kafka consumer.
		c := k.NewConsumer(brokers, consumerGroup, k.MessageTopic)
		consumers = append(consumers, c)
		go c.ConsumeFromMessageTopic(ctx, uc)
	}

	// Create ctx for listening to SIGINT and SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	<-ctx.Done()
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	gracefulShutdown(kc, consumers, db)
}

func gracefulShutdown(kc *k.KafkaClient, consumers []*k.KafkaConsumer, db *repository.Querier) {
	fmt.Println("performing graceful shutdown...")
	kc.Close()
	db.Close()

	for _, c := range consumers {
		c.Close()
	}	
}