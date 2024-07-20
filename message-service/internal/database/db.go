package db

import (
	"context"
	"fmt"
	"message-service/internal/config"
	"strings"
	"sync"

	"github.com/gocql/gocql"
)

var (
	syncOnce sync.Once
	cluster *gocql.ClusterConfig
)

const (
	keyspaceMessage = "message"
)

func ProvideCluster(cfg *config.Config) *gocql.ClusterConfig {
	syncOnce.Do(func() {
		cluster = gocql.NewCluster(strings.Split(cfg.Cassandra.HostAddresses, ",")...)
		cluster.ProtoVersion = 4
		cluster.Consistency = gocql.Quorum
	})
	return cluster
}

func CreateKeyspace(ctx context.Context, cluster *gocql.ClusterConfig) error {
	session, err := cluster.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()
	if err := session.Query(
		fmt.Sprintf("CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class': 'NetworkTopologyStrategy', 'replication_factor': 1}", keyspaceMessage),
	).WithContext(ctx).Exec(); err != nil {
		return err
	}
	return nil
}

func ProvideDBSession(cfg *config.Config, cluster *gocql.ClusterConfig) (*gocql.Session, error) {
	// Should not be parameterized but explicitly defined.
	cluster.Keyspace = keyspaceMessage

	// Session is safe to use from multiple goroutines.
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}