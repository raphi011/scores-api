package scores

import "time"

// GroupService exposes various operations on Groups
type GroupService struct {
	repo          GroupRepository
	matchRepo     MatchRepository
	playerRepo    PlayerRepository
	statisticRepo StatisticRepository
}

// NewGroupService creates a new group service
func NewGroupService(
	repo GroupRepository,
	matchRepo MatchRepository,
	playerRepo PlayerRepository,
	statisticRepo StatisticRepository,
) *GroupService {
	return &GroupService{
		repo:          repo,
		matchRepo:     matchRepo,
		playerRepo:    playerRepo,
		statisticRepo: statisticRepo,
	}
}

// Group retrieves a group with the first few matches, all its players
// and their statistics
func (s *GroupService) Group(groupID uint) (*Group, error) {
	g, err := s.repo.Get(groupID)

	if err != nil {
		return g, err
	}

	g.Matches, err = s.matchRepo.ByGroup(groupID, time.Time{}, 25)

	if err != nil {
		return g, err
	}

	g.Players, err = s.playerRepo.ByGroup(groupID)

	if err != nil {
		return g, err
	}

	g.PlayerStatistics, err = s.statisticRepo.PlayersByGroup(groupID, "all")

	if err != nil {
		return g, err
	}

	return g, nil
}
