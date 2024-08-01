package main

import (
	"context"
	"fmt"
	"message-service/internal"
	"message-service/internal/config"
	g "message-service/internal/delivery/grpc"
	"message-service/internal/delivery/kafka"
	rmq "message-service/internal/delivery/rabbitmq"
	"message-service/internal/repository"
	"message-service/internal/usecase"
	"net"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	// Setup DB.
	if err := repository.SetupDB(ctx, cfg); err != nil {
		logger.Fatal("error setting up DB", zap.String("trace", err.Error()))
	}

	// Create gRPC client.
	client, err := g.NewClient(cfg)
	if err != nil {
		logger.Fatal("error setting up grpc user client", zap.String("trace", err.Error()))
	}

	// Create usecase with dependencies.
	db, err := repository.New(cfg)
	if err != nil {
		logger.Fatal("error setting up DB instance", zap.String("trace", err.Error()))
	}
	mb := &rmq.RabbitMQClient{} // Dummy as it is not needed in server.
	eb := &kafka.KafkaClient{}	// Dummy as it is not needed in server.
	uc := usecase.NewUseCaseService(mb, eb, db, client)

	// Listen to protocol and port.
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.Port))
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to listen on port %v", cfg.Port), zap.String("trace", err.Error()))
	}

	// Create server.
	s := g.New(logger, uc)

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
	gracefulShutdown(s, db)
}

func gracefulShutdown(s *grpc.Server, db *repository.Querier) {
	fmt.Println("performing graceful shutdown...")

	s.GracefulStop()

	db.Close()
}