package sqlite

import (
	"database/sql"
	scores "scores-backend"
)

type services struct {
	db               *sql.DB
	playerService    *PlayerService
	userService      *UserService
	teamService      *TeamService
	matchService     *MatchService
	statisticService *StatisticService
}

func createServices() *services {
	db, _ := Open("file::memory:?mode=memory&cache=shared")
	s := &services{
		playerService:    &PlayerService{DB: db},
		userService:      &UserService{DB: db},
		teamService:      &TeamService{DB: db},
		matchService:     &MatchService{DB: db},
		statisticService: &StatisticService{DB: db},
		db:               db,
	}

	return s
}

func newMatch(s *services) *scores.Match {
	u, _ := s.userService.Create(&scores.User{Email: "test@test.at"})
	p1, _ := s.playerService.Create(&scores.Player{Name: "p1"})
	p2, _ := s.playerService.Create(&scores.Player{Name: "p2"})
	p3, _ := s.playerService.Create(&scores.Player{Name: "p3"})
	p4, _ := s.playerService.Create(&scores.Player{Name: "p4"})
	t1, _ := s.teamService.Create(&scores.Team{Name: "", Player1ID: p1.ID, Player2ID: p2.ID})
	t2, _ := s.teamService.Create(&scores.Team{Name: "", Player1ID: p3.ID, Player2ID: p4.ID})

	return &scores.Match{
		Team1:      t1,
		Team2:      t2,
		ScoreTeam1: 15,
		ScoreTeam2: 13,
		CreatedBy:  u,
	}
}
