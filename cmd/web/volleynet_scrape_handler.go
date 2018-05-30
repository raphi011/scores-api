package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raphi011/scores/sqlite"
	"github.com/raphi011/scores/volleynet"
)

type volleynetScrapeHandler struct {
	volleynetService *sqlite.VolleynetService
	userService      *sqlite.UserService
}

func (h *volleynetScrapeHandler) scrapeLadder(c *gin.Context) {
	gender := c.DefaultQuery("gender", "M")

	client := volleynet.DefaultClient()
	ranks, err := client.Ladder(gender)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	err = h.volleynetService.SyncPlayers(gender, ranks)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Status(http.StatusOK)
}

func (h *volleynetScrapeHandler) scrapeTournaments(c *gin.Context) {
	gender := c.DefaultQuery("gender", "M")
	league := c.DefaultQuery("league", "AMATEUR TOUR")
	season := c.DefaultQuery("season", strconv.Itoa(time.Now().Year()))

	seasonNumber, err := strconv.Atoi(season)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	client := volleynet.DefaultClient()

	// get list of tournaments
	tournaments, err := client.AllTournaments(gender, league, season)

	if err != nil {
		// return early
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	// find out which have to be updated
	syncInformation, err := h.volleynetService.SyncTournamentInformation(tournaments...)

	// update one after another
	for _, t := range syncInformation {
		link := client.GetApiTournamentLink(&t.Tournament)
		fullTournament, err := client.GetTournament(t.ID, link)
		fullTournament.Gender = gender
		fullTournament.League = league
		fullTournament.Season = seasonNumber

		if err != nil {
			c.AbortWithError(http.StatusServiceUnavailable, err)
			return
		}

		if t.New {
			log.Printf("adding new tournament id: %v, name: %v",
				fullTournament.ID,
				fullTournament.Name)

			err = h.volleynetService.NewTournament(fullTournament)
		} else {
			// err = h.volleynetService.UpdateTournament(fullTournament)
		}

		if err != nil {
			log.Print(err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	c.Status(http.StatusOK)
}
