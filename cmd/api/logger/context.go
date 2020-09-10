package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Get retrieves a logger from the context
func Get(ctx *gin.Context) *zap.SugaredLogger {
	if log, ok := ctx.Get("log"); ok {
		return log.(*zap.SugaredLogger)
	}

	log, err := zap.NewProduction()

	if err != nil {
		panic("could not create logger")
	}

	return log.Sugar()
}
