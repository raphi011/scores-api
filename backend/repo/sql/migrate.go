package sql

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"

	// supported sql drivers
    "github.com/golang-migrate/migrate/v4/database"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    "github.com/golang-migrate/migrate/v4/database/mysql"
    "github.com/golang-migrate/migrate/v4/database/sqlite3"
)

func migrateAll(provider string, db *sqlx.DB) error {
	var dbDriver database.Driver
	var err error

	driver, err := (&packrDriver{}).Open(provider)

	if err != nil {
		return errors.Wrap(err, "load migration scripts")
	}

	switch provider {
	case "postgres": dbDriver, err = postgres.WithInstance(db.DB, &postgres.Config{})
	case "mysql": dbDriver, err = mysql.WithInstance(db.DB, &mysql.Config{})
	case "sqlite3": dbDriver, err = sqlite3.WithInstance(db.DB, &sqlite3.Config{})
	default: return fmt.Errorf("invalid migration db provider: %s", provider)
	}

	m, err := migrate.NewWithInstance("packr", driver, provider, dbDriver)

	if err != nil {
		return errors.Wrap(err, "initialize migration")
	}

	err = m.Up()

	if err == migrate.ErrNoChange {
		return nil
	}

	return errors.Wrap(err, "migrateAll")
}
