// +build portable

package migrate

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
)

func migrationDriver(provider string, db *sqlx.DB) (database.Driver, error) {
	switch provider {
	case "postgres":
		return postgres.WithInstance(db.DB, &postgres.Config{})
	case "mysql":
		return mysql.WithInstance(db.DB, &mysql.Config{})
	default:
		return nil, fmt.Errorf("invalid migration db provider: %s", provider)
	}
}
