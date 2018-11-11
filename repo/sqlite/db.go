package sqlite

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"

	"github.com/raphi011/scores/migrate"
	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/repo/sqlite/migrations"
)

type scan interface {
	Scan(src ...interface{}) error
}

func CreateTest(provider, connectionString string) (*repo.Repositories, *sql.DB, error) {
	return create(provider, connectionString)
}

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
		Volleynet: &VolleynetRepository{DB: db},
	}, db, nil
}

func Migrate(db *sql.DB) error {
	return migrate.Migrate(db, migrations.MigrationSet)
}

func Reset(db *sql.DB) error {
	return migrate.Reset(db, migrations.ResetSet)
}
