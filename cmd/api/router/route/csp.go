package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/raphi011/scores-backend/cmd/api/csp"
	"github.com/raphi011/scores-backend/cmd/api/logger"
)

// CspHandler is the constructor for the csp routes handler.
func CspHandler() Csp {
	return Csp{}
}

// Csp wraps the dependencies of the CspHandler.
type Csp struct{}

// PostViolationReport handles the csp violation repot route.
func (a *Csp) PostViolationReport(c *gin.Context) {
	report := csp.ViolationReport{}

	if err := c.ShouldBindWith(&report, binding.JSON); err != nil {
		responseBadRequest(c)
		return
	}

	logger.Get(c).WithField("report", report).Warn("received violation report")

	c.Status(http.StatusOK)
}
