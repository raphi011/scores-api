package app

import (
	"fmt"
	"testing"
	"time"

	"github.com/raphi011/scores-api/cmd/api/auth"
	"github.com/raphi011/scores-api/cmd/api/cron"
	"github.com/raphi011/scores-api/events"
	"github.com/raphi011/scores-api/job"
	"github.com/raphi011/scores-api/repo"
	"github.com/raphi011/scores-api/repo/sql"
	"github.com/raphi011/scores-api/volleynet/sync"
	"go.uber.org/zap"
)

// WithMode sets the production mode to true if "production" is passed.
func WithMode(mode string) Option {
	return func(r *App) {
		r.production = mode == "production"
	}
}

// WithVersion sets the api version.
func WithVersion(version string) Option {
	return func(r *App) {
		r.version = version
	}
}

// WithRepository sets the repository provider and connectionstring.
func WithRepository(provider, connectionString string) Option {
	return func(r *App) {
		var err error
		var repos *repo.Repositories

		switch provider {
		case "sqlite3":
			fallthrough
		case "postgres":
			repos, err = sql.Repositories(provider, connectionString)
		default:
			err = fmt.Errorf("invalid repo provider %q", provider)
		}

		if err != nil {
			zap.S().Fatalf("Could not initialize repository: %s", err)
		}

		r.services = servicesFromRepository(repos)
	}
}

// WithOAuth sets the oauth configuration.
func WithOAuth(configPath, host string) Option {
	return func(r *App) {
		var err error
		r.conf, err = auth.GoogleOAuthConfig(configPath, host)

		if err != nil {
			zap.S().Infof("Could not read google secret: %v, continuing without google oauth\n", err)
		}
	}
}

// WithTestRepository configures the repository with the test database
// via test environment variables.
func WithTestRepository(t testing.TB) Option {
	t.Helper()

	return func(r *App) {
		repos, _ := sql.RepositoriesTest(t)

		r.services = servicesFromRepository(repos)

	}
}

// WithEventQueue configures the eventqueue.
func WithEventQueue() Option {
	return func(r *App) {
		r.eventBroker = &events.Broker{}

		// we never unsubcribe
		events, _ := r.eventBroker.Subscribe(sync.ScrapeEventsType)

		go func() {
			for event := range events {
				zap.S().Debugf("scrape event: %v", event)
			}
		}()

		// return broker
		return
	}
}

// WithCron enable cron jobs.
func WithCron() Option {
	return func(r *App) {
		ladderJob := cron.LadderJob{
			SyncService: r.services.Scrape,
			Genders:     []string{"M", "W"},
		}

		tournamentsJob := cron.TournamentsJob{
			SyncService: r.services.Scrape,
			Genders:     []string{"M", "W"},
			Leagues:     []string{"AMATEUR TOUR", "PRO TOUR", "JUNIOR TOUR"},
			Season:      time.Now().Year(),
		}

		lastYearsTournamentsJob := tournamentsJob
		lastYearsTournamentsJob.Season = lastYearsTournamentsJob.Season - 1

		r.services.JobManager.Start(
			job.Job{
				Name:        "Players",
				MaxFailures: 3,
				Interval:    1 * time.Hour,
				Do:          ladderJob.Do,
			},
			job.Job{
				Name:    "Last years tournaments",
				MaxRuns: 1, // only run once on startup
				Do:      tournamentsJob.Do,
			},
			job.Job{
				Name:        "Tournaments",
				MaxFailures: 3,
				Interval:    5 * time.Minute,
				Delay:       1 * time.Minute,
				Do:          tournamentsJob.Do,
			},
		)

	}
}
