package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/raphi011/scores/cmd/api/csp"
	"github.com/raphi011/scores/cmd/api/logger"
)

type cspHandler struct{}

func (a *cspHandler) violationReportHandler(c *gin.Context) {
	report := csp.ViolationReport{}

	if err := c.ShouldBindWith(&report, binding.JSON); err != nil {
		responseBadRequest(c)
		return
	}

	logger.Get(c).WithField("report", report).Warn("received violation report")

	c.Status(http.StatusOK)
}
