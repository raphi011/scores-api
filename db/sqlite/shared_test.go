package sqlite

import (
	"database/sql"
	"os"
	"runtime"
	"testing"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/db/sqlite/setup"
)

type services struct {
	db               *sql.DB
	groupService     *GroupService
	playerService    *PlayerService
	userService      *UserService
	teamService      *TeamService
	matchService     *MatchService
	statisticService *StatisticService
	pwService        scores.PasswordService
}

func createServices(t *testing.T) *services {
	dbProvider := "sqlite3"
	connectionString := "file::memory:?_busy_timeout=5000&mode=memory"

	p := os.Getenv("TEST_DB_PROVIDER")
	c := os.Getenv("TEST_DB_CONNECTION")

	if p != "" && c != "" {
		dbProvider = p
		connectionString = c
	}

	t.Helper()

	db, err := Open(dbProvider, connectionString)

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

	s := &services{
		groupService:     &GroupService{DB: db},
		playerService:    &PlayerService{DB: db},
		userService:      &UserService{DB: db},
		teamService:      &TeamService{DB: db},
		matchService:     &MatchService{DB: db},
		statisticService: &StatisticService{DB: db},
		db:               db,
		pwService: &scores.PBKDF2PasswordService{
			SaltBytes:  16,
			Iterations: 10000,
		},
	}

	return s
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
		"DELETE FROM volleynet_tournament_teams",
		"DELETE FROM volleynet_players",
		"DELETE FROM volleynet_tournaments",
		"DELETE FROM matches",
		"DELETE FROM teams",
		"DELETE FROM group_players",
		"DELETE FROM groups",
		"DELETE FROM players",
		"DELETE FROM users",
		"DELETE FROM db_version",
	)

	if err != nil {
		t.Fatal(err)
	}
}

func setupSQLite(t *testing.T, db *sql.DB) {
	_, err := db.Exec(setup.SQLITE)

	if err != nil {
		t.Fatal(err)
	}
}

func newMatch(s *services) *scores.Match {
	g, _ := s.groupService.Create(&scores.Group{Name: "TestGroup"})
	u, _ := s.userService.Create(&scores.User{Email: "test@test.at"})
	p1, _ := s.playerService.Create(&scores.Player{Name: "p1"})
	p2, _ := s.playerService.Create(&scores.Player{Name: "p2"})
	p3, _ := s.playerService.Create(&scores.Player{Name: "p3"})
	p4, _ := s.playerService.Create(&scores.Player{Name: "p4"})
	t1, _ := s.teamService.Create(&scores.Team{Name: "", Player1ID: p1.ID, Player2ID: p2.ID})
	t2, _ := s.teamService.Create(&scores.Team{Name: "", Player1ID: p3.ID, Player2ID: p4.ID})

	return &scores.Match{
		Group:      g,
		Team1:      t1,
		Team2:      t2,
		ScoreTeam1: 15,
		ScoreTeam2: 13,
		CreatedBy:  u,
	}
}
