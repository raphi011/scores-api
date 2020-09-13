package app

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/raphi011/scores-api/cmd/api/middleware"
	"github.com/raphi011/scores-api/cmd/api/route"
)

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
	{
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
	}

	admin := auth.Group("/admin")
	{
		admin.Use(middleware.Admin(s.User))

		admin.GET("/users", adminHandler.GetUsers)
		admin.POST("/users", adminHandler.PostUser)
	}

	if !r.production {
		debug := router.Group("/debug")
		debug.Use(middleware.LocalhostOnly())

		debug.POST("/new-admin", debugHandler.PostCreateAdmin)
	}

	volleynetAdmin := admin.Group("/volleynet")

	volleynetAdmin.GET("/scrape/report", scrapeHandler.GetReport)

	return router
}
