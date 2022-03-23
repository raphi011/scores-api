package sql

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/raphi011/scores-api"
	"github.com/raphi011/scores-api/repo"
	"github.com/raphi011/scores-api/repo/sql/crud"
	"github.com/raphi011/scores-api/volleynet"
)

var _ repo.TeamRepository = &teamRepository{}

type teamRepository struct {
	DB *sqlx.DB
}

// New creates a new team.
func (s *teamRepository) New(t *volleynet.TournamentTeam) (*volleynet.TournamentTeam, error) {
	err := crud.Create(s.DB, "team/insert", t)

	return t, fmt.Errorf("insert team: %w", err)
}

// NewBatch creates a new team.
func (s *teamRepository) NewBatch(teams ...*volleynet.TournamentTeam) error {
	ts := make([]scores.Tracked, len(teams))

	for i, t := range teams {
		ts[i] = t
	}

	err := crud.Create(s.DB, "team/insert", ts...)

	return fmt.Errorf("batch insert team: %w", err)
}

// Update updates a tournament team.
func (s *teamRepository) Update(t *volleynet.TournamentTeam) error {
	err := crud.Update(s.DB, "team/update", t)

	return fmt.Errorf("update team: %w", err)
}

// UpdateBatch updates a tournament team.
func (s *teamRepository) UpdateBatch(teams ...*volleynet.TournamentTeam) error {
	ts := make([]scores.Tracked, len(teams))

	for i, t := range teams {
		ts[i] = t
	}

	err := crud.Update(s.DB, "team/update", ts...)

	return fmt.Errorf("batch update team: %w", err)
}

// Delete deletes a team.
func (s *teamRepository) Delete(t *volleynet.TournamentTeam) error {
	err := crud.Delete(s.DB, "team/delete", t)

	return fmt.Errorf("delete team: %w", err)
}

// ByTournament loads all teams of a tournament.
func (s *teamRepository) ByTournament(tournamentID int) (
	[]*volleynet.TournamentTeam, error) {

	teams := []*volleynet.TournamentTeam{}
	err := crud.Read(s.DB, "team/select-by-tournament-id", &teams, tournamentID)

	return teams, fmt.Errorf("byTournament team: %w", err)
}
