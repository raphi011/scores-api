package router

import (
	"fmt"
	"net"
	"testing"
	"time"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/cenkalti/backoff"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"

	"github.com/raphi011/scores/cmd/api/auth"
	"github.com/raphi011/scores/cmd/api/cron"
	"github.com/raphi011/scores/cmd/api/middleware"
	"github.com/raphi011/scores/cmd/api/router/route"
	"github.com/raphi011/scores/events"
	"github.com/raphi011/scores/job"
	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/repo/sql"
	"github.com/raphi011/scores/services"
	"github.com/raphi011/scores/volleynet/client"
	"github.com/raphi011/scores/volleynet/sync"
)

type Router struct {
	conf        *oauth2.Config
	log         logrus.FieldLogger
	repository  *repo.Repositories
	eventBroker *events.Broker
	version     string
	production  bool
}

// Option is used to configure a new Router.
type Option func(*Router)

// New creates a new router and configures it with `opts`.
func New(opts ...Option) *Router {
	router := &Router{
		log: logrus.New(),
	}

	for _, o := range opts {
		o(router)
	}

	return router
}

func (r *Router) Build() *gin.Engine {
	var router *gin.Engine

	s := servicesFromRepository(r.repository, true, r.log)

	router = gin.New()
	router.Use(gin.Recovery())

	if !r.production {
		gin.SetMode(gin.TestMode)
	}

	authHandler := route.AuthHandler(
		s.User,
		s.Password,
		r.conf,
	)

	playerHandler := route.PlayerHandler(s.Volleynet, s.User)
	tournamentHandler := route.TournamentHandler(s.Volleynet, s.User)
	scrapeHandler := route.ScrapeHandler(s.JobManager)
	infoHandler := route.InfoHandler(r.version)
	adminHandler := route.AdminHandler(s.User)
	debugHandler := route.DebugHandler(s.User)
	cspHandler := route.CspHandler()

	// Generate keys on startup for HMAC signing + encryption.
	// This means that on every restart previously authenticated
	// users are logged out and have to login again, for now this is good enough.
	var store = cookie.NewStore(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32),
	)

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24, // 1 day
		HttpOnly: true,
		Secure:   true,
	})

	router.Use(sessions.Sessions("session", store), middleware.Logger(r.log))

	router.GET("/version", infoHandler.GetVersion)

	router.GET("/user-or-login", authHandler.GetLoginRouteOrUser)
	router.GET("/auth", authHandler.GetGoogleAuthenticate)
	router.POST("/pw-auth", authHandler.PostPasswordAuthenticate)

	auth := router.Group("/")
	auth.Use(middleware.Auth())
	auth.POST("/logout", authHandler.PostLogout)
	auth.POST("/csp-violation-report", cspHandler.PostViolationReport)

	auth.GET("/filters", tournamentHandler.GetFilterOptions)
	auth.GET("/tournaments", tournamentHandler.GetTournaments)
	auth.GET("/tournaments/:tournamentID", tournamentHandler.GetTournament)
	auth.POST("/signup", tournamentHandler.PostSignup)

	auth.GET("/ladder", playerHandler.GetLadder)
	auth.GET("/players/search", playerHandler.GetSearchPlayers)
	auth.GET("/players/partners/:playerID", playerHandler.GetPartners)
	auth.POST("/players/login", playerHandler.PostLogin)

	admin := auth.Group("/admin")
	admin.Use(middleware.Admin(s.User))

	admin.GET("/users", adminHandler.GetUsers)
	admin.POST("/users", adminHandler.PostUser)

	if !r.production {
		debug := router.Group("/debug")
		debug.Use(middleware.LocalhostOnly())

		debug.POST("/new-admin", debugHandler.PostCreateAdmin)
	}

	volleynetAdmin := admin.Group("/volleynet")

	volleynetAdmin.GET("/scrape/report", scrapeHandler.GetReport)

	return router
}

func (r *Router) Run() {
	router := r.Build()

	err := router.Run()

	if err != nil {
		r.log.Errorf("could not start router: %+v", err)
	}
}

func WithMode(mode string) Option {
	return func(r *Router) {
		r.production = mode == "production"
	}
}

func WithVersion(version string) Option {
	return func(r *Router) {
		r.version = version
	}
}

