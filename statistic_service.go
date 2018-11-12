package scores

type StatisticService struct {
	Repository StatisticRepository
}

func (s *StatisticService) Player(playerID uint, filter string) (*PlayerStatistic, error) {
	return s.Repository.Player(playerID, filter)
}

func (s *StatisticService) Players(filter string) ([]PlayerStatistic, error) {
	return s.Repository.Players(filter)
}

func (s *StatisticService) PlayersByGroup(groupID uint, filter string) ([]PlayerStatistic, error) {
	return s.Repository.PlayersByGroup(groupID, filter)
}

func (s *StatisticService) PlayerTeams(playerID uint, filter string) ([]PlayerStatistic, error) {
	return s.Repository.PlayerTeams(playerID, filter)
}
