package sqlite

import (
	"database/sql"
	"os"
	"runtime"
	"testing"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/repo"
)

func createRepositories(t *testing.T) *repo.Repositories {
	dbProvider := "sqlite3"
	connectionString := "file::memory:?_busy_timeout=5000&mode=memory"

	p := os.Getenv("TEST_DB_PROVIDER")
	c := os.Getenv("TEST_DB_CONNECTION")

	if p != "" && c != "" {
		dbProvider = p
		connectionString = c
	}

	t.Helper()

	repos, db, err := CreateTest(dbProvider, connectionString)

	if err != nil {
		t.Fatal("unable to open db")
	}

	switch dbProvider {
	case "sqlite3":
		setupSQLite(t, db)
	case "mysql":
		setupMysql(t, db)
	default:
		t.Fatal("Unsupported db provider")
	}

	return repos
}

func execMultiple(db *sql.DB, statements ...string) error {
	for _, statement := range statements {
		_, err := db.Exec(statement)

		if err != nil {
			return err
		}
	}

	return nil
}

func setupMysql(t *testing.T, db *sql.DB) {
	if runtime.GOMAXPROCS(0) > 1 {
		t.Fatal("Mysql testing can not run in parallel")
	}

	err := execMultiple(
		db,
		query("test/delete-all"),
	)

	if err != nil {
		t.Fatal(err)
	}
}

func setupSQLite(t *testing.T, db *sql.DB) {
	_, err := db.Exec(query("ddl/sqlite"))

	if err != nil {
		t.Fatal(err)
	}
}

func newMatch(s *repo.Repositories) *scores.Match {
	g, _ := s.Group.Create(&scores.Group{Name: "TestGroup"})
	u, _ := s.User.New(&scores.User{Email: "test@test.at"})
	p1, _ := s.Player.Create(&scores.Player{Name: "p1"})
	p2, _ := s.Player.Create(&scores.Player{Name: "p2"})
	p3, _ := s.Player.Create(&scores.Player{Name: "p3"})
	p4, _ := s.Player.Create(&scores.Player{Name: "p4"})
	t1, _ := s.Team.Create(&scores.Team{Name: "", Player1ID: p1.ID, Player2ID: p2.ID})
	t2, _ := s.Team.Create(&scores.Team{Name: "", Player1ID: p3.ID, Player2ID: p4.ID})

	return &scores.Match{
		Group:           g,
		Team1:           t1,
		Team2:           t2,
		ScoreTeam1:      15,
		ScoreTeam2:      13,
		CreatedByUserID: u.ID,
	}
}
