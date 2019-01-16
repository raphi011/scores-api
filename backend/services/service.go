package services

import (
	"github.com/pkg/errors"
	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/volleynet"
)

// Service allows loading / mutation of volleynet data
type Service struct {
	TeamRepo repo.TeamRepository
	PlayerRepo repo.PlayerRepository
	TournamentRepo repo.TournamentRepository
}

// ValidGender returns true if if the passed gender string is valid
func (s *Service) ValidGender(gender string) bool {
	return gender == "M" || gender == "W"
}

// Ladder loads all players of the passed gender and with a rank > 0
func (s *Service) Ladder(gender string) ([]*volleynet.Player, error) {
	return s.PlayerRepo.Ladder(gender)
}

// GetTournaments loads all tournaments of a certain `gender`, `league` and `season`
func (s *Service) GetTournaments(seasons []int, genders, leagues []string) (
	[]*volleynet.FullTournament, error) {
	tournaments, err := s.TournamentRepo.Filter(seasons, genders, leagues)

	if err != nil {
		return nil, err
	}

	return s.addTeams(tournaments...)
}

func (s *Service) addTeams(tournaments ...*volleynet.FullTournament) ([]*volleynet.FullTournament, error) {
	var err error

	for _, t := range tournaments {
		t.Teams, err = s.TeamRepo.ByTournament(t.ID)

		if err != nil {
			return nil, errors.Wrapf(err, "adding teams of tournamentID %d", t.ID)
		}
	}

	return tournaments, nil
}

// Tournament loads a tournament and its teams
func (s *Service) Tournament(tournamentID int) (
	*volleynet.FullTournament, error) {
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
