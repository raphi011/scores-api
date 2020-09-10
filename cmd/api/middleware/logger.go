package middleware

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger middleware populates a logger with request specific fields
// and adds it to the context
func Logger(production bool) gin.HandlerFunc {

	return func(c *gin.Context) {
		var log *zap.Logger
		var err error

		requestID := c.GetHeader("x-request-id")

		if production {
			log, err = zap.NewProduction()
		} else {
			log, err = zap.NewDevelopment()
		}

		if err != nil {
			panic("could not create logger")
		}

		log = log.With(
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.String()),
			zap.String("ip", ipFromRequest(c.Request)),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("request-id", requestID),
		)

		c.Set("log", log)

		c.Next()
	}
}

func ipFromRequest(request *http.Request) string {
	if ip, ok := request.Header["X-Forwarded-For"]; ok {
		return ip[0]
	}

	host, _, _ := net.SplitHostPort(request.RemoteAddr)

	return host
}
