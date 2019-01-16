package sql

import (
	"fmt"

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

func loadQuery(db *sqlx.DB, name string) string {
	var q string
	var err error

	dbSpecificName := fmt.Sprintf("%s.%s.sql", name, db.DriverName())
	genericName := fmt.Sprintf("%s.sql", name)

	if q, err = queries.FindString(dbSpecificName); err == nil {
		return q
	} else if q, err = queries.FindString(genericName); err == nil {
		return q
	}

	log.Fatalf("could not load sql query %s: %v", name, err)

	return ""
}

func namedQuery(db *sqlx.DB, name string) string {
	return loadQuery(db, name)
}

func query(db *sqlx.DB, queryName string) string {
	return db.Rebind(loadQuery(db, queryName))
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
		namedQuery(db, queryName),
		entity,
	)

	return result, mapError(err)
}

func execMultiple(db *sqlx.DB, queryName string, entities ...interface{}) error {
	stmt, err := db.PrepareNamed(namedQuery(db, queryName))

	if err != nil {
		return mapError(err)
	}

	for _, entity := range entities {
		_, err := stmt.Exec(entity)

		if err != nil {
			return mapError(err)
		}
	}

	return nil
}

func insertSetID(db *sqlx.DB, queryName string, entity model) error {
	var id int
	var err error

	if db.DriverName() == "postgres" {
		var rows *sqlx.Rows
		// doesn't support `LastInsertID()`
		rows, err = db.NamedQuery(
			namedQuery(db, queryName),
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