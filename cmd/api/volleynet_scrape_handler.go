package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores/job"
)

type volleynetScrapeHandler struct {
	jobManager *job.Manager
}

func (h *volleynetScrapeHandler) report(c *gin.Context) {
	execs := h.jobManager.Executions()

	response(c, http.StatusOK, execs)
}

func (h *volleynetScrapeHandler) run(c *gin.Context) {
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

func (h *volleynetScrapeHandler) stop(c *gin.Context) {
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

func (h *volleynetScrapeHandler) runAll(c *gin.Context) {

}

func (h *volleynetScrapeHandler) stopAll(c *gin.Context) {

}
