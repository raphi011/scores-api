package main

import (
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores/db/sqlite"
	"github.com/raphi011/scores/volleynet"
)

type volleynetScrapeHandler struct {
	volleynetService sqlite.VolleynetService
	userService      *sqlite.UserService
}

func (h *volleynetScrapeHandler) scrapeLadder(c *gin.Context) {
	gender := c.DefaultQuery("gender", "M")
	client := volleynet.DefaultClient()

	sync := volleynet.SyncService{
		Client:           client,
		VolleynetService: h.volleynetService,
	}

	report, err := sync.Ladder(gender)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	jsonn(c, http.StatusOK, report, "")
}

func (h *volleynetScrapeHandler) scrapeTournament(c *gin.Context) {
	tournamentID, err := strconv.Atoi(c.Param("tournamentID"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tournament, err := h.volleynetService.Tournament(tournamentID)

	if err != nil {
		log.Warn(err)
	}

	client := volleynet.DefaultClient()
	fullTournament, err := client.ComplementTournament(tournament.Tournament)

	if err != nil {
		log.Warn(err)
	}

	syncInformation := volleynet.SyncTournaments([]volleynet.FullTournament{*tournament}, fullTournament.Tournament)
	sync := syncInformation[0]

	mergedTournament := volleynet.MergeTournament(sync.SyncType, sync.OldTournament, fullTournament)

	err = h.volleynetService.UpdateTournament(mergedTournament)

	if err != nil {
		log.Warn(err)
	}
}

func (h *volleynetScrapeHandler) scrapeTournaments(c *gin.Context) {
	gender := c.DefaultQuery("gender", "M")
	league := c.DefaultQuery("league", "AMATEUR TOUR")
	season := c.DefaultQuery("season", strconv.Itoa(time.Now().Year()))
	seasonInt, err := strconv.Atoi(season)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := volleynet.DefaultClient()

	sync := volleynet.SyncService{
		Client:           client,
		VolleynetService: h.volleynetService,
	}

	report, err := sync.Tournaments(gender, league, seasonInt)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	jsonn(c, http.StatusOK, report, "")
}
