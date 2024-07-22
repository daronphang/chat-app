package repository

import (
	"context"
	"message-service/internal/config"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

type Querier struct {
	session *gocql.Session
}

func New(ctx context.Context, cfg *config.Config, logger *zap.Logger) *Querier {
	// Connect to Cassandra.
	cluster := provideCluster(cfg)

	// Create keyspace if required.
	if err := createKeyspace(ctx, cluster); err != nil {
		logger.Fatal("error creating keyspace", zap.String("trace", err.Error()))
	}

	// Migrate db.
	if err := migrateDB(cfg, logger); err != nil {
		logger.Fatal("error migrating DB", zap.String("trace", err.Error()))
	}

	// Create session.
	session, err := provideSession(cluster)
	if err != nil {
		logger.Fatal("error creating DB session", zap.String("trace", err.Error()))
	}

	return &Querier{session: session}
}

func (q *Querier) Close() {
	q.session.Close()
}