package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/cmd/api/logger"
)

func Admin(userService *scores.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.Get(c)
		session := sessions.Default(c)
		userID := session.Get("user-id").(uint)

		if !userService.HasRole(userID, "admin") {
			log.Print("unauthorized")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
