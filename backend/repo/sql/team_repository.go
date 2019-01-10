package sql

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/raphi011/scores/volleynet"
)

// TeamRepository implements VolleynetRepository interface
type TeamRepository struct {
	DB *sql.DB
}

// New creates a new team
func (s *TeamRepository) New(t *volleynet.TournamentTeam) error {
	_, err := s.DB.Exec(query("volleynet/insert-team"),
		t.TournamentID,
		t.Player1.ID,
		t.Player2.ID,
		t.Rank,
		t.Seed,
		t.TotalPoints,
		t.WonPoints,
		t.PrizeMoney,
		t.Deregistered,
	)

	return err
}

// NewBatch creates multiple new teams
func (s *TeamRepository) NewBatch(teams []volleynet.TournamentTeam) error {
	for _, t := range teams {
		err := s.New(&t)

		if err != nil {
			return err
		}
	}

	return nil
}

// Update updates a tournament team
func (s *TeamRepository) Update(t *volleynet.TournamentTeam) error {
	result, err := s.DB.Exec(
		query("volleynet/update-team"),
		t.Rank,
		t.Seed,
		t.TotalPoints,
		t.WonPoints,
		t.PrizeMoney,
		t.Deregistered,
		t.TournamentID,
		t.Player1.ID,
		t.Player2.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("Team not found")
	}

	return nil
}

// UpdateBatch updates tournament teams
func (s *TeamRepository) UpdateBatch(teams []volleynet.TournamentTeam) error {
	for _, t := range teams {
		if err := s.Update(&t); err != nil {
			return err
		}
	}

	return nil
}

// Delete deletes a team
func (s *TeamRepository) Delete(t *volleynet.TournamentTeam) error {
	_, err := s.DB.Exec(query("volleynet/delete-team"), t.TournamentID, t.Player1.ID, t.Player2.ID)

	return err
}

// ByTournament loads all teams of a tournament
func (s *TeamRepository) ByTournament(tournamentID int) ([]volleynet.TournamentTeam, error) {
	return scanTournamentTeams(s.DB,
		query("volleynet/select-team-by-tournament-id"),
		tournamentID)
}


func scanTournamentTeams(db *sql.DB, query string, args ...interface{}) ([]volleynet.TournamentTeam, error) {
	teams := []volleynet.TournamentTeam{}
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		team, err := scanTournamentTeam(rows)

		if err != nil {
			return nil, err
		}

		teams = append(teams, *team)
	}

	return teams, nil
}

func scanTournamentTeam(scanner scan) (*volleynet.TournamentTeam, error) {
	t := &volleynet.TournamentTeam{}
	t.Player1 = &volleynet.Player{}
	t.Player2 = &volleynet.Player{}

	err := scanner.Scan(
		&t.TournamentID,
		&t.Player1.ID,
		&t.Player1.FirstName,
		&t.Player1.LastName,
		&t.Player1.TotalPoints,
		&t.Player1.CountryUnion,
		&t.Player1.Birthday,
		&t.Player1.License,
		&t.Player1.Gender,
		&t.Player2.ID,
		&t.Player2.FirstName,
		&t.Player2.LastName,
		&t.Player2.TotalPoints,
		&t.Player2.CountryUnion,
		&t.Player2.Birthday,
		&t.Player2.License,
		&t.Player2.Gender,
		&t.Rank,
		&t.Seed,
		&t.TotalPoints,
		&t.WonPoints,
		&t.PrizeMoney,
		&t.Deregistered,
	)

	return t, err
}
