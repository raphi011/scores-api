package scores

type PlayerStatistic struct {
	PlayerID uint    `json:"playerId"`
	Player   *Player `json:"player"`
	Statistic
}

type TeamStatistic struct {
	Player1ID uint `json:"player1Id"`
	Player2ID uint `json:"player2Id"`
	Team      Team `json:"team"`
	Statistic
}

type Statistic struct {
	Played        int     `json:"played"`
	GamesWon      int     `json:"gamesWon"`
	GamesLost     int     `json:"gamesLost"`
	PointsWon     int     `json:"pointsWon"`
	PointsLost    int     `json:"pointsLost"`
	PercentageWon float32 `json:"percentageWon"`
	Rank          string  `json:"rank"`
}

type StatisticRepository interface {
	Player(playerID uint, filter string) (*PlayerStatistic, error)
	Players(filter string) ([]PlayerStatistic, error)
	PlayersByGroup(groupID uint, filter string) ([]PlayerStatistic, error)
	PlayerTeams(playerID uint, filter string) ([]PlayerStatistic, error)
}

func CalculateRank(percentage int) string {
	for _, r := range ranks {
		if percentage >= r.percentage {
			return r.name
		}
	}

	return ""
}
