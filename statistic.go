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
}

type TeamStatistics []TeamStatistic
type PlayerStatistics []PlayerStatistic

type StatisticService interface {
	Team(teamID uint) (*TeamStatistic, error)
	Teams() (TeamStatistics, error)
	Player(playerID uint, filter string) (*PlayerStatistic, error)
	Players(filter string) (PlayerStatistics, error)
}
