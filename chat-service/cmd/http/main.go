package main

import (
	"chat-service/internal"
	"chat-service/internal/config"
	"chat-service/internal/delivery/kafka"
	"chat-service/internal/delivery/rest"
	svcdis "chat-service/internal/delivery/service-discovery"
	ws "chat-service/internal/delivery/websocket"
	uc "chat-service/internal/usecase"
	"context"
	"fmt"
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

	// Create service discovery client.
	sd, err := svcdis.New(cfg)
	if err != nil {
		logger.Fatal("error creating etcd client", zap.String("trace", err.Error()))
	}
	defer sd.Close()

	// Create kafka dependency.
	eb := kafka.New(cfg)

	// Create websocket dependency.
	sc := ws.New()

	// Create usecase.
	uc := uc.NewUseCaseService(eb, sc)

	// Init websocket hub.
	hub := ws.NewHub(uc)
	brokers := strings.Split(cfg.Kafka.BrokerAddresses, ",")
	go hub.Run(ctx, brokers)

	// Create server.
	s := rest.New(logger, uc, ws.ServeWs)

	// Run server.
	go func() {
		fmt.Printf("starting REST server in port %v", cfg.Port)
		if err := s.Echo.Start(fmt.Sprintf(":%v", cfg.Port)); err != nil {
			gracefulShutdown(ctx, s, eb, hub)
			logger.Fatal("failed to start REST server", zap.String("trace", err.Error()))
		}
	}()

	// Send server metadata to service discovery.
	go sd.SendHeartbeatToServiceDiscovery(ctx)

	// Create ctx for listening to SIGINT and SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	gracefulShutdown(ctx, s, eb, hub)
}

func gracefulShutdown(ctx context.Context, s *rest.RestServer, k *kafka.KafkaClient, hub *ws.Hub) {
	fmt.Println("performing graceful shutdown...")
	k.Close()
	hub.Close()
	if err := s.Echo.Shutdown(ctx); err != nil {
		logger.Error("failed to shutdown REST server", zap.String("trace", err.Error()))
	}
}