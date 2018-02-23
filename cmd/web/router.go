package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"scores-backend/sqlite"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var store = sessions.NewCookieStore([]byte("ultrasecret"))

func initRouter(app app) *gin.Engine {
	var router *gin.Engine
	if app.production {
		router = gin.Default()
	} else {
		gin.SetMode(gin.TestMode)
		router = gin.New()
		router.Use(gin.Recovery())
	}

	teamService := &sqlite.TeamService{DB: app.db}
	userService := &sqlite.UserService{DB: app.db}
	matchService := &sqlite.MatchService{DB: app.db}
	playerService := &sqlite.PlayerService{DB: app.db}
	statisticService := &sqlite.StatisticService{DB: app.db}

	authHandler := authHandler{userService: userService, conf: app.conf}
	playerHandler := playerHandler{playerService: playerService}
	matchHandler := matchHandler{matchService: matchService, userService: userService, playerService: playerService, teamService: teamService}
	statisticHandler := statisticHandler{statisticService: statisticService}

	router.Use(sessions.Sessions("goquestsession", store))

	router.GET("/userOrLoginRoute", authHandler.loginRouteOrUser)
	router.GET("/auth", authHandler.authenticate)

	auth := router.Group("/")
	auth.Use(authRequired())
	{
		auth.POST("/logout", authHandler.logout)
		auth.GET("/matches", matchHandler.index)
		auth.GET("/playerMatches/:playerID", matchHandler.byPlayer)
		auth.GET("/matches/:matchID", matchHandler.matchShow)
		auth.GET("/players", playerHandler.playerIndex)
		auth.GET("/players/:playerID", playerHandler.playerShow)
		auth.GET("/statistics", statisticHandler.players)
		auth.GET("/statistics/:playerID", statisticHandler.player)
		auth.GET("/playerTeamStatistics/:playerID", statisticHandler.playerTeams)
		auth.DELETE("/matches/:matchID", matchHandler.matchDelete)
		auth.POST("/players", playerHandler.playerCreate)
		auth.POST("/matches", matchHandler.matchCreate)
	}

	return router
}

func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user-id")

		if userID == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
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
