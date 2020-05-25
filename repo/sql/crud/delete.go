package crud

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/raphi011/scores-backend"
)

// Delete deletes an entity.
func Delete(db *sqlx.DB, queryName string, entities ...scores.Tracked) error {
	stmt, err := db.PrepareNamed(namedQuery(db, queryName))

	if err != nil {
		return mapError(err)
	}

	now := time.Now()

	for _, entity := range entities {
		entity.Delete(now)

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
