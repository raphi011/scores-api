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
	jsonn(c, http.StatusOK, h.jobManager.Executions(), "")
}

func (h *volleynetScrapeHandler) run(c *gin.Context) {
	jobName := c.Query("job")

	exists := h.jobManager.HasJob(jobName)

	if !exists {
		jsonn(c, http.StatusNotFound, nil, "Not Found")
	}

	err := h.jobManager.Run(jobName)

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad Request")
	}

	jsonn(c, http.StatusOK, nil, "")
}

func (h *volleynetScrapeHandler) stop(c *gin.Context) {
	jobName := c.Query("job")

	exists := h.jobManager.HasJob(jobName)

	if !exists {
		jsonn(c, http.StatusNotFound, nil, "Not Found")
	}

	err := h.jobManager.StopJob(jobName)

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad Request")
	}

	jsonn(c, http.StatusOK, nil, "")

}

func (h *volleynetScrapeHandler) runAll(c *gin.Context) {

}

func (h *volleynetScrapeHandler) stopAll(c *gin.Context) {

}
