package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Logger middleware populates a logger with request specific fields
// and adds it to the context
func Logger(log logrus.FieldLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log = log.WithFields(logrus.Fields{
			"method":     c.Request.Method,
			"url":        c.Request.URL.String(),
			"ip":         ipFromRequest(c.Request),
			"user-agent": c.Request.UserAgent(),
			"request-id": uuid.New().String(),
		})

		c.Set("log", log)

		c.Next()
	}
}

func ipFromRequest(request *http.Request) string {
	if ip, ok := request.Header["X-Forwarded-For"]; ok {
		return ip[0]
	}

	return strings.Split(request.RemoteAddr, ":")[0]
}
