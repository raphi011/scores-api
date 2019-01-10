package sql

import (
	"database/sql"

	//revive:disable:blank-imports
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"

	"github.com/raphi011/scores/repo"
)

type scan interface {
	Scan(src ...interface{}) error
}

// CreateTest creates the Repositories struct and returns the sql.DB connection
func CreateTest(provider, connectionString string) (*repo.Repositories, *sql.DB, error) {
	return create(provider, connectionString)
}

// Create creates the Repositories struct and returns a function that
// closes the underlying db when called
func Create(provider, connectionString string) (*repo.Repositories, func(), error) {
	repo, db, err := create(provider, connectionString)

	return repo, func() { db.Close() }, err
}

func create(provider, connectionString string) (*repo.Repositories, *sql.DB, error) {
	db, err := sql.Open(provider, connectionString)

	if err != nil {
		return nil, nil, errors.Wrapf(err, "open db provider: %s conncetion: %s failed",
			provider,
			connectionString,
		)
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
