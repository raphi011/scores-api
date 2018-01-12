package sqlite

import (
	"database/sql"
	scores "scores-backend"
)

var _ scores.TeamService = &TeamService{}

type TeamService struct {
	DB *sql.DB
}

func TeamPlayerOrder(player1ID, player2ID uint) (uint, uint) {
	if player1ID > player2ID {
		return player2ID, player1ID
	}

	return player1ID, player2ID
}

const teamInsertSQL = `
	INSERT INTO teams (created_at, name, player1_id, player2_id)
	VALUES (CURRENT_TIMESTAMP, $1, $2, $3)
`

func (s *TeamService) Create(team *scores.Team) (*scores.Team, error) {
	_, err := s.DB.Exec(teamInsertSQL, team.Name, team.Player1ID, team.Player2ID)

	if err != nil {
		return nil, err
	}

	return team, nil
}

const teamSelectSQL = `
	SELECT created_at, name, player1_id, player2_id FROM teams
	WHERE player1_id = $1 and player2_id = $2
`

func (s *TeamService) ByPlayers(player1ID, player2ID uint) (*scores.Team, error) {
	team := &scores.Team{}

	err := s.DB.QueryRow(teamSelectSQL, player1ID, player2ID).
		Scan(&team.CreatedAt, &team.Name, &team.Player1ID, &team.Player2ID)

	if err != nil {
		return nil, err
	}

	return team, nil
}
