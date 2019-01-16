package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/raphi011/scores/cmd/api/middleware"
)

var store = cookie.NewStore([]byte("ultrasecret"))

func initRouter(app app, services *Services) *gin.Engine {
	var router *gin.Engine
	if app.production {
		router = gin.Default()
	} else {
		gin.SetMode(gin.TestMode)
		router = gin.New()
		router.Use(gin.Recovery())
	}

	authHandler := authHandler{
		userService: services.User,
		password:    services.Password,
		conf:        app.conf,
	}

	volleynetHandler := volleynetHandler{
		volleynetService: services.Volleynet,
		userService:      services.User,
	}

	scrapeHandler := scrapeHandler{
		jobManager: services.JobManager,
	}

	infoHandler := infoHandler{}

	adminHandler := adminHandler{
		userService: services.User,
	}

	debugHandler := debugHandler{
		userService: services.User,
	}

	router.Use(sessions.Sessions("session", store), middleware.Logger(app.log), middleware.Metric())

	router.GET("/version", infoHandler.version)

	router.GET("/user-or-login", authHandler.loginRouteOrUser)
	router.GET("/auth", authHandler.googleAuthenticate)
	router.POST("/pw-auth", authHandler.passwordAuthenticate)

	auth := router.Group("/")
	auth.Use(middleware.Auth())
	auth.POST("/logout", authHandler.logout)

	auth.GET("/ladder", volleynetHandler.getLadder)
	auth.GET("/tournaments", volleynetHandler.getTournaments)
	auth.GET("/tournaments/:tournamentID", volleynetHandler.getTournament)
	auth.GET("/players/search", volleynetHandler.getSearchPlayers)
	auth.POST("/signup", volleynetHandler.postSignup)

	admin := auth.Group("/admin")
	admin.Use(middleware.Admin(services.User))

	admin.GET("/users", adminHandler.getUsers)
	admin.POST("/users", adminHandler.postUser)

	if !app.production {
		debug := router.Group("/debug")
		debug.Use(middleware.LocalhostOnly())

		debug.POST("/new-admin", debugHandler.postCreateAdmin)
	}

	volleynetAdmin := admin.Group("/volleynet")

	volleynetAdmin.GET("/scrape/report", scrapeHandler.report)
	volleynetAdmin.POST("/scrape/run", scrapeHandler.run)
	volleynetAdmin.POST("/scrape/stop", scrapeHandler.stop)

	return router
}
