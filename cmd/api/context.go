package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func logger(ctx *gin.Context) logrus.FieldLogger {
	if log, ok := ctx.Get("log"); ok {
		return log.(logrus.FieldLogger)
	}

	return logrus.StandardLogger()
}
