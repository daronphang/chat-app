package svcdis

import (
	"strings"
	"time"
	"user-service/internal"
	"user-service/internal/config"

	etcd "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	logger, _ = internal.WireLogger()
)

type ServiceDiscoveryClient struct {
	cli *etcd.Client
}

func New(cfg *config.Config) (*ServiceDiscoveryClient, error) {
	/*
	etcd v3 uses gRPC for remote procedure calls, and clientv3 
	uses grpc-go to connect to etcd. Make sure to close the 
	client after using it. If the client is not closed, the 
	connection will have leaky goroutines.
	*/
	cli, err := etcd.New(etcd.Config{
		Endpoints: strings.Split(cfg.Etcd.BrokerAddresses, ","),
		DialTimeout: 	3 * time.Second,
		DialOptions: 	[]grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		return nil, err
	}

	return &ServiceDiscoveryClient{cli: cli}, nil
}

func (s *ServiceDiscoveryClient) Close() {
	if err := s.cli.Close(); err != nil {
		logger.Error(
			"unable to close etcd client",
			zap.String("trace", err.Error()),
		)
	}
}