package crud

import (
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/repo"
)

func Update(db *sqlx.DB, queryName string, entities ...repo.Tracked) error {
	stmt, err := db.PrepareNamed(namedQuery(db, queryName))

	if err != nil {
		return mapError(err)
	}

	now := time.Now()

	for _, entity := range entities {
		entity.SetUpdatedAt(now)

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
