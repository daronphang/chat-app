package repository

import (
	"database/sql"
	"path"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
)

// If migration is dirty, force the migration version down by 1
// and perform the migration again.
// For clean slates, force to version 0 first and then 1.

func provideDriver(conn *sql.DB) (database.Driver, error) {
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func provideMigrateInstance(d database.Driver) (*migrate.Migrate, error) {
	_, filename, _, _ := runtime.Caller(0)
	m, err := migrate.NewWithDatabaseInstance(
		path.Join("file:///", path.Dir(filename), "migration"), 
		"migration", // For logging purpose only.
		d,
	)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func migrateDB(conn *sql.DB) error {
	d, err := provideDriver(conn)
	if err != nil {
		return err
	}

	defer func() {
		if err := d.Close(); err != nil {
			logger.Error("unable to close db driver", zap.String("trace", err.Error()))
		}
	}()
	
	// Create migrate instance.
	m, err := provideMigrateInstance(d)
	if err != nil {
		return err
	}

	// Migrate.
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}