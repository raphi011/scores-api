package sqlite

import (
	"database/sql"
	"testing"

	"github.com/raphi011/scores"
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
	db, err := Open("file::memory:", "&mode=memory&cache=shared")

	if err != nil {
		t.Fatal("unable to create db")
	}

	Migrate(db)

	saltBytes := 16
	iterations := 10000

	s := &services{
		groupService:     &GroupService{DB: db},
		playerService:    &PlayerService{DB: db},
		userService:      &UserService{DB: db},
		teamService:      &TeamService{DB: db},
		matchService:     &MatchService{DB: db},
		statisticService: &StatisticService{DB: db},
		db:               db,
		pwService: &scores.PBKDF2PasswordService{
			SaltBytes:  saltBytes,
			Iterations: iterations,
		},
	}

	return s
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
