package scores

import (
	"github.com/raphi011/scores/job"
	"github.com/raphi011/scores/volleynet"
	"github.com/raphi011/scores/volleynet/sync"
)

// Services contains all available services
type Services struct {
	VolleynetScrape *sync.Service

	JobManager *job.Manager
	Password   *PBKDF2Password
	User       *UserService
	Volleynet  *volleynet.Service
}
