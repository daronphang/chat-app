package svcdis

import (
	"chat-service/internal/config"
	"strings"
	"time"

	etcd "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	Creation of client does not guarantee that connection is established.
	Operations on etcd will hang until connection is resolved.

	When using DialTimeout:
	1. The code doesn't block at the New() call, but instead in the following cli.Status() call
	2. The code blocks indefinitely, so the DialTimeout doesn't seem to have any effect

	Need to use grpc.WithBlock().
	*/
	cli, err := etcd.New(etcd.Config{
		Endpoints: 		strings.Split(cfg.Etcd.BrokerAddresses, ","),
		DialTimeout: 	3 * time.Second,
		DialOptions: 	[]grpc.DialOption{grpc.WithBlock()},
		// DialKeepAliveTime:    	2 * time.Second,
		// DialKeepAliveTimeout: 	2 * time.Second,
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