package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
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

	router.Use(sessions.Sessions("session", store), loggerMiddleware(app.log), metricMiddleware())

	router.GET("/version", infoHandler.version)

	router.GET("/userOrLoginRoute", authHandler.loginRouteOrUser)
	router.GET("/auth", authHandler.googleAuthenticate)
	router.POST("/pwAuth", authHandler.passwordAuthenticate)

	auth := router.Group("/")
	auth.Use(authMiddleware())
	auth.POST("/logout", authHandler.logout)

	auth.GET("/groups/:groupID/matches", groupHandler.getMatches)
	auth.GET("/groups/:groupID", groupHandler.getGroup)
	auth.GET("/groups/:groupID/players", groupHandler.getPlayers)
	auth.GET("/groups/:groupID/playerStatistics", groupHandler.getPlayerStatistics)
	auth.POST("/groups/:groupID/matches", groupHandler.postMatch)

	auth.GET("/matches/:matchID", matchHandler.getMatch)
	auth.DELETE("/matches/:matchID", matchHandler.deleteMatch)

	auth.GET("/players/:playerID", playerHandler.getPlayer)
	auth.GET("/players/:playerID/statistics", playerHandler.getStatistics)
	auth.GET("/players/:playerID/teamStatistics", playerHandler.getTeamStatistics)
	auth.GET("/players/:playerID/matches", playerHandler.getMatches)
	auth.POST("/players", playerHandler.postPlayer)

	auth.GET("/volleynet/ladder", volleynetHandler.getLadder)
	auth.GET("/volleynet/tournaments", volleynetHandler.getTournaments)
	auth.GET("/volleynet/tournaments/:tournamentID", volleynetHandler.getTournament)
	auth.GET("/volleynet/players/search", volleynetHandler.getSearchPlayers)
	auth.POST("/volleynet/signup", volleynetHandler.postSignup)

	admin := auth.Group("/admin")
	admin.Use(adminMiddlware(app.services.User))

	admin.GET("/users", adminHandler.getUsers)

	volleynetAdmin := admin.Group("/volleynet")

	volleynetAdmin.GET("/scrape/report", volleynetScrapeHandler.report)
	volleynetAdmin.POST("/scrape/run", volleynetScrapeHandler.run)
	volleynetAdmin.POST("/scrape/stop", volleynetScrapeHandler.stop)
	// volleynetAdmin.POST("/scrape/run-all", volleynetScrapeHandler.runAll)
	// volleynetAdmin.POST("/scrape/stop-all", volleynetScrapeHandler.stopAll)

	return router
}
