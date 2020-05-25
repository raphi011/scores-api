package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores-backend/cmd/api/logger"
)

// Auth middleware restricts routes for authenticated users only
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.Get(c)
		session := sessions.Default(c)
		userID := session.Get("user-id")

		if userID == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user-id", userID)

		log = log.WithFields(logrus.Fields{"user-id": userID})

		c.Set("log", log)

		c.Next()
	}
}
