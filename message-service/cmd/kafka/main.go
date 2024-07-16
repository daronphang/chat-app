package main

import (
	"context"
	"fmt"
	"message-service/internal"
	"message-service/internal/config"
	"message-service/internal/delivery/kafka"
	"message-service/internal/usecase"
	"os"
	"os/signal"
	"time"

	kg "github.com/segmentio/kafka-go"
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
		panic(fmt.Sprintf("error setting up logger: %v", err))
    }

	// Create usecase and dependencies.
	uc := usecase.NewUseCaseService()

	// Important to close readers when a process exits.
	readers := make([]*kg.Reader, 0, 10)

	// Spin up goroutines equivalent to the number of partitions per topic.
	// One consumer per goroutine.
	topic := kafka.Messages.String()
	consumerGroupID := cfg.Kafka.MessageConsumerGroupID
	for range 10 {
		go func(){
			k := kafka.New(cfg, consumerGroupID, topic)
			readers = append(readers, k.Reader)
			for {
				if ok := k.ConsumeMsg(ctx, uc); !ok {
					break
				}
			}
		}()
	}

	fmt.Println("running goroutines for Kafka readers...")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("shutting down gracefully...")
	gracefulShutdown(readers)
}

func gracefulShutdown(readers []*kg.Reader) {
	for _, r := range readers {
		if err := r.Close(); err != nil {
			logger.Error(
				"unable to close Kafka reader",
				zap.String("trace", err.Error()),
			)
		}
	}
}