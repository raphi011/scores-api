package scores

import "time"

type Group struct {
	Model
	Players          Players          `json:"players"`
	Matches          Matches          `json:"matches"`
	PlayerStatistics PlayerStatistics `json:"playerStatistics"`
	Name             string           `json:"name"`
	Role             string           `json:"role"`
	ImageURL         string           `json:"imageUrl"`
}

type Groups []Group

type GroupRepository interface {
	GroupsByPlayer(playerID uint) (Groups, error)
	Groups() (Groups, error)
	Group(groupID uint) (*Group, error)
	Create(*Group) (*Group, error)
	AddPlayerToGroup(playerID, groupID uint, role string) error
}

type GroupService struct {
	repo          GroupRepository
	matchRepo     MatchRepository
	playerRepo    PlayerRepository
	statisticRepo StatisticRepository
}

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

func (s *GroupService) Group(groupID uint) (*Group, error) {
	g, err := s.repo.Group(groupID)

	if err != nil {
		return g, err
	}

	g.Matches, err = s.matchRepo.GroupMatches(groupID, time.Time{}, 25)

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
