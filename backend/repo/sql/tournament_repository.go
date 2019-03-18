package sql

import (
	"sort"

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
	seasons []string,
	leagues []string,
	gender []string) ([]*volleynet.Tournament, error) {

	tournaments := []*volleynet.Tournament{}
	err := crud.ReadIn(s.DB, "tournament/select-by-filter", &tournaments,
		seasons,
		leagues,
		gender,
	)

	return tournaments, errors.Wrap(err, "filtered tournaments")
}

// Seasons returns all available seasons
func (s *tournamentRepository) Seasons() ([]string, error) {
	seasons := []string{}
	err := crud.ReadIn(s.DB, "tournament/select-seasons", &seasons)

	sort.Strings(seasons)

	return seasons, errors.Wrap(err, "available seasons")
}

// Leagues returns all available leagues
func (s *tournamentRepository) Leagues() ([]string, error) {
	leagues := []string{}

	err := crud.ReadIn(s.DB, "tournament/select-leagues", &leagues)

	sort.Strings(leagues)

	return leagues, errors.Wrap(err, "leagues")
}

// SubLeagues returns all available sub-leagues
func (s *tournamentRepository) SubLeagues() ([]string, error) {
	subLeagues := []string{}
	err := crud.ReadIn(s.DB, "tournament/select-sub-leagues", &subLeagues)

	sort.Strings(subLeagues)

	return subLeagues, errors.Wrap(err, "subleagues")
}
