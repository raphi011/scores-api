package migrate

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	// supported sql drivers

	"github.com/gobuffalo/packr/v2"
)

var (
	migrations = packr.New("migrations", "../migrations")
)

// All runs all available migrations on the db connection.
func All(provider string, db *sqlx.DB) error {
	dbDriver, err := migrationDriver(provider, db)

	if err != nil {
		return errors.Wrap(err, "create db migration driver")
	}

	driver, err := (&packrDriver{}).Open(provider)

	if err != nil {
		return errors.Wrap(err, "load migration scripts")
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
