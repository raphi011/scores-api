package crud

import (
	"github.com/pkg/errors"
	"github.com/jmoiron/sqlx"
)

func ReadIn(db *sqlx.DB, queryName string, dest interface{}, args ...interface{}) error {
	q, args, err := sqlx.In(loadQuery(db, queryName), args...)

	if err != nil {
		return errors.Wrap(err, "creating query")
	}

	q = db.Rebind(q)

	err = db.Select(dest, q, args...)

	return mapError(err)
}

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
