package svcdis

import (
	"context"
	"encoding/json"
	"user-service/internal/domain"

	etcd "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

func (sdc *ServiceDiscoveryClient) GetServersMetdata(ctx context.Context) ([]domain.ServerMetadata, error) {
	resp, err := sdc.cli.Get(ctx, "chat-server", etcd.WithPrefix())
	if err != nil {
		return nil, err
	}

	rv := make([]domain.ServerMetadata, 0)
	for _, kv := range resp.Kvs {
		p := new(domain.ServerMetadata)
		v := kv.Value
		if err := json.Unmarshal(v, p); err != nil {
			logger.Error(
				"unable to unmarshal value in etcd",
				zap.String("payload", string(v)),
				zap.String("trace", err.Error()),
			)
		}
		rv = append(rv, *p)
	}

	return rv, nil
}