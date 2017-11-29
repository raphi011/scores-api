package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var store = sessions.NewCookieStore([]byte("ultrasecret"))

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

func newRouter() *gin.Engine {
	r := gin.Default()

	r.Use(sessions.Sessions("goquestsession", store))

	r.GET("/", index)
	r.GET("/matches", matchIndex)
	r.GET("/matches/:matchID", matchShow)
	r.GET("/players", playerIndex)
	r.GET("/players/:playerID/statistic", playerStatistic)

	r.GET("/loginRoute", loginHandler)
	r.GET("/auth", authHandler)
	r.POST("/logout", logoutHandler)

	auth := r.Group("/")
	auth.Use(authRequired())
	{
		auth.DELETE("/matches/:matchID", matchDelete)
		auth.POST("/players", playerCreate)
		auth.POST("/matches", matchCreate)
	}

	return r
}
