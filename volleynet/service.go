package volleynet

import (
	"time"

	"github.com/pkg/errors"
)

type teamRepository interface{
	ByTournament(tournamentID int) ([]TournamentTeam, error)
}

type tournamentRepository interface{
	UpdatedSince(since time.Time) ([]*FullTournament, error)
	Filter(season int, genders, leagues []string) ([]*FullTournament, error)
	Get(tournamentID int) (*FullTournament, error)
}
type playerRepository interface{
	Ladder(gender string) ([]Player, error)
}

// Service allows loading / mutation of volleynet data
type Service struct {
	TeamRepository teamRepository
	PlayerRepository playerRepository
	TournamentRepository tournamentRepository
}

// ValidGender returns true if if the passed gender string is valid
func (s *Service) ValidGender(gender string) bool {
	return gender == "M" || gender == "W"
}

// Ladder loads all players of the passed gender and with a rank > 0
func (s *Service) Ladder(gender string) ([]Player, error) {
	return s.PlayerRepository.Ladder(gender)
}

// GetTournamentsUpdatedSince loads all tournaments that were updated after `since`
func (s *Service) GetTournamentsUpdatedSince(since time.Time) (
	[]*FullTournament, error) {
	tournaments, err := s.TournamentRepository.UpdatedSince(since)

	if err != nil {
		return nil, errors.Wrapf(err, "loading tournaments updated since %s", since)
	}

	return s.addTeams(tournaments...)
}

// GetTournaments loads all tournaments of a certain `gender`, `league` and `season`
func (s *Service) GetTournaments(season int, gender, league []string) (
	[]*FullTournament, error) {
	tournaments, err := s.TournamentRepository.Filter(season, gender, league)

	if err != nil {
		return nil, err
	}

	return s.addTeams(tournaments...)
}

func (s *Service) addTeams(tournaments ...*FullTournament) ([]*FullTournament, error) {
	var err error

	for _, t := range tournaments {
		t.Teams, err = s.TeamRepository.ByTournament(t.ID)

		if err != nil {
			return nil, errors.Wrapf(err, "adding teams of tournamentID %d", t.ID)
		}
	}

	return tournaments, nil
}

// Tournament loads a tournament and its teams
func (s *Service) Tournament(tournamentID int) (
	*FullTournament, error) {
	tournament, err := s.TournamentRepository.Get(tournamentID)

	if err != nil {
		return nil, err
	}

	result, err := s.addTeams(tournament)

	if err == nil {
		return result[0], nil
	}

	return nil, err
}
