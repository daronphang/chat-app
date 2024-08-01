package repository

import (
	"context"
	"user-service/internal/config"
	"user-service/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Querier struct {
	db *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.Config) (*Querier, error) {
	pool, err := providePGConnPool(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &Querier{db: pool}, nil
}

func (q *Querier) ExecWithTx(ctx context.Context, cb func(domain.Repository) (interface{}, error)) func() (interface{}, error) {
	// Creates closure.
	return func() (interface{}, error) {
		tx, err := q.db.Begin(ctx)
		if err != nil {
			return nil, err
		}
		defer tx.Rollback(ctx)

		rv, err := cb(q)
		if err != nil {
			return nil, err
		}
		if err := tx.Commit(ctx); err != nil {
			return nil, err
		}
		return rv, nil
	}
}

func (q *Querier) Close() {
	q.db.Close()
}