package sqlite

import (
	"database/sql"

	"github.com/raphi011/scores"
)

var _ scores.TeamRepository = &TeamRepository{}

type TeamRepository struct {
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
	VALUES (CURRENT_TIMESTAMP, ?, ?, ?)
`

func (s *TeamRepository) Create(team *scores.Team) (*scores.Team, error) {
	_, err := s.DB.Exec(teamInsertSQL, team.Name, team.Player1ID, team.Player2ID)

	return team, err
}

const teamSelectSQL = `
	SELECT created_at, name, player1_id, player2_id FROM teams
	WHERE player1_id = ? and player2_id = ?
`

func (s *TeamRepository) ByPlayers(player1ID, player2ID uint) (*scores.Team, error) {
	team := &scores.Team{}

	player1ID, player2ID = TeamPlayerOrder(player1ID, player2ID)

	err := s.DB.QueryRow(teamSelectSQL, player1ID, player2ID).
		Scan(&team.CreatedAt, &team.Name, &team.Player1ID, &team.Player2ID)

	return team, err
}

func (s *TeamRepository) GetOrCreate(player1ID, player2ID uint) (*scores.Team, error) {
	player1ID, player2ID = TeamPlayerOrder(player1ID, player2ID)

	t, err := s.ByPlayers(player1ID, player2ID)

	if err == nil {
		return t, nil
	}

	return s.Create(&scores.Team{Player1ID: player1ID, Player2ID: player2ID})
}
