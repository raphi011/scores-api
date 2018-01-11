package sqlite

import (
	"database/sql"
	scores "scores-backend"
)

var _ scores.TeamService = &TeamService{}

type TeamService struct {
	db *sql.DB
}

func (s *TeamService) ByPlayers(player1ID, player2ID uint) (*scores.Team, error) {
	if player1ID > player2ID {
		player1ID, player2ID = player2ID, player1ID
	}

	// db.Where(Team{Player1ID: player1ID, Player2ID: player2ID}).FirstOrCreate(&t)

	return nil, nil
}
