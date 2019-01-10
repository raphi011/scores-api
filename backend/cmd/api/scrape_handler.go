package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores/job"
)

type scrapeHandler struct {
	jobManager *job.Manager
}

func (h *scrapeHandler) report(c *gin.Context) {
	execs := h.jobManager.Executions()

	response(c, http.StatusOK, execs)
}

func (h *scrapeHandler) run(c *gin.Context) {
	jobName := c.Query("job")

	exists := h.jobManager.HasJob(jobName)

	if !exists {
		response(c, http.StatusNotFound, nil)
		return
	}

	err := h.jobManager.Run(jobName)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, nil)
}

func (h *scrapeHandler) stop(c *gin.Context) {
	jobName := c.Query("job")

	exists := h.jobManager.HasJob(jobName)

	if !exists {
		response(c, http.StatusNotFound, nil)
		return
	}

	err := h.jobManager.StopJob(jobName)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, nil)
}
