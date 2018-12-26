package scores

// PlayerStatistic contains match statistics data of a player
type PlayerStatistic struct {
	PlayerID uint    `json:"playerId"`
	Player   *Player `json:"player"`
	Statistic
}

// TeamStatistic contains match statistics data of a team
type TeamStatistic struct {
	Player1ID uint `json:"player1Id"`
	Player2ID uint `json:"player2Id"`
	Team      Team `json:"team"`
	Statistic
}

// Statistic contains match statistics data
type Statistic struct {
	Played        int     `json:"played"`
	GamesWon      int     `json:"gamesWon"`
	GamesLost     int     `json:"gamesLost"`
	PointsWon     int     `json:"pointsWon"`
	PointsLost    int     `json:"pointsLost"`
	PercentageWon float32 `json:"percentageWon"`
	Rank          string  `json:"rank"`
}

// StatisticRepository retrieves statistics
type StatisticRepository interface {
	Player(playerID uint, filter string) (*PlayerStatistic, error)
	Players(filter string) ([]PlayerStatistic, error)
	PlayersByGroup(groupID uint, filter string) ([]PlayerStatistic, error)
	PlayerTeams(playerID uint, filter string) ([]PlayerStatistic, error)
}

// CalculateRank returns the a rank depending on the win percentage
func CalculateRank(percentage int) string {
	for _, r := range ranks {
		if percentage >= r.percentage {
			return r.name
		}
	}

	return ""
}
