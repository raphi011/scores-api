package scores

import (
	"github.com/pkg/errors"
	"github.com/raphi011/scores/volleynet"
)

type VolleynetService struct {
	Repository VolleynetRepository
}

func (s *VolleynetService) Ladder(gender string) ([]volleynet.Player, error) {
	return s.Repository.Ladder(gender)
}

func (s *VolleynetService) GetTournaments(gender, league string, season int) (
	[]volleynet.FullTournament, error) {
	tournaments, err := s.Repository.GetTournaments(gender, league, season)

	if err != nil {
		return nil, err
	}

	for _, t := range tournaments {
		t.Teams, err = s.Repository.TournamentTeams(t.ID)

		if err != nil {
			return nil, errors.Wrapf(err, "could teams of tournament %d", t.ID)
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

	tournament.Teams, err = s.Repository.TournamentTeams(tournament.ID)

	if err != nil {
		return nil, errors.Wrapf(err, "could teams of tournament %d", tournament.ID)
	}

	return tournament, nil
}
