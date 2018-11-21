package main

import (
	"encoding/json"
	"strconv"

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
		syncService: app.services.VolleynetScrape,
		userService: app.services.User,
	}

	infoHandler := infoHandler{}

	router.Use(sessions.Sessions("session", store), loggerMiddleware(app.log))

	router.GET("/version", infoHandler.version)

	router.GET("/userOrLoginRoute", authHandler.loginRouteOrUser)
	router.GET("/auth", authHandler.googleAuthenticate)
	router.POST("/pwAuth", authHandler.passwordAuthenticate)

	localhost := router.Group("/")
	localhost.Use(localhostOnlyMiddleware())
	localhost.GET("/volleynet/scrape/tournaments", volleynetScrapeHandler.scrapeTournaments)
	localhost.GET("/volleynet/scrape/ladder", volleynetScrapeHandler.scrapeLadder)

	auth := router.Group("/")
	auth.Use(authMiddleware())
	auth.POST("/logout", authHandler.logout)

	auth.GET("/groups/:groupID/matches", groupHandler.getMatches)
	auth.GET("/groups/:groupID", groupHandler.getGroup)
	auth.GET("/groups/:groupID/players", groupHandler.getPlayers)
	auth.GET("/groups/:groupID/playerStatistics", groupHandler.getPlayerStatistics)
	auth.POST("/groups/:groupID/matches", groupHandler.postMatch)

	auth.GET("/volleynet/ladder", volleynetHandler.getLadder)
	auth.GET("/volleynet/tournaments", volleynetHandler.getAllTournaments)
	auth.GET("/volleynet/tournaments/:tournamentID", volleynetHandler.getTournament)
	auth.GET("/volleynet/players/search", volleynetHandler.getSearchPlayers)
	auth.POST("/volleynet/signup", volleynetHandler.postSignup)

	auth.GET("/matches/:matchID", matchHandler.getMatch)
	auth.DELETE("/matches/:matchID", matchHandler.deleteMatch)

	auth.GET("/players/:playerID", playerHandler.getPlayer)
	auth.GET("/players/:playerID/statistics", playerHandler.getStatistics)
	auth.GET("/players/:playerID/teamStatistics", playerHandler.getTeamStatistics)
	auth.GET("/players/:playerID/matches", playerHandler.getMatches)
	auth.POST("/players", playerHandler.postPlayer)

	return router
}

func jsonn(c *gin.Context, code int, data interface{}, message string) {
	out, _ := json.Marshal(gin.H{
		"status":  code,
		"message": message,
		"data":    data,
	})

	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Writer.Header().Set("Content-Length", strconv.Itoa(len(out)))

	c.Writer.WriteHeader(code)
	c.Writer.Write(out)
}
