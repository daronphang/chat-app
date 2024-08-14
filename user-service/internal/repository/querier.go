package repository

import (
	"context"
	"user-service/internal/config"
	"user-service/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

type Querier struct {
	db DBTX
}

func New(ctx context.Context, cfg *config.Config) (*Querier, error) {
	pool, err := providePGConnPool(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &Querier{db: pool}, nil
}

func (q *Querier) withTx(tx pgx.Tx) *Querier {
	return &Querier{db: tx}
}

func (q *Querier) ExecWithTx(ctx context.Context, cb func(domain.Repository) (interface{}, error)) func() (interface{}, error) {
	// Creates closure.
	return func() (interface{}, error) {
		db := q.db.(*pgxpool.Pool)

		tx, err := db.Begin(ctx)
		if err != nil {
			return nil, err
		}
		defer tx.Rollback(ctx)

		qtx := q.withTx(tx)
		rv, err := cb(qtx)
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
	db := q.db.(*pgxpool.Pool)
	db.Close()
}