package scores

// StatisticService allows loading of match statistics for players and teams
type StatisticService struct {
	Repository StatisticRepository
}

// Player loads match statistics of a player for a given timeframe `filter`.
func (s *StatisticService) Player(playerID uint, filter string) (*PlayerStatistic, error) {
	return s.Repository.Player(playerID, filter)
}

// Players loads match statistics of all player for a given timeframe `filter`.
func (s *StatisticService) Players(filter string) ([]PlayerStatistic, error) {
	return s.Repository.Players(filter)
}

// PlayersByGroup loads match statistics of all players in the group
// `groupID` for a given timeframe `filter`.
func (s *StatisticService) PlayersByGroup(groupID uint, filter string) ([]PlayerStatistic, error) {
	return s.Repository.PlayersByGroup(groupID, filter)
}

// PlayerTeams loads match statistics for all teammates the player `playerID` has played with.
func (s *StatisticService) PlayerTeams(playerID uint, filter string) ([]PlayerStatistic, error) {
	return s.Repository.PlayerTeams(playerID, filter)
}
