package main

import (
	"flag"
	"fmt"
	"time"
	"net"
	"os"

	"golang.org/x/oauth2"

	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/cenkalti/backoff"
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/repo/sqlite"
	"github.com/raphi011/scores/job"
	"github.com/raphi011/scores/volleynet/sync"
	"github.com/raphi011/scores/volleynet/client"
	"github.com/sirupsen/logrus"
)

type app struct {
	conf       *oauth2.Config
	services   *scores.Services
	log        logrus.FieldLogger
	production bool
}

var version = "undefined"

func main() {
	dbProvider := flag.String("provider", "sqlite3", "DB Driver (sqlite3 or mysql)")
	connectionString := flag.String("connection", "./scores.db", "Path to sqlite db")
	gSecret := flag.String("goauth", "./client_secret.json", "Path to google oauth secret")
	logstashURL := flag.String("logstash", "", "logstash url")

	flag.Parse()

	log := setupLogger(*logstashURL)

	production := os.Getenv("APP_ENV") == "production"
	host := os.Getenv("BACKEND_URL")

	if host == "" {
		host = "http://localhost:3000"
	}

	services, repoClose, err := createServices(*dbProvider, *connectionString)

	if err != nil {
		log.Fatalf("Could not initialize services: %s", err)
	}

	defer repoClose()

	googleOAuth, err := googleOAuthConfig(*gSecret, host)

	if err != nil {
		log.Printf("Could not read google secret: %v, continuing without google oauth\n", err)
	}

	app := app{
		production: production,
		services:   services,
		conf:       googleOAuth,
		log:        log,
	}

	router := initRouter(app)
	router.Run()
}

func setupLogger(logstashURL string) logrus.FieldLogger {
	log := logrus.New()

	if logstashURL != "" {
		var con net.Conn

		err := backoff.Retry(func() error {
			var err error

			con, err = net.Dial("tcp", logstashURL)

			if err != nil {
				log.Printf("Retrying connection to logstash: %s", err)
			}

			return err
		}, backoff.NewExponentialBackOff())

		if err != nil {
			log.Warnf("unable to setup logstash hook: %s", err)
			return log
		}

		log.Print("Successfully connected to logstash")

		hook, err := logrustash.NewHookWithConn(con, "scores")

		if err != nil {
			log.Warnf("unable to setup logstash hook: %s", err)
			return log
		}

		log.Hooks.Add(hook)
	}

	return log
}

func createServices(provider string, connectionString string) (*scores.Services, func(), error) {
	var repos *repo.Repositories
	var services *scores.Services
	var closerFunc func()
	var err error

	switch provider {
	case "sqlite3":
		fallthrough
	case "mysql":
		repos, closerFunc, err = sqlite.Create(provider, connectionString)
	default:
		return nil, nil, fmt.Errorf("invalid repo provider %q", provider)
	}

	if err != nil {
		return nil, nil, err
	}

	password := &scores.PBKDF2Password{
		SaltBytes:  16,
		Iterations: 10000,
	}

	groupService := scores.NewGroupService(
		repos.Group,
		repos.Match,
		repos.Player,
		repos.Statistic,
	)

	userService := &scores.UserService{
		Repository:       repos.User,
		PlayerRepository: repos.Player,
		Password:         password,
	}

	matchService := &scores.MatchService{
		Repository:       repos.Match,
		PlayerRepository: repos.Player,
		GroupRepository:  repos.Group,
		UserRepository:   repos.User,
		TeamRepository:   repos.Team,
	}

	statisticService := &scores.StatisticService{
		Repository: repos.Statistic,
	}

	teamService := &scores.TeamService{
		Repository: repos.Team,
	}

	playerService := &scores.PlayerService{
		Repository: repos.Player,
	}

	volleynetService := &scores.VolleynetService{
		Repository: repos.Volleynet,
	}

	scrapeService := &sync.Service{
		Client: client.DefaultClient(),
		VolleynetRepository: repos.Volleynet,
	}

	manager := &job.Manager{}

	ladderJob := ladderJob{
		syncService: scrapeService,
		genders: []string{ "M", "W" },
	}

	tournamentsJob := tournamentsJob{
		syncService: scrapeService,
		genders: []string{ "M", "W" },
		leagues: []string{"AMATEUR TOUR", "PRO TOUR", "JUNIOR TOUR"},
	}

	manager.Start(
		job.Job{
			Name:        "Players",
			MaxFailures: 3,
			Interval:    1 * time.Hour,

			Do:          ladderJob.do,
		},
		job.Job{
			Name:        "Tournaments",
			MaxFailures: 3,
			Interval:    5 * time.Minute,
			Delay:       1 * time.Minute,

			Do:          tournamentsJob.do,
		},
	)


	services = &scores.Services{
		JobManager: manager,
		VolleynetScrape: scrapeService,
		Volleynet: volleynetService,
		Group:     groupService,
		Password:  password,
		User:      userService,
		Match:     matchService,
		Statistic: statisticService,
		Team:      teamService,
		Player:    playerService,


	}

	return services, closerFunc, nil
}
