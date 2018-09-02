package main

import (
	"net/http"
	"strconv"
	"time"

	msync "sync"

	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores/db/sqlite"
	"github.com/raphi011/scores/volleynet/client"
	"github.com/raphi011/scores/volleynet/sync"
)

type volleynetScrapeHandler struct {
	volleynetService sqlite.VolleynetService
	userService      *sqlite.UserService

	mux msync.Mutex
}

func (h *volleynetScrapeHandler) scrapeLadder(c *gin.Context) {
	h.mux.Lock()
	defer h.mux.Unlock()

	gender := c.DefaultQuery("gender", "M")
	vnClient := client.Default()

	sync := sync.SyncService{
		Client:           vnClient,
		VolleynetService: h.volleynetService,
	}

	report, err := sync.Ladder(gender)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
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
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	vnClient := client.Default()

	sync := sync.SyncService{
		Client:           vnClient,
		VolleynetService: h.volleynetService,
	}

	report, err := sync.Tournaments(gender, league, seasonInt)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	jsonn(c, http.StatusOK, report, "")
}
