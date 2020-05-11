package crud

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/raphi011/scores"
)

// Create creates a new entity.
func Create(db *sqlx.DB, queryName string, entities ...scores.Tracked) error {
	var err error

	stmt, err := db.PrepareNamed(namedQuery(db, queryName))

	if err != nil {
		return mapError(err)
	}

	now := time.Now()

	for _, entity := range entities {
		setTrackedCreate(entity, now)

		_, err = stmt.Exec(entity)

		if err != nil {
			return mapError(err)
		}
	}

	return nil
}

// CreateSetID creates a new entity and sets the newly assigned primary key.
func CreateSetID(db *sqlx.DB, queryName string, entities ...scores.Model) error {
	var err error

	if db.DriverName() == "postgres" {
		err = createQueryID(db, queryName, entities...)
	} else {
		err = createResultID(db, queryName, entities...)
	}

	return mapError(err)
}

func createResultID(db *sqlx.DB, queryName string, entities ...scores.Model) error {
	q := query(db, queryName)

	stmt, err := db.PrepareNamed(q)

	if err != nil {
		return err
	}

	now := time.Now()

	for _, entity := range entities {
		setTrackedCreate(entity, now)
		result, err := stmt.Exec(entity)

		if err != nil {
			return err
		}

		bigID, err := result.LastInsertId()

		if err != nil {
			return err
		}

		entity.SetID(int(bigID))
	}

	return nil
}

func createQueryID(db *sqlx.DB, queryName string, entities ...scores.Model) error {
	q := namedQuery(db, queryName)

	stmt, err := db.PrepareNamed(q)

	if err != nil {
		return err
	}

	now := time.Now()

	for _, entity := range entities {
		setTrackedCreate(entity, now)
		rows, err := stmt.Query(entity)

		if err != nil {
			return err
		}

		defer rows.Close()

		var id int

		if rows.Next() {
			err = rows.Scan(&id)
		}

		if err != nil {
			return err
		}

		entity.SetID(id)
	}

	return nil
}

func setTrackedCreate(entity interface{}, now time.Time) {
	if tracked, ok := entity.(scores.Tracked); ok {
		tracked.Create(now)
	}
}
