package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores/job"
)

func ScrapeHandler(jobManager *job.Manager) Scrape {
	return Scrape{
		jobManager: jobManager,
	}
}

type Scrape struct {
	jobManager *job.Manager
}

func (h *Scrape) GetReport(c *gin.Context) {
	execs := h.jobManager.Jobs()

	response(c, http.StatusOK, execs)
}

// func (h *Scrape) run(c *gin.Context) {
// 	jobName := c.Query("job")

// 	exists := h.jobManager.HasJob(jobName)

// 	if !exists {
// 		response(c, http.StatusNotFound, nil)
// 		return
// 	}

// 	err := h.jobManager.Run(jobName)

// 	if err != nil {
// 		responseErr(c, err)
// 		return
// 	}

// 	response(c, http.StatusOK, nil)
// }

// func (h *Scrape) stop(c *gin.Context) {
// 	jobName := c.Query("job")

// 	exists := h.jobManager.HasJob(jobName)

// 	if !exists {
// 		response(c, http.StatusNotFound, nil)
// 		return
// 	}

// 	err := h.jobManager.StopJob(jobName)

// 	if err != nil {
// 		responseErr(c, err)
// 		return
// 	}

// 	response(c, http.StatusOK, nil)
// }
