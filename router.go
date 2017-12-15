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

func JSONN(c *gin.Context, code int, data interface{}, message string) {
	c.JSON(code, gin.H{
		"status":  code,
		"message": message,
		"data":    data,
	})
}
