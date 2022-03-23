package migrate

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
)

// All runs all available migrations on the db connection.
func All(provider string, db *sqlx.DB) error {
	dbDriver, err := migrationDriver(provider, db)

	if err != nil {
		return fmt.Errorf("create db migration driver: %w", err)
	}

	driver, err := (&pkgerDriver{}).Open(provider)

	if err != nil {
		return fmt.Errorf("load migration scripts: %w", err)
	}

	m, err := migrate.NewWithInstance("pkger", driver, provider, dbDriver)

	if err != nil {
		return fmt.Errorf("initialize migration: %w", err)
	}

	err = m.Up()

	if err == migrate.ErrNoChange {
		return nil
	}

	return fmt.Errorf("migrateAll: %w", err)
}
