package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raphi011/scores-api/services"
)

// Admin middleware restricts routes for admins only
func Admin(userService *services.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user-id").(*uuid.UUID)

		if userID == nil || !userService.HasRole(*userID, "admin") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
