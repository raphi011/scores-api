package crud

import (
	"fmt"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/raphi011/scores"
	log "github.com/sirupsen/logrus"

	"github.com/raphi011/scores/packr"
)

var (
	queries = packr.New("queries", "./packr/sql/queries")
)

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

func mapError(err error) error {
	if err == nil {
		return nil
	}

	if err == sql.ErrNoRows {
		return scores.ErrNotFound
	}

	return err
}