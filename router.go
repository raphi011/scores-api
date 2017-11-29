package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var store = sessions.NewCookieStore([]byte("ultrasecret"))

func newRouter() *gin.Engine {
	r := gin.Default()

	r.Use(sessions.Sessions("goquestsession", store))

	r.GET("/", index)
	r.GET("/matches", matchIndex)
	r.GET("/matches/:matchID", matchShow)
	r.DELETE("/matches/:matchID", matchDelete)
	r.GET("/players", playerIndex)
	r.GET("/players/:playerID/statistic", playerStatistic)
	r.POST("/players", playerCreate)
	r.POST("/matches", matchCreate)

	r.GET("/loginRoute", loginHandler)
	r.GET("/auth", authHandler)

	return r
}
