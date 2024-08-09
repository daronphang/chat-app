package main

import (
	"broadcast-service/internal"
	"broadcast-service/internal/config"
	g "broadcast-service/internal/delivery/grpc"
	k "broadcast-service/internal/delivery/kafka"
	rmq "broadcast-service/internal/delivery/rabbitmq"
	"broadcast-service/internal/repository"
	"broadcast-service/internal/usecase"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var logger *zap.Logger

func main() {
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

	// Create gRPC client dependency.
	u, err := g.NewClient(cfg)
	if err != nil {
		logger.Fatal("error setting up grpc message client", zap.String("trace", err.Error()))
	}

	// Setup rabbitmq dependency.
	hub := rmq.NewHub(logger)
	mb := rmq.NewClient(logger)
	hub.AddClient(mb)

	// Create db dependency.
	db := repository.New(cfg)

	// Create kafka dependency.
	eb := k.New(cfg)

	// Create usecase.
	uc := usecase.NewUseCaseService(db, u, mb, eb)

	// Create a single TCP rabbitmq connection for all goroutines to use.
	go hub.Run(cfg)

	// Listen to protocol and port.
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.Port))
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to listen on port %v", cfg.Port), zap.String("trace", err.Error()))
	}

	// Create server.
	s := g.NewServer(logger, uc)

	// Start server.
	go func() {
		fmt.Printf("starting gRPC server in port %v...", cfg.Port)
		if err := s.Serve(lis); err != nil {
			logger.Fatal("failed to start gRPC server", zap.String("trace", err.Error()))
		}
	}()

	// Create ctx for listening to SIGINT and SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	<-ctx.Done()
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	gracefulShutdown(s, db, hub, eb)
}

func gracefulShutdown(s *grpc.Server, db *repository.RedisClient, hub *rmq.RabbitMQHub, eb *k.KafkaClient) {
	fmt.Println("performing graceful shutdown...")
	s.GracefulStop()
	db.Close()
	hub.Close()
	eb.Close()
}