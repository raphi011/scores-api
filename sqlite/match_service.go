package sqlite

import (
	"database/sql"
	scores "scores-backend"
)

var _ scores.MatchService = &MatchService{}

type MatchService struct {
	DB *sql.DB
}

func (s *MatchService) Delete(matchID uint) error {
	_, err := s.DB.Exec("UPDATE matches SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1", matchID)

	return err
}

func (s *MatchService) PlayerMatches(playerID uint) (*scores.Matches, error) {

	return nil, nil
}

func (s *MatchService) Create(match *scores.Match) error {

	// player1ID uint,
	// player2ID uint,
	// player3ID uint,
	// player4ID uint,
	// scoreTeam1 int,
	// scoreTeam2 int,
	// userEmail string) {

	// user := &User{}
	// team1 := &Team{}
	// team2 := &Team{}

	// user.GetUserByEmail(db, userEmail)
	// team1.GetTeam(db, player1ID, player2ID)
	// team2.GetTeam(db, player3ID, player4ID)

	// m.Team1 = *team1
	// m.Team2 = *team2
	// m.ScoreTeam1 = scoreTeam1
	// m.ScoreTeam2 = scoreTeam2
	// m.CreatedByID = user.ID

	// db.Create(&m)
	return nil
}

func (s *MatchService) Match(ID uint) (*scores.Match, error) {
	// db.
	// 	Preload("Team1.Player1").
	// 	Preload("Team1.Player2").
	// 	Preload("Team2.Player1").
	// 	Preload("Team2.Player2").
	// 	Preload("CreatedBy").
	// 	First(&m, ID)

	return nil, nil
}

func (s *MatchService) Matches() (*scores.Matches, error) {
	// var matches []Match

	// db.
	// 	Preload("Team1.Player1").
	// 	Preload("Team1.Player2").
	// 	Preload("Team2.Player1").
	// 	Preload("Team2.Player2").
	// 	Order("created_at desc").
	// 	Find(&matches)

	// return matches

	return nil, nil
}
