package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"presence-service/internal"
	"presence-service/internal/config"
	g "presence-service/internal/delivery/grpc"
	"presence-service/internal/repository"
	"presence-service/internal/usecase"
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

	// Create gRPC client.
	mc, err := g.NewClient(cfg)
	if err != nil {
		logger.Fatal("error setting up grpc message client", zap.String("trace", err.Error()))
	}

	// Create usecase with dependencies.
	db := repository.New(cfg)
	uc := usecase.NewUseCaseService(db, mc)

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
	gracefulShutdown(s, db)
}

func gracefulShutdown(s *grpc.Server, db *repository.RedisClient) {
	fmt.Println("performing graceful shutdown...")

	s.GracefulStop()

	db.Close()
}