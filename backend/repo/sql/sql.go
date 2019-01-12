package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/raphi011/scores"
	"github.com/pkg/errors"
	"github.com/gobuffalo/packr"
	log "github.com/sirupsen/logrus"
)

var queries packr.Box

func init() {
	queries = packr.NewBox("./queries")
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

func query(name string) string {
	return sqlx.Rebind(loadQuery(q))
}

type model interface {
	SetID(id int)
}

func insert(db *sqlx.DB, queryName string, entity model) (int, error) {
	var id int
	var err error

	// doesn't support `LastInsertID()`
	if db.DriverName() == "postgres" {
		err = db.NamedQuery(
			namedQuery(queryName),
			entity
		).Scan(&id)

	} else {
		var result db.Result

		result, err = db.NamedExec(
			query,
			entity
		)

		if err == nil {
			id, err = result.LastInsertId()
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
		return scores.ErrorNotFound
	}

	return err
}