package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/raphi011/scores/cmd/api/middleware"
)

var store = cookie.NewStore([]byte("ultrasecret"))

func initRouter(app app) *gin.Engine {
	var router *gin.Engine
	if app.production {
		router = gin.Default()
	} else {
		gin.SetMode(gin.TestMode)
		router = gin.New()
		router.Use(gin.Recovery())
	}

	authHandler := authHandler{
		userService: app.services.User,
		password:    app.services.Password,
		conf:        app.conf,
	}

	playerHandler := playerHandler{
		playerService:    app.services.Player,
		statisticService: app.services.Statistic,
		matchService:     app.services.Match,
	}

	matchHandler := matchHandler{
		matchService: app.services.Match,
		userService:  app.services.User,
	}

	groupHandler := groupHandler{
		service:          app.services.Group,
		playerService:    app.services.Player,
		matchService:     app.services.Match,
		statisticService: app.services.Statistic,
	}

	volleynetHandler := volleynetHandler{
		volleynetService: app.services.Volleynet,
		userService:      app.services.User,
	}

	volleynetScrapeHandler := volleynetScrapeHandler{
		jobManager: app.services.JobManager,
	}

	infoHandler := infoHandler{}

	adminHandler := adminHandler{
		userService: app.services.User,
	}

	debugHandler := debugHandler{
		userService: app.services.User,
	}

	router.Use(sessions.Sessions("session", store), middleware.Logger(app.log), middleware.Metric())

	router.GET("/version", infoHandler.version)

	router.GET("/user-or-login", authHandler.loginRouteOrUser)
	router.GET("/auth", authHandler.googleAuthenticate)
	router.POST("/pw-auth", authHandler.passwordAuthenticate)

	auth := router.Group("/")
	auth.Use(middleware.Auth())
	auth.POST("/logout", authHandler.logout)

	auth.GET("/groups/:groupID/matches", groupHandler.getMatches)
	auth.GET("/groups/:groupID", groupHandler.getGroup)
	auth.GET("/groups/:groupID/players", groupHandler.getPlayers)
	auth.GET("/groups/:groupID/player-statistics", groupHandler.getPlayerStatistics)
	auth.POST("/groups/:groupID/matches", groupHandler.postMatch)

	auth.GET("/matches/:matchID", matchHandler.getMatch)
	auth.DELETE("/matches/:matchID", matchHandler.deleteMatch)

	auth.GET("/players/:playerID", playerHandler.getPlayer)
	auth.GET("/players/:playerID/statistics", playerHandler.getStatistics)
	auth.GET("/players/:playerID/team-statistics", playerHandler.getTeamStatistics)
	auth.GET("/players/:playerID/matches", playerHandler.getMatches)
	auth.POST("/players", playerHandler.postPlayer)

	auth.GET("/volleynet/ladder", volleynetHandler.getLadder)
	auth.GET("/volleynet/tournaments", volleynetHandler.getTournaments)
	auth.GET("/volleynet/tournaments/:tournamentID", volleynetHandler.getTournament)
	auth.GET("/volleynet/players/search", volleynetHandler.getSearchPlayers)
	auth.POST("/volleynet/signup", volleynetHandler.postSignup)

	admin := auth.Group("/admin")
	admin.Use(middleware.Admin(app.services.User))

	admin.GET("/users", adminHandler.getUsers)
	admin.POST("/users", adminHandler.postUser)

	if !app.production {
		debug := router.Group("/debug")
		debug.Use(middleware.LocalhostOnly())

		debug.POST("/new-admin", debugHandler.postCreateAdmin)
	}

	volleynetAdmin := admin.Group("/volleynet")

	volleynetAdmin.GET("/scrape/report", volleynetScrapeHandler.report)
	volleynetAdmin.POST("/scrape/run", volleynetScrapeHandler.run)
	volleynetAdmin.POST("/scrape/stop", volleynetScrapeHandler.stop)

	return router
}
