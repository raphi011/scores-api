package crud

import (
	"github.com/jmoiron/sqlx"
)

func Read(db *sqlx.DB, queryName string, dest interface{}, args ...interface{}) error {
	q := query(db, queryName)

	err := db.Select(dest, q, args...)

	return mapError(err)
}

func ReadOne(db *sqlx.DB, queryName string, dest interface{}, args ...interface{}) error {
	q := query(db, queryName)

	err := db.Get(dest, q, args...)

	return mapError(err)
}
