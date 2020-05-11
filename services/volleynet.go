package services

import (
	"github.com/pkg/errors"
	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/volleynet"

	"strconv"
	"time"
)

// Volleynet allows loading / mutation of volleynet data
type Volleynet struct {
	TeamRepo       repo.TeamRepository
	PlayerRepo     repo.PlayerRepository
	TournamentRepo repo.TournamentRepository
}

// ValidGender returns true if if the passed gender string is valid
func (s *Volleynet) ValidGender(gender string) bool {
	return gender == "M" || gender == "W"
}

// Ladder loads all players of the passed gender and with a rank > 0
func (s *Volleynet) Ladder(gender string) ([]*volleynet.Player, error) {
	return s.PlayerRepo.Ladder(gender)
}

// FilterOptions are the available tournament filters.
type FilterOptions struct {
	Seasons []string `json:"seasons"`
	Leagues []string `json:"leagues"`
	Genders []string `json:"genders"`
}

// SearchTournaments searches for tournaments that satisfy the passed filter.
func (s *Volleynet) SearchTournaments(filter repo.TournamentFilter) (
	[]*volleynet.Tournament, error) {
	return s.TournamentRepo.Search(filter)
}

// SearchPlayers searches for players that satisfy the passed filter.
func (s *Volleynet) SearchPlayers(filter repo.PlayerFilter) (
	[]*volleynet.Player, error) {
	return s.PlayerRepo.Search(filter)
}

// SetDefaultFilters sets filters to the users's previous setting - or the default value
// if no value has been provided for a filter.
func (s *Volleynet) SetDefaultFilters(filter repo.TournamentFilter) repo.TournamentFilter {
	if len(filter.Seasons) == 0 {
		filter.Seasons = append(filter.Seasons, strconv.Itoa(time.Now().Year()))
	}
	if len(filter.Leagues) == 0 {
		// TODO read this from DB
		filter.Leagues = append(filter.Leagues, "amateur-tour")
	}
	if len(filter.Genders) == 0 {
		filter.Genders = append(filter.Genders, "M", "W")
	}

	return filter
}

// TournamentFilterOptions returns available filter options
func (s *Volleynet) TournamentFilterOptions() (*FilterOptions, error) {
	leagues, err := s.Leagues()
	if err != nil {
		return nil, errors.Wrap(err, "loading leagues")
	}

	seasons, err := s.Seasons()

	if err != nil {
		return nil, errors.Wrap(err, "loading seasons")
	}

	options := &FilterOptions{
		Genders: []string{"M", "W"},
		Leagues: leagues,
		Seasons: seasons,
	}

	return options, nil
}

// Leagues loads all available Leagues as Name/Value pairs.
func (s *Volleynet) Leagues() ([]string, error) {
	leagues, err := s.TournamentRepo.Leagues()

	return leagues, errors.Wrap(err, "loading leagues")
}

// SubLeagues loads all available SubLeagues as Name/Value pairs.
func (s *Volleynet) SubLeagues() ([]string, error) {
	leagues, err := s.TournamentRepo.Leagues()

	return leagues, errors.Wrap(err, "loading leagues")
}

// PreviousPartners returns a list of all partners a player has played with before.
func (s *Volleynet) PreviousPartners(playerID int) ([]*volleynet.Player, error) {
	partners, err := s.PlayerRepo.PreviousPartners(playerID)

	return partners, errors.Wrap(err, "loading parners")
}

// Seasons loads all available seasons.
func (s *Volleynet) Seasons() ([]string, error) {
	leagues, err := s.TournamentRepo.Seasons()

	return leagues, errors.Wrap(err, "loading leagues")
}

func (s *Volleynet) addTeams(tournaments ...*volleynet.Tournament) ([]*volleynet.Tournament, error) {
	var err error

	for _, t := range tournaments {
		t.Teams, err = s.TeamRepo.ByTournament(t.ID)

		if err != nil {
			return nil, errors.Wrapf(err, "adding teams of tournamentID %d", t.ID)
		}
	}

	return tournaments, nil
}

// TournamentInfo loads a tournament and its teams
func (s *Volleynet) TournamentInfo(tournamentID int) (
	*volleynet.Tournament, error) {
	tournament, err := s.TournamentRepo.Get(tournamentID)

	if err != nil {
		return nil, err
	}

	result, err := s.addTeams(tournament)

	if err == nil {
		return result[0], nil
	}

	return nil, err
}
