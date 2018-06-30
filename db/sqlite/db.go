package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"

	"github.com/raphi011/scores/db/sqlite/migrations"
	"github.com/raphi011/scores/migrate"
)

type scan interface {
	Scan(src ...interface{}) error
}

func Open(filename string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filename)

	if err != nil {
		return nil, errors.Wrapf(err, "open sqlite db: %s failed", filename)
	}

	return db, nil
}

func Migrate(db *sql.DB) error {
	return migrate.Migrate(db, migrations.MigrationSet)
}

func Reset(db *sql.DB) error {
	return migrate.Reset(db, migrations.ResetSet)
}
