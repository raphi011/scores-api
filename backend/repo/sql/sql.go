package sql

import (
	"time"
	"fmt"
	"os"
	"path"

	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/repo"
)

var (
	queries *packr.Box
	migrations *packr.Box
)

func init() {
	// this is necessary for testing since the working directory and
	// thus the location of the SQL files changes depending on the tests
	// that are executed
	folder := os.Getenv("TEST_SQL_FILES")

	queriesFolder := "./queries"
	migrationsFolder := "./migrations"

	if folder != "" {
		queriesFolder = path.Join(folder, queriesFolder)
		migrationsFolder = path.Join(folder, migrationsFolder)
	}

	queries = packr.New("sql", queriesFolder)
	migrations = packr.New("migrations", migrationsFolder)
}

// Repositories returns a collection of all repositories with an SQL backend.
func Repositories(provider, connectionString string) (*repo.Repositories, error) {
	db, err := sqlx.Open(provider, connectionString)

	if err != nil {
		return nil, errors.Wrap(err, "open db: %v")
	}

	err = migrateAll(provider, db)

	return &repo.Repositories{
		UserRepo: &userRepository{DB: db},
		PlayerRepo: &playerRepository{DB: db},
		TournamentRepo: &tournamentRepository{DB: db},
		TeamRepo: &teamRepository{DB: db},
	}, err
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

func insert(db *sqlx.DB, queryName string, entity interface{}) error {
	var err error

	if tracked, ok := entity.(repo.Tracked); ok {
		now := time.Now()
		tracked.SetCreatedAt(now)
		tracked.SetUpdatedAt(now)
	}

	if model, ok := entity.(repo.Model); ok {
		err = insertSetID(db, queryName, model)

		return mapError(err)
	}

	_, err = db.NamedExec(
		query(db, queryName),
		entity,
	)

	return mapError(err)
}

func insertSetID(db *sqlx.DB, queryName string, entity repo.Model) error {
	var id int

	if db.DriverName() == "postgres" {
		var rows *sqlx.Rows
		// doesn't support `LastInsertID()`
		rows, err := db.NamedQuery(
			namedQuery(db, queryName),
			entity)

		if err != nil {
			return mapError(err)
		}

		defer rows.Close()

		if rows.Next() {
			rows.Scan(&id)
		}
	} else {
		var result sql.Result
		var bigID int64

		result, err := db.NamedExec(
			query(db, queryName),
			entity,
		)

		if err != nil {
			return mapError(err)
		}

		bigID, err = result.LastInsertId()
		id = int(bigID)
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