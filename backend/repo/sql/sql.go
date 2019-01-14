package sql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/raphi011/scores"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
)

var queries *packr.Box

func init() {
	queries = packr.New("sql", "./queries")
}

func loadQuery(name string) string {
	q, err := queries.FindString(name + ".sql")

	if err != nil {
		log.Fatalf("could not load sql query %s", name)
	}

	return q
}

func namedQuery(name string) string {
	return loadQuery(name)
}

func query(db *sqlx.DB, queryName string) string {
	return db.Rebind(loadQuery(queryName))
}

type model interface {
	SetID(id int)
}

func update(db *sqlx.DB, queryName string, entity interface{}) error {
	result, err := exec(db, queryName, entity)

	if err != nil {
		return mapError(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return scores.ErrNotFound
	}

	return nil
}

func exec(db *sqlx.DB, queryName string, entity interface{}) (sql.Result, error) {
	result, err := db.NamedExec(
		namedQuery(queryName),
		entity,
	)

	return result, mapError(err)
}

func insertSetID(db *sqlx.DB, queryName string, entity model) error {
	var id int
	var err error

	if db.DriverName() == "postgres" {
		var rows *sqlx.Rows
		// doesn't support `LastInsertID()`
		rows, err = db.NamedQuery(
			namedQuery(queryName),
			entity)

		if err != nil {
			return mapError(nil)
		}

		defer rows.Close()

		if rows.Next() {
			rows.Scan(&id)
		}
	} else {
		var result sql.Result
		var bigID int64

		result, err = db.NamedExec(
			query(db, queryName),
			entity,
		)

		if err == nil {
			bigID, err = result.LastInsertId()
			id = int(bigID)
		}
	}

	if err = mapError(err); err != nil {
		return err
	}

	entity.SetID(id)

	return nil
}

func mapError(err error) error {
	if err == nil {
		return nil
	}

	if err == sql.ErrNoRows {
		return scores.ErrNotFound
	}

	return err
}