package scores

import (
	"time"

	"github.com/pkg/errors"
	"github.com/raphi011/scores/volleynet"
)

type VolleynetService struct {
	Repository VolleynetRepository
}

func (s *VolleynetService) ValidGender(gender string) bool {
	return gender == "M" || gender == "W"
}

func (s *VolleynetService) Ladder(gender string) ([]volleynet.Player, error) {
	return s.Repository.Ladder(gender)
}

func (s *VolleynetService) GetTournamentsUpdatedSince(since time.Time) (
	[]*volleynet.FullTournament, error) {
	tournaments, err := s.Repository.TournamentsUpdatedSince(since)

	if err != nil {
		return nil, errors.Wrapf(err, "loading tournaments updated since %s", since)
	}

	return s.addTeams(tournaments...)
}

func (s *VolleynetService) GetTournaments(gender, league string, season int) (
	[]*volleynet.FullTournament, error) {
	tournaments, err := s.Repository.GetTournaments(gender, league, season)

	if err != nil {
		return nil, err
	}

	return s.addTeams(tournaments...)
}

func (s *VolleynetService) addTeams(tournaments ...*volleynet.FullTournament) ([]*volleynet.FullTournament, error) {
	var err error

	for _, t := range tournaments {
		t.Teams, err = s.Repository.TournamentTeams(t.ID)

		if err != nil {
			return nil, errors.Wrapf(err, "adding teams of tournamentID %d", t.ID)
		}
	}

	return tournaments, nil
}

func (s *VolleynetService) Tournament(tournamentID int) (
	*volleynet.FullTournament, error) {
	tournament, err := s.Repository.Tournament(tournamentID)

	if err != nil {
		return nil, err
	}

	result, err := s.addTeams(tournament)

	if err == nil {
		return result[0], nil
	}

	return nil, err
}
