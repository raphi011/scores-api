package sql

import (
	"fmt"
	"sort"

	"github.com/jmoiron/sqlx"

	"github.com/raphi011/scores-api"
	"github.com/raphi011/scores-api/repo"
	"github.com/raphi011/scores-api/repo/sql/crud"
	"github.com/raphi011/scores-api/volleynet"
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

	return tournament, fmt.Errorf("get tournament: %w", err)
}

// New creates a new tournament.
func (s *tournamentRepository) New(t *volleynet.Tournament) (*volleynet.Tournament, error) {
	err := crud.Create(s.DB, "tournament/insert", t)

	return t, fmt.Errorf("insert tournament: %w", err)
}

// NewBatch creates a new tournament.
func (s *tournamentRepository) NewBatch(tournaments ...*volleynet.Tournament) error {
	ts := make([]scores.Tracked, len(tournaments))

	for i, t := range tournaments {
		ts[i] = t
	}
	err := crud.Create(s.DB, "tournament/insert", ts...)

	return fmt.Errorf("insert tournament: %w", err)
}

// Update updates a tournament.
func (s *tournamentRepository) Update(t *volleynet.Tournament) error {
	err := crud.Update(s.DB, "tournament/update", t)

	return fmt.Errorf("update tournament: %w", err)
}

// UpdateBatch updates a tournament.
func (s *tournamentRepository) UpdateBatch(tournaments ...*volleynet.Tournament) error {
	ts := make([]scores.Tracked, len(tournaments))

	for i, t := range tournaments {
		ts[i] = t
	}
	err := crud.Update(s.DB, "tournament/update", ts...)

	return fmt.Errorf("update tournament: %w", err)
}

// Filter loads all tournaments by season, league and gender.
func (s *tournamentRepository) Search(filter repo.TournamentFilter) (
	[]*volleynet.Tournament, error) {

	tournaments := []*volleynet.Tournament{}
	err := crud.ReadIn(s.DB, "tournament/select-by-filter", &tournaments,
		filter.Seasons,
		filter.Leagues,
		filter.Genders,
	)

	return tournaments, fmt.Errorf("filtered tournaments: %w", err)
}

// Seasons returns all available seasons
func (s *tournamentRepository) Seasons() ([]string, error) {
	seasons := []string{}
	err := crud.ReadIn(s.DB, "tournament/select-seasons", &seasons)

	sort.Strings(seasons)

	return seasons, fmt.Errorf("available seasons: %w", err)
}

// Leagues returns all available leagues
func (s *tournamentRepository) Leagues() ([]string, error) {
	leagues := []string{}

	err := crud.ReadIn(s.DB, "tournament/select-leagues", &leagues)

	sort.Strings(leagues)

	return leagues, fmt.Errorf("leagues: %w", err)
}

// SubLeagues returns all available sub-leagues
func (s *tournamentRepository) SubLeagues() ([]string, error) {
	subLeagues := []string{}
	err := crud.ReadIn(s.DB, "tournament/select-sub-leagues", &subLeagues)

	sort.Strings(subLeagues)

	return subLeagues, fmt.Errorf("subleagues: %w", err)
}
