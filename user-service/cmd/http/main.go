package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
	"user-service/internal"
	"user-service/internal/config"
	g "user-service/internal/delivery/grpc"
	k "user-service/internal/delivery/kafka"
	svcdis "user-service/internal/delivery/service-discovery"
	"user-service/internal/repository"
	"user-service/internal/usecase"

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

	// Create usecase with dependencies.
	db, err := repository.New(ctx, cfg)
	if err != nil {
		logger.Fatal("error creating db", zap.String("trace", err.Error()))
	}
	sd, err := svcdis.New(cfg)
	if err != nil {
		logger.Fatal("error connecting to etcd", zap.String("trace", err.Error()))
	}
	kc := k.New(cfg)
	uc := usecase.NewUseCaseService(db, sd, kc)

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
	gracefulShutdown(s, db, sd, kc)
}

func gracefulShutdown(s *grpc.Server, db *repository.Querier, sd *svcdis.ServiceDiscoveryClient, kc *k.KafkaClient) {
	fmt.Println("performing graceful shutdown...")
	s.GracefulStop()
	db.Close()
	sd.Close()
	kc.Close()
}