package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Get retrieves a logger from the context
func Get(ctx *gin.Context) logrus.FieldLogger {
	if log, ok := ctx.Get("log"); ok {
		return log.(logrus.FieldLogger)
	}

	return logrus.StandardLogger()
}
