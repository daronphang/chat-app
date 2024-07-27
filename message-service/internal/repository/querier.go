package repository

import (
	"message-service/internal/config"

	"github.com/gocql/gocql"
)

type Querier struct {
	session *gocql.Session
}

func New(cfg *config.Config) (*Querier, error) {
	cluster := provideCluster(cfg)
	session, err := provideSession(cluster)
	if err != nil {
		return nil, err
	}
	return &Querier{session: session}, nil
}

func (q *Querier) Close() {
	q.session.Close()
}