package crud

import (
	"github.com/pkg/errors"
	"github.com/jmoiron/sqlx"
)

// ReadIn reads rows into `dest` and expands the query's `IN` parameters.
func ReadIn(db *sqlx.DB, queryName string, dest interface{}, args ...interface{}) error {
	q, args, err := sqlx.In(loadQuery(db, queryName), args...)

	if err != nil {
		return errors.Wrap(err, "creating query")
	}

	q = db.Rebind(q)

	err = db.Select(dest, q, args...)

	return mapError(err)
}

// Read reads rows into `dest`.
func Read(db *sqlx.DB, queryName string, dest interface{}, args ...interface{}) error {
	q := query(db, queryName)

	err := db.Select(dest, q, args...)

	return mapError(err)
}

// ReadOne reads  one row into `dest`.
func ReadOne(db *sqlx.DB, queryName string, dest interface{}, args ...interface{}) error {
	q := query(db, queryName)

	err := db.Get(dest, q, args...)

	return mapError(err)
}
