package app

import (
	"fmt"
	"testing"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/raphi011/scores-api/cmd/api/auth"
	"github.com/raphi011/scores-api/cmd/api/cron"
	"github.com/raphi011/scores-api/cmd/api/middleware"
	"github.com/raphi011/scores-api/cmd/api/route"
	"github.com/raphi011/scores-api/events"
	"github.com/raphi011/scores-api/job"
	"github.com/raphi011/scores-api/repo"
	"github.com/raphi011/scores-api/repo/sql"
	"github.com/raphi011/scores-api/services"
	volleynet_client "github.com/raphi011/scores-api/volleynet/client"
	"github.com/raphi011/scores-api/volleynet/sync"
)

// App wraps all the services and configuration needed
// to serve the api.
type App struct {
	conf        *oauth2.Config
	services    *handlerServices
	eventBroker *events.Broker
	version     string
	production  bool
}

// Option is used to configure a new Router.
type Option func(*App)

// New creates a new router and configures it with `opts`.
func New(opts ...Option) *App {
	app := &App{}

	for _, o := range opts {
		o(app)
	}

	return app
}

// Build builds the routes from the configuration.
func (r *App) Build() *gin.Engine {
	var router *gin.Engine

	router = gin.New()
	router.Use(gin.Recovery())

	if r.production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.TestMode)
	}

	s := r.services

	authHandler := route.AuthHandler(
		s.User,
		s.Password,
		r.conf,
	)

	playerHandler := route.PlayerHandler(s.Volleynet, s.User)
	tournamentHandler := route.TournamentHandler(s.Volleynet, s.VolleynetClient, s.User)
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

		Secure: true,
	})

	router.Use(sessions.Sessions("session", store), middleware.Logger(r.production))

	if r.production {
		router.Handle("GET", "/metrics", gin.WrapH(promhttp.Handler()))
	}

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

// Run builds the router and opens the port.
func (r *App) Run() {
	router := r.Build()

	err := router.Run()

	if err != nil {
		zap.S().Errorf("could not start router: %+v", err)
	}
}

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
