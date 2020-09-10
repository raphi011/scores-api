package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raphi011/scores-api/cmd/api/logger"
	"go.uber.org/zap"
)

// Auth middleware restricts routes for authenticated users only
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.Get(c)
		session := sessions.Default(c)

		if userID := session.Get("user-id"); userID == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		} else {
			userID := userID.(*uuid.UUID)
			c.Set("user-id", userID)

			log = log.With(zap.String("user-id", userID.String()))
			c.Set("log", log)

			c.Next()
		}
	}
}
