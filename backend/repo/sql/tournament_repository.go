package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/repo/sql/crud"
	"github.com/raphi011/scores/volleynet"
)

type tournamentRepository struct {
	DB *sqlx.DB
}

var _ repo.TournamentRepository = &tournamentRepository{}

// Get loads a tournament by its id.
func (s *tournamentRepository) Get(tournamentID int) (*volleynet.Tournament, error) {
	tournament := &volleynet.Tournament{
		Teams: []*volleynet.TournamentTeam{},
	}
	err := crud.ReadOne(s.DB, "tournament/select-by-id", tournament, tournamentID)

	return tournament, errors.Wrap(err, "get tournament")
}

// New creates a new tournament.
func (s *tournamentRepository) New(t *volleynet.Tournament) (*volleynet.Tournament, error) {
	err := crud.Create(s.DB, "tournament/insert", t)

	return t, errors.Wrap(err, "insert tournament")
}

// NewBatch creates a new tournament.
func (s *tournamentRepository) NewBatch(tournaments ...*volleynet.Tournament) error {
	ts := make([]scores.Tracked, len(tournaments))

	for i, t := range tournaments {
		ts[i] = t
	}
	err := crud.Create(s.DB, "tournament/insert", ts...)

	return errors.Wrap(err, "insert tournament")
}

// Update updates a tournament.
func (s *tournamentRepository) Update(t *volleynet.Tournament) error {
	err := crud.Update(s.DB, "tournament/update", t)

	return errors.Wrap(err, "update tournament")
}

// UpdateBatch updates a tournament.
func (s *tournamentRepository) UpdateBatch(tournaments ...*volleynet.Tournament) error {
	ts := make([]scores.Tracked, len(tournaments))

	for i, t := range tournaments {
		ts[i] = t
	}
	err := crud.Update(s.DB, "tournament/update", ts...)

	return errors.Wrap(err, "update tournament")
}

// Filter loads all tournaments by season, league and gender.
func (s *tournamentRepository) Filter(
	seasons []int,
	leagues []string,
	formats []string) ([]*volleynet.Tournament, error) {

	tournaments := []*volleynet.Tournament{}
	err := crud.ReadIn(s.DB, "tournament/select-by-filter", &tournaments,
		formats,
		leagues,
		seasons,
	)

	return tournaments, errors.Wrap(err, "filtered tournaments")
}
