package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/sqlite"
	"github.com/raphi011/scores/volleynet"
)

type volleynetHandler struct {
	volleynetService *sqlite.VolleynetService
}

func (h *volleynetHandler) allTournaments(c *gin.Context) {
	gender := c.DefaultQuery("gender", "M")
	league := c.DefaultQuery("league", "AMATEUR TOUR")
	season := c.DefaultQuery("season", strconv.Itoa(time.Now().Year()))

	seasonNumber, err := strconv.Atoi(season)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	tournaments, err := h.volleynetService.GetTournaments(gender, league, seasonNumber)

	if err != nil {
		log.Print(err)
		c.AbortWithError(http.StatusBadRequest, err)
	}

	jsonn(c, http.StatusOK, tournaments, "")
}

type signupForm struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	PartnerID    string `json:"partnerId"`
	PartnerName  string `json:"partnerName"`
	TournamentID string `json:"tournamentId"`
}

func (h *volleynetHandler) tournament(c *gin.Context) {
	tournamentID := c.Param("tournamentID")

	if tournamentID == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tournament, err := h.volleynetService.Tournament(tournamentID)

	if err == scores.ErrorNotFound {
		c.AbortWithError(http.StatusNotFound, err)
	} else if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		jsonn(c, http.StatusOK, tournament, "")
	}
}

func (h *volleynetHandler) signup(c *gin.Context) {
	su := signupForm{}
	if err := c.ShouldBindWith(&su, binding.JSON); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if su.Username == "" ||
		su.Password == "" ||
		su.PartnerID == "" ||
		su.TournamentID == "" {

		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	client := volleynet.DefaultClient()
	err := client.Login(su.Username, su.Password)

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = client.TournamentEntry(su.PartnerName, su.PartnerID, su.TournamentID)

	if err != nil {
		log.Printf("entry to tournamentID %v with partnerID %v did not work", su.TournamentID, su.PartnerID)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *volleynetHandler) searchPlayers(c *gin.Context) {
	client := volleynet.DefaultClient()
	firstName := c.Query("fname")
	lastName := c.Query("lname")
	birthday := c.Query("bday")
	players, err := client.SearchPlayers(firstName, lastName, birthday)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	jsonn(c, http.StatusOK, players, "")
}

func (h *volleynetHandler) scrapeTournaments(c *gin.Context) {
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
	syncInformation, err := h.volleynetService.SyncInformation(tournaments)

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
			err = h.volleynetService.NewTournament(fullTournament)
		} else {
			err = h.volleynetService.UpdateTournament(fullTournament)
		}

		if err != nil {
			log.Print(err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	c.Status(http.StatusOK)
}
