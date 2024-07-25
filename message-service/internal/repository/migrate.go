package repository

import (
	"fmt"
	"message-service/internal/config"
	"path"
	"runtime"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	c "github.com/golang-migrate/migrate/v4/database/cassandra"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
)

// If migration is dirty, force the migration version down by 1
// and perform the migration again.
// For clean slates, force to version 0 first and then 1.

func provideDriver(cfg *config.Config) (database.Driver, error) {
	// Create driver.
	// Default consistency is gocql.All; hence, only one host is needed for migration.
	var d database.Driver
	var err error
	hosts := strings.Split(cfg.Cassandra.HostAddresses, ",")
	for _, host := range hosts {
		// cassandra://host:port/keyspace?param1=value&param2=value2
		addr := fmt.Sprintf("cassandra://%v/%s", host, keyspaceMessage)
		p := &c.Cassandra{}
		d, err = p.Open(addr)
		if err == nil {
			break
		}
	}
	if d == nil {
		return nil, err
	}
	return d, nil
}

func provideMigrateInstance(d database.Driver) (*migrate.Migrate, error) {
	_, filename, _, _ := runtime.Caller(0)
	m, err := migrate.NewWithDatabaseInstance(
		path.Join("file:///", path.Dir(filename), "migration"), 
		keyspaceMessage, // For logging purpose only.
		d,
	)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func migrateDB(cfg *config.Config, logger *zap.Logger) error {
	// Create driver.
	d, err := provideDriver(cfg)
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