package main

import (
	"chat-service/internal"
	"chat-service/internal/config"
	"chat-service/internal/delivery/kafka"
	"chat-service/internal/delivery/rest"
	ws "chat-service/internal/delivery/websocket"
	"chat-service/internal/domain"
	uc "chat-service/internal/usecase"
	"context"
	"fmt"
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
	if err := kafka.CreateKafkaTopics(cfg, domain.MessageTopicConfig); err != nil {
		logger.Fatal("error creating kafka topics", zap.String("trace", err.Error()))
	}

	// Create UseCase with dependencies.
	eb := kafka.New(cfg)
	sc := ws.New()
	uc := uc.NewUseCaseService(eb, sc)

	// Init websocket hub.
	hub := ws.NewHub(uc)
	go hub.Run(ctx, cfg)

	// Create server.
	s := rest.New(logger, uc, ws.ServeWs)

	// Run server.
	go func() {
		fmt.Printf("starting REST server in port %v", cfg.Port)
		if err := s.Echo.Start(fmt.Sprintf(":%v", cfg.Port)); err != nil {
			gracefulShutdown(ctx, s, eb)
			logger.Fatal("failed to start REST server", zap.String("trace", err.Error()))
		}
	}()

	// Create ctx for listening to SIGINT and SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	gracefulShutdown(ctx, s, eb)
}

func gracefulShutdown(ctx context.Context, s *rest.RestServer, k *kafka.KafkaClient) {
	fmt.Println("performing graceful shutdown...")
	if err := k.Writer.Close(); err != nil {
		logger.Error("failed to close Kafka writer", zap.String("trace", err.Error()))
	}

	if err := s.Echo.Shutdown(ctx); err != nil {
		logger.Error("failed to shutdown REST server", zap.String("trace", err.Error()))
	}
}