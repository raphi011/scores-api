package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	// supported sql drivers
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"

	"github.com/raphi011/scores/repo"
)

// Create creates the Repositories struct and returns a function that
// closes the underlying db when called
func Create(provider, connectionString string) (*repo.Repositories, func(), error) {
	repo, db, err := create(provider, connectionString)

	return repo, func() { db.Close() }, err
}

func create(provider, connectionString string) (*repo.Repositories, *sqlx.DB, error) {
	db, err := sqlx.Open(provider, connectionString)

	if err != nil {
		return nil, nil, errors.Wrapf(err, "open db provider: %s", provider)
	}

	err = db.Ping()

	if err != nil  {
		return nil, nil, errors.Wrapf(err, "ping db provider: %s", provider)
	}

	return &repo.Repositories{
		Group:     &GroupRepository{DB: db},
		Match:     &MatchRepository{DB: db},
		Player:    &PlayerRepository{DB: db},
		Statistic: &StatisticRepository{DB: db},
		Team:      &TeamRepository{DB: db},
		User:      &UserRepository{DB: db},
		VolleynetPlayer: &VolleynetPlayerRepository{DB: db},
		VolleynetTeam: &VolleynetTeamRepository{DB: db},
		VolleynetTournament: &VolleynetTournamentRepository{DB: db},
	}, db, nil
}
