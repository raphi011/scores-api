package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores-backend/services"
)

// Admin middleware restricts routes for admins only
func Admin(userService *services.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user-id")

		if !userService.HasRole(userID.(int), "admin") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
