package middleware

import (
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Logger middleware populates a logger with request specific fields
// and adds it to the context
func Logger(log logrus.FieldLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log = log.WithFields(logrus.Fields{
			"method":     c.Request.Method,
			"url":        c.Request.URL.String(),
			"ip":         c.Request.RemoteAddr,
			"user-agent": c.Request.UserAgent(),
			"request-id": uuid.New().String(),
		})

		c.Set("log", log)

		c.Next()
	}
}
