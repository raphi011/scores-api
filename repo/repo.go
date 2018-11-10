package repo

import "github.com/raphi011/scores"

type Repositories struct {
	Group      scores.GroupRepository
	Match      scores.MatchRepository
	Player     scores.PlayerRepository
	Statistic  scores.StatisticRepository
	Team       scores.TeamRepository
	User       scores.UserRepository
	Volleynet  scores.VolleynetRepository
	CloserFunc func()
}