func WithRepository(provider, connectionString string) Option {
	return func(r *Router) {
		var err error

		switch provider {
		case "sqlite3":
			fallthrough
		case "postgres":
			fallthrough
		case "mysql":
			r.repository, err = sql.Repositories(provider, connectionString)
		default:
			err = fmt.Errorf("invalid repo provider %q", provider)
		}

		if err != nil {
			r.log.Fatalf("Could not initialize repository: %s", err)
		}
	}
}

func WithOAuth(configPath, host string) Option {
	return func(r *Router) {
		var err error
		r.conf, err = auth.GoogleOAuthConfig(configPath, host)

		if err != nil {
			r.log.Warnf("Could not read google secret: %v, continuing without google oauth\n", err)
		}
	}
}

func WithLogstash(logstashURL string, level logrus.Level) Option {
	return func(r *Router) {
		log := logrus.New()
		log.SetLevel(level)

		r.log = log

		if logstashURL != "" {
			var con net.Conn

			err := backoff.Retry(func() error {
				var err error

				con, err = net.Dial("tcp", logstashURL)

				if err != nil {
					log.Infof("Retrying connection to logstash: %s", err)
				}

				return err
			}, backoff.NewExponentialBackOff())

			if err != nil {
				log.Warnf("unable to setup logstash hook: %s", err)
				return
			}

			log.Info("Successfully connected to logstash")

			hook, err := logrustash.NewHookWithConn(con, "scores")

			if err != nil {
				log.Warnf("unable to setup logstash hook: %s", err)
				return
			}

			log.Hooks.Add(hook)
		}
	}
}

func WithTestRepository(t testing.TB) Option {
	t.Helper()

	return func(r *Router) {
		r.repository, _ = sql.RepositoriesTest(t)
	}
}

func WithEventQueue() Option {
	return func(r *Router) {
		r.eventBroker = &events.Broker{}

		// we never unsubcribe
		events, _ := r.eventBroker.Subscribe(sync.ScrapeEventsType)

		go func() {
			for event := range events {
				r.log.Debugf("scrape event: %v", event)
			}
		}()

		// return broker
		return
	}
}

type handlerServices struct {
	JobManager *job.Manager
	User       *services.User
	Volleynet  *services.Volleynet
	Scrape     *sync.Service
	Password   services.Password
}

func servicesFromRepository(repos *repo.Repositories, startManager bool, log logrus.FieldLogger) *handlerServices {
	password := &services.PBKDF2Password{
		SaltBytes:  16,
		Iterations: 10000,
	}

	userService := &services.User{
		Repo:        repos.UserRepo,
		PlayerRepo:  repos.PlayerRepo,
		SettingRepo: repos.SettingRepo,
		Password:    password,
	}

	volleynetService := &services.Volleynet{
		PlayerRepo:     repos.PlayerRepo,
		TeamRepo:       repos.TeamRepo,
		TournamentRepo: repos.TournamentRepo,
	}

	scrapeService := &sync.Service{
		Log: log,

		PlayerRepo:     repos.PlayerRepo,
		TeamRepo:       repos.TeamRepo,
		TournamentRepo: repos.TournamentRepo,

		Client: client.WithLogger(log),
		// Subscriptions: r.eventBroker,
	}

	manager := job.NewManager(log)

	ladderJob := cron.LadderJob{
		SyncService: scrapeService,
		Genders:     []string{"M", "W"},
	}

	tournamentsJob := cron.TournamentsJob{
		SyncService: scrapeService,
		Genders:     []string{"M", "W"},
		Leagues:     []string{"AMATEUR TOUR", "PRO TOUR", "JUNIOR TOUR"},
		Season:      time.Now().Year(),
	}

	lastYearsTournamentsJob := tournamentsJob
	lastYearsTournamentsJob.Season = lastYearsTournamentsJob.Season - 1

	if startManager {
		manager.Start(
			job.Job{
				Name:        "Players",
				MaxFailures: 3,
				Interval:    1 * time.Hour,

				Do: ladderJob.Do,
			},
			job.Job{
				Name:    "Last years tournaments",
				MaxRuns: 1, // only run once on startup

				Do: tournamentsJob.Do,
			},
			job.Job{
				Name:        "Tournaments",
				MaxFailures: 3,
				Interval:    5 * time.Minute,
				Delay:       1 * time.Minute,

				Do: tournamentsJob.Do,
			},
		)
	}

	s := &handlerServices{
		JobManager: manager,
		Scrape:     scrapeService,
		Volleynet:  volleynetService,
		Password:   password,
		User:       userService,
	}

	return s
}
