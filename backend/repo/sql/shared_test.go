package sql

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
)

func setupDB(t *testing.T) *sqlx.DB {
	t.Helper()

	dbProvider := "sqlite3"
	connectionString := ":memory:" // "file::memory:?_busy_timeout=5000&mode=memory"

	p := os.Getenv("TEST_DB_PROVIDER")
	c := os.Getenv("TEST_DB_CONNECTION")

	if p != "" && c != "" {
		dbProvider = p
		connectionString = c
	}

	db, err := sqlx.Open(dbProvider, connectionString)

	if err != nil {
		t.Fatalf("unable to open db: %v", err)
	}

	err = db.Ping()

	if err != nil {
		t.Fatalf("unable to connect to db: %v", err)
	}

	if err = migrateAll(dbProvider, db); err != nil {
		t.Fatalf("migration failed: %v", err)
	}
	
	return db
}

// func newMatch(s *repo.Repositories) *scores.Match {
// 	g, _ := s.Group.Create(&scores.Group{Name: "TestGroup"})
// 	u, _ := s.User.New(&scores.User{Email: "test@test.at"})
// 	p1, _ := s.Player.Create(&scores.Player{Name: "p1"})
// 	p2, _ := s.Player.Create(&scores.Player{Name: "p2"})
// 	p3, _ := s.Player.Create(&scores.Player{Name: "p3"})
// 	p4, _ := s.Player.Create(&scores.Player{Name: "p4"})
// 	t1, _ := s.Team.Create(&scores.Team{Name: "", Player1ID: p1.ID, Player2ID: p2.ID})
// 	t2, _ := s.Team.Create(&scores.Team{Name: "", Player1ID: p3.ID, Player2ID: p4.ID})

// 	return &scores.Match{
// 		Group:           g,
// 		Team1:           t1,
// 		Team2:           t2,
// 		ScoreTeam1:      15,
// 		ScoreTeam2:      13,
// 		CreatedByUserID: u.ID,
// 	}
// }
