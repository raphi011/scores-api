package crud

import (
	"github.com/jmoiron/sqlx"
	"github.com/raphi011/scores/repo"
)

func Delete(db *sqlx.DB, queryName string, entities ...repo.Tracked) error {
	return nil
}