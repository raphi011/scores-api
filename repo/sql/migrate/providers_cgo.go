// +build cgo

package migrate

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/jmoiron/sqlx"
)

func migrationDriver(provider string, db *sqlx.DB) (database.Driver, error) {
	switch provider {
	case "postgres":
		return postgres.WithInstance(db.DB, &postgres.Config{})
	case "mysql":
		return mysql.WithInstance(db.DB, &mysql.Config{})
	case "sqlite3":
		return sqlite3.WithInstance(db.DB, &sqlite3.Config{})
	default:
		return nil, fmt.Errorf("invalid migration db provider: %s", provider)
	}
}
