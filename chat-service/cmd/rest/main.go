package main

import (
	"chat-service/internal"
	"chat-service/internal/config"
	"chat-service/internal/delivery/kafka"
	"chat-service/internal/delivery/rest"
	ws "chat-service/internal/delivery/websocket"
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
		panic(fmt.Sprintf("error setting up logger: %v", err))
    }

	// Create UseCase with dependencies.
	eb := kafka.New(cfg)
	sc := ws.New()
	uc := uc.NewUseCaseService(eb, sc)

	// Init websocket hub.
	go sc.Hub.Run(ctx, uc)

	// Create Kafka topics.
	if err := eb.CreateKafkaTopics(cfg); err != nil {
		logger.Fatal("error creating kafka topics", zap.String("trace", err.Error()))
	}

	// Create server.
	s := rest.New(ctx, logger, uc, sc.Hub)

	// Server and dependencies have been successfully initialized from here.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Run server.
	go func() {
		fmt.Printf("starting ECHO server in port %v", cfg.Port)
		if err := s.Echo.Start(fmt.Sprintf(":%v", cfg.Port)); err != nil {
			logger.Fatal("server error", zap.String("trace", err.Error()))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("shutting down gracefully...")
	gracefulShutdown(ctx, s, eb)
}

func gracefulShutdown(ctx context.Context, s *rest.Server, k kafka.Kafka) {
	if err := k.Writer.Close(); err != nil {
		logger.Error("failed to close Kafka writer", zap.String("trace", err.Error()))
	}

	if err := s.Echo.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("failed to shutdown echo server: %v", err.Error()))
	}
}