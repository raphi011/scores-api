package scores

import (
	"time"

	"github.com/pkg/errors"
)

// MatchService exposes various operations on Matches
type MatchService struct {
	Repository       MatchRepository
	PlayerRepository PlayerRepository
	GroupRepository  GroupRepository
	UserRepository   UserRepository
	TeamRepository   TeamRepository
}

// Get retrieves a match by its ID
func (s *MatchService) Get(ID uint) (*Match, error) {
	return s.Repository.Get(ID)
}

// ByPlayer retrieves matches of a player if they were created after the passed time
func (s *MatchService) ByPlayer(playerID uint, after time.Time, count uint) ([]Match, error) {
	return s.Repository.ByPlayer(playerID, after, count)
}

// ByGroup retrieves matches of a group if they were created after the passed time
func (s *MatchService) ByGroup(groupID uint, after time.Time, count uint) ([]Match, error) {
	return s.Repository.ByGroup(groupID, after, count)
}

// After retrieves matches if they were created after the passed time
func (s *MatchService) After(after time.Time, count uint) (
	[]Match, error) {
	return s.Repository.After(after, count)
}

// Delete deletes a match
func (s *MatchService) Delete(matchID uint) error {
	return s.Repository.Delete(matchID)
}

// Create creates a new match
func (s *MatchService) Create(newMatch *Match) (*Match, error) {
	group, err := s.GroupRepository.Get(newMatch.Group.ID)

	if err != nil {
		return nil, errors.Wrap(err, "error loading group")
	}

	team1, err := s.TeamRepository.
		GetOrCreate(newMatch.Team1.Player1ID, newMatch.Team1.Player2ID)

	if err != nil {
		return nil, errors.Wrap(err, "error loading team")
	}

	team2, err := s.TeamRepository.
		GetOrCreate(newMatch.Team2.Player1ID, newMatch.Team2.Player2ID)

	if err != nil {
		return nil, errors.Wrap(err, "error loading team")
	}

	// TODO: additional score validation
	match, err := s.Repository.Create(&Match{
		Group:           group,
		Team1:           team1,
		Team2:           team2,
		ScoreTeam1:      newMatch.ScoreTeam1,
		ScoreTeam2:      newMatch.ScoreTeam2,
		TargetScore:     newMatch.TargetScore,
		CreatedByUserID: newMatch.CreatedByUserID,
	})

	return match, errors.Wrap(err, "error persisting match")
}
