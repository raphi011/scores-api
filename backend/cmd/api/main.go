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
	"github.com/sirupsen/logrus"

	"github.com/raphi011/scores/events"
	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/repo/sql"
	"github.com/raphi011/scores/job"
	"github.com/raphi011/scores/services"
	"github.com/raphi011/scores/volleynet/sync"
	"github.com/raphi011/scores/volleynet/client"
)

type app struct {
	conf       *oauth2.Config
	log        logrus.FieldLogger
	production bool
}

var version = "undefined"

func main() {
	dbProvider := flag.String("provider", "sqlite3", "DB Driver (sqlite3, mysql or postgres)")
	connectionString := flag.String("connection", "./scores.db", "provider specific connectionstring")
	gSecret := flag.String("gauth", "./client_secret.json", "Path to google oauth secret")
	logstashURL := flag.String("logstash", "", "logstash url")

	flag.Parse()

	log := setupLogger(*logstashURL)

	production := os.Getenv("APP_ENV") == "production"
	host := os.Getenv("BACKEND_URL")

	if host == "" {
		host = "http://localhost:3000"
	}

	services, err := createServices(*dbProvider, *connectionString)

	if err != nil {
		log.Fatalf("Could not initialize services: %s", err)
	}

	googleOAuth, err := googleOAuthConfig(*gSecret, host)

	if err != nil {
		log.Printf("Could not read google secret: %v, continuing without google oauth\n", err)
	}

	app := app{
		production: production,
		conf:       googleOAuth,
		log:        log,
	}

	router := initRouter(app, services)
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

func newBroker() *events.Broker {
	broker := &events.Broker{}

	// we never unsubcribe
	events, _ := broker.Subscribe(sync.ScrapeEventsType)

	go func() {
		for event := range events {
			// TODO log the changes properly
			fmt.Printf("scrape event: %v", event)
		}
	}()

	return broker
}

type handlerServices struct {
	JobManager *job.Manager
	User *services.User
	Volleynet *services.Service
	Scrape *sync.Service
	Password services.Password
}

func createServices(provider string, connectionString string) (*handlerServices, error) {
	var repos *repo.Repositories
	var s *handlerServices
	var err error

	switch provider {
	case "sqlite3":
		fallthrough
	case "postgres": 
		fallthrough
	case "mysql":
		repos, err = sql.Repositories(provider, connectionString)
	default:
		return nil, fmt.Errorf("invalid repo provider %q", provider)
	}

	if err != nil {
		return nil, err
	}

	password := &services.PBKDF2Password{
		SaltBytes:  16,
		Iterations: 10000,
	}

	userService := &services.User{
		Repo:       repos.UserRepo,
		PlayerRepo: repos.PlayerRepo,
		Password:         password,
	}

	volleynetService := &services.Service{
		PlayerRepo: repos.PlayerRepo,
		TeamRepo: repos.TeamRepo,
		TournamentRepo: repos.TournamentRepo,
	}

	broker := newBroker()

	scrapeService := &sync.Service{
		Client: client.DefaultClient(),
		PlayerRepo: repos.PlayerRepo,
		TeamRepo: repos.TeamRepo,
		TournamentRepo: repos.TournamentRepo,
		Subscriptions: broker,
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


	s = &handlerServices{
		JobManager: manager,
		Scrape: scrapeService,
		Volleynet: volleynetService,
		Password:  password,
		User:      userService,
	}

	return s, nil
}
