package repository

import (
	"context"
	"net/url"
	"time"
	"user-service/internal"
	"user-service/internal/config"

	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	logger, _ = internal.WireLogger()
)

// When using a connection pool, if the existing connection to db is broken, it will automatically
// perform a reconnection for every new connection a thread requests from the pool.
// Hence, it is safe to initialize the connection pool once and reusing it for all threads.
func providePGConnPool(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	addr := &url.URL{
		Scheme: "postgres",
		User: url.UserPassword(cfg.Postgres.Username, cfg.Postgres.Password),
		Host: cfg.Postgres.HostAddress,
		Path: cfg.Postgres.DBName,
	}

	// Create config.
	dbConfig, err := pgxpool.ParseConfig(addr.String())
	dbConfig.MaxConns = int32(4)
	dbConfig.MinConns = int32(0)
	dbConfig.MaxConnLifetime = time.Hour
	dbConfig.MaxConnIdleTime = time.Minute * 30
	dbConfig.HealthCheckPeriod = time.Minute
	dbConfig.ConnConfig.ConnectTimeout = time.Second * 5

	pool, err :=  pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, err
	}
	return pool, err
}

func providePGConn(cfg *config.Config, withDB bool) (*sql.DB, error) {
	var path = cfg.Postgres.DBName
	if !withDB {
		path = ""
	}
	addr := &url.URL{
		Scheme: "postgres",
		User: url.UserPassword(cfg.Postgres.Username, cfg.Postgres.Password),
		Host: cfg.Postgres.HostAddress,
		Path: path,
	}

	conn, err := sql.Open("pgx", addr.String())
	if err != nil {
		return nil, err
	}
	return conn, err
}

func SetupDB(ctx context.Context, cfg *config.Config) error {
	// Establish conn without database name.
	conn, err := providePGConn(cfg, false)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create database.
	rv := conn.QueryRow("SELECT 1 FROM pg_database WHERE datname = $1", cfg.Postgres.DBName)
	if err := rv.Scan(); err == sql.ErrNoRows {
		_, err = conn.Exec("CREATE DATABASE " + cfg.Postgres.DBName)
		if err != nil {
			return err
		}
	}

	// Establish conn with database name.
	connWithDB, err := providePGConn(cfg, true)
	if err != nil {
		return err
	}
	defer connWithDB.Close()

	// Migrate DB.
	if err := migrateDB(connWithDB); err != nil {
		return err
	}
	return nil
}