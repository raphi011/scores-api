package scores

import "github.com/raphi011/scores/volleynet/sync"

type Services struct {
	VolleynetScrape *sync.SyncService
	Group           *GroupService

	Password  *PBKDF2Password
	User      *UserService
	Match     *MatchService
	Statistic *StatisticService
	Team      *TeamService
	Player    *PlayerService
	Volleynet *VolleynetService
}
