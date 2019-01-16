package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores/services"
	"github.com/raphi011/scores/cmd/api/logger"
)

// Admin middleware restricts routes for admins only
func Admin(userService *services.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.Get(c)
		session := sessions.Default(c)
		userID := session.Get("user-id")


		if !userService.HasRole(userID.(int), "admin") {
			log.Print("unauthorized")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
