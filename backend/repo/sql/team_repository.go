package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/raphi011/scores/volleynet"
)

// TeamRepository implements VolleynetRepository interface.
type TeamRepository struct {
	DB *sqlx.DB
}

// New creates a new team.
func (s *TeamRepository) New(t *volleynet.TournamentTeam) error {
	_, err := exec(s.DB,  "team/insert", t)

	return errors.Wrap(err, "insert team")
}

// NewBatch creates multiple new teams.
func (s *TeamRepository) NewBatch(teams []*volleynet.TournamentTeam) error {
	for _, t := range teams {
		err := s.New(t)

		if err != nil {
			return err
		}
	}

	return nil
}

// Update updates a tournament team.
func (s *TeamRepository) Update(t *volleynet.TournamentTeam) error {
	_, err := exec(s.DB, "team/update", t)

	return errors.Wrap(err, "update team")
}

// UpdateBatch updates tournament teams.
func (s *TeamRepository) UpdateBatch(teams []*volleynet.TournamentTeam) error {
	for _, t := range teams {
		if err := s.Update(t); err != nil {
			return err
		}
	}

	return nil
}

// Delete deletes a team.
func (s *TeamRepository) Delete(t *volleynet.TournamentTeam) error {
	_, err := exec(
		s.DB,
		"team/delete",
		struct{
			TournamentID int 
			Player1ID int 
			Player2ID int 
		}{
			t.TournamentID,
			t.Player1.ID,
			t.Player2.ID,
		},
	)

	return errors.Wrap(err, "delete team")
}

// ByTournament loads all teams of a tournament.
func (s *TeamRepository) ByTournament(tournamentID int) ([]*volleynet.TournamentTeam, error) {
	return s.scan(
		"team/select-by-tournament-id",
		tournamentID,
	)
}

func (s *TeamRepository) scan(query string, args ...interface{}) ([]*volleynet.TournamentTeam, error) {
	teams := []*volleynet.TournamentTeam{}
	rows, err := s.DB.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		t := &volleynet.TournamentTeam{}
		t.Player1 = &volleynet.Player{}
		t.Player2 = &volleynet.Player{}

		err = rows.Scan(
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

		if err != nil {
			return nil, err
		}

		teams = append(teams, t)
	}

	return teams, nil
}
