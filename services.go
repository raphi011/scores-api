package scores

import (
	"github.com/raphi011/scores/job"
	"github.com/raphi011/scores/volleynet/sync"
)

// Services contains all available services
type Services struct {
	VolleynetScrape *sync.Service
	Group           *GroupService

	JobManager *job.Manager
	Password   *PBKDF2Password
	User       *UserService
	Match      *MatchService
	Statistic  *StatisticService
	Team       *TeamService
	Player     *PlayerService
	Volleynet  *VolleynetService
}
