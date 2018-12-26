package sqlite

import (
	"database/sql"

	"github.com/raphi011/scores"
)

var _ scores.TeamRepository = &TeamRepository{}

// TeamRepository stores teams
type TeamRepository struct {
	DB *sql.DB
}

// TeamPlayerOrder returns the playerIDs in the order the repository expects
// (lower ID then the other)
func TeamPlayerOrder(player1ID, player2ID uint) (uint, uint) {
	if player1ID > player2ID {
		return player2ID, player1ID
	}

	return player1ID, player2ID
}

// Create creates a new team
func (s *TeamRepository) Create(team *scores.Team) (*scores.Team, error) {
	_, err := s.DB.Exec(query("team/insert"), team.Name, team.Player1ID, team.Player2ID)

	return team, err
}

// ByPlayers returns a team
func (s *TeamRepository) ByPlayers(player1ID, player2ID uint) (*scores.Team, error) {
	team := &scores.Team{}

	player1ID, player2ID = TeamPlayerOrder(player1ID, player2ID)

	err := s.DB.QueryRow(query("team/select-by-player-ids"), player1ID, player2ID).
		Scan(&team.CreatedAt, &team.Name, &team.Player1ID, &team.Player2ID)

	return team, err
}

// GetOrCreate loads a team or creates one if it doesn't exist yet
func (s *TeamRepository) GetOrCreate(player1ID, player2ID uint) (*scores.Team, error) {
	player1ID, player2ID = TeamPlayerOrder(player1ID, player2ID)

	t, err := s.ByPlayers(player1ID, player2ID)

	if err == nil {
		return t, nil
	}

	return s.Create(&scores.Team{Player1ID: player1ID, Player2ID: player2ID})
}
