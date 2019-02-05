package services

import (
	"time"

	"github.com/pkg/errors"
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/volleynet"
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

// TournamentFilters contains all available Tournament filters.
type TournamentFilters struct {
	Seasons []int
	Leagues []string
	Genders []string
}

// SearchTournaments loads all tournaments of a certain `gender`, `league` and `season`
func (s *Volleynet) SearchTournaments(filters TournamentFilters) (
	[]*volleynet.Tournament, error) {
	if len(filters.Seasons) == 0 {
		filters.Seasons = append(filters.Seasons, time.Now().Year())
	}
	if len(filters.Leagues) == 0 {
		// TODO read this from DB
		filters.Leagues = append(filters.Leagues, "amateur-tour")
	}
	if len(filters.Genders) == 0 {
		filters.Genders = append(filters.Genders, "M", "W")
	}

	return s.TournamentRepo.Filter(filters.Seasons, filters.Leagues, filters.Genders)
}

// Leagues loads all available Leagues as Name/Value pairs.
func (s *Volleynet) Leagues() ([]scores.NameValue, error) {
	leagues, err := s.TournamentRepo.Leagues()

	return leagues, errors.Wrap(err, "loading leagues")
}

// SubLeagues loads all available SubLeagues as Name/Value pairs.
func (s *Volleynet) SubLeagues() ([]scores.NameValue, error) {
	leagues, err := s.TournamentRepo.Leagues()

	return leagues, errors.Wrap(err, "loading leagues")
}

// Seasons loads all available seasons.
func (s *Volleynet) Seasons() ([]scores.NameValue, error) {
	leagues, err := s.TournamentRepo.Leagues()

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
