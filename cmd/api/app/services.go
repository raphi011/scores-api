package app

import (
	"github.com/raphi011/scores-api/job"
	"github.com/raphi011/scores-api/repo"
	"github.com/raphi011/scores-api/services"
	volleynet_client "github.com/raphi011/scores-api/volleynet/client"
	"github.com/raphi011/scores-api/volleynet/sync"
)

type handlerServices struct {
	JobManager      *job.Manager
	User            *services.User
	Volleynet       *services.Volleynet
	Scrape          *sync.Service
	Password        services.Password
	VolleynetClient volleynet_client.Client
}

func servicesFromRepository(repos *repo.Repositories) *handlerServices {
	password := &services.PBKDF2Password{
		SaltBytes:  16,
		Iterations: 10000,
	}

	metrics := services.NewMetrics()

	userService := &services.User{
		Repo:        repos.UserRepo,
		PlayerRepo:  repos.PlayerRepo,
		SettingRepo: repos.SettingRepo,
		Password:    password,
	}

	volleynetService := services.NewVolleynetService(
		repos.TeamRepo,
		repos.PlayerRepo,
		repos.TournamentRepo,
		metrics,
	)

	manager := job.NewManager()

	scrapeService := &sync.Service{
		PlayerRepo:     repos.PlayerRepo,
		TeamRepo:       repos.TeamRepo,
		TournamentRepo: repos.TournamentRepo,

		Client: volleynet_client.Default(),
	}

	s := &handlerServices{
		Scrape:     scrapeService,
		Volleynet:  volleynetService,
		Password:   password,
		User:       userService,
		JobManager: manager,
	}

	return s
}
