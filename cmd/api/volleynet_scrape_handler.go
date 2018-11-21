package main

import (
	"net/http"
	"strconv"
	"time"

	msync "sync"

	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/volleynet/sync"
)

type volleynetScrapeHandler struct {
	volleynetRepository scores.VolleynetRepository
	userService         *scores.UserService
	syncService         *sync.SyncService

	mux msync.Mutex
}

func (h *volleynetScrapeHandler) scrapeLadder(c *gin.Context) {
	h.mux.Lock()
	defer h.mux.Unlock()

	gender := c.DefaultQuery("gender", "M")

	report, err := h.syncService.Ladder(gender)

	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	jsonn(c, http.StatusOK, report, "")
}

func (h *volleynetScrapeHandler) scrapeTournaments(c *gin.Context) {
	h.mux.Lock()
	defer h.mux.Unlock()

	gender := c.DefaultQuery("gender", "M")
	league := c.DefaultQuery("league", "AMATEUR TOUR")
	season := c.DefaultQuery("season", strconv.Itoa(time.Now().Year()))
	seasonInt, err := strconv.Atoi(season)

	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	report, err := h.syncService.Tournaments(gender, league, seasonInt)

	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	jsonn(c, http.StatusOK, report, "")
}
