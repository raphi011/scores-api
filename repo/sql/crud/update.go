package crud

import (
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/raphi011/scores-api"
)

// Update updates multiple entities and updates the `UpdatedAt` field.
func Update(db *sqlx.DB, queryName string, entities ...scores.Tracked) error {
	stmt, err := db.PrepareNamed(namedQuery(db, queryName))

	if err != nil {
		return mapError(err)
	}

	now := time.Now()

	for _, entity := range entities {
		entity.Update(now)

		result, err := stmt.Exec(entity)

		if err != nil {
			return mapError(err)
		}

		rowsAffected, err := result.RowsAffected()

		if err != nil {
			return mapError(err)
		}

		if rowsAffected != 1 {
			return scores.ErrNotFound
		}
	}

	return nil
}
