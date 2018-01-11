package scores

type Statistic struct {
	ID           int     `json:"playerId"`
	Pname        string  `json:"name"`
	Played       int     `json:"played"`
	Wongames     int     `json:"gamesWon"`
	Lostgames    int     `json:"gamesLost"`
	Wonpoints    int     `json:"pointsWon"`
	Lost         int     `json:"pointsLost"`
	Percentage   float32 `json:"percentageWon"`
	Profileimage string  `json:"profileImage"`
}

type Statistics []Statistic

type StatisticService interface {
	Team(teamID uint) (*Statistic, error)
	Teams() (*Statistics, error)
	Player(playerID uint) (*Statistic, error)
	Players(filter string) (*Statistics, error)
}
