package sql

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/raphi011/scores-api/repo"
	"github.com/raphi011/scores-api/repo/sql/migrate"
)

// Repositories returns a collection of all repositories with an SQL backend.
func Repositories(provider, connectionString string) (*repo.Repositories, error) {
	db, err := sqlx.Open(provider, connectionString)

	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	err = migrate.All(provider, db)

	return &repo.Repositories{
		UserRepo:       &userRepository{DB: db},
		PlayerRepo:     &playerRepository{DB: db},
		TournamentRepo: &tournamentRepository{DB: db},
		TeamRepo:       &teamRepository{DB: db},
		SettingRepo:    &settingRepository{DB: db},
	}, err
}
