package sql

import (
	"github.com/pkg/errors"
	"github.com/jmoiron/sqlx"

	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/repo/sql/migrate"
)

// Repositories returns a collection of all repositories with an SQL backend.
func Repositories(provider, connectionString string) (*repo.Repositories, error) {
	db, err := sqlx.Open(provider, connectionString)

	if err != nil {
		return nil, errors.Wrap(err, "open db: %v")
	}

	err = migrate.All(provider, db)

	return &repo.Repositories{
		UserRepo: &userRepository{DB: db},
		PlayerRepo: &playerRepository{DB: db},
		TournamentRepo: &tournamentRepository{DB: db},
		TeamRepo: &teamRepository{DB: db},
	}, err
}


