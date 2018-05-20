package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/volleynet"
)

type volleynetHandler struct{}

func (h *volleynetHandler) allTournaments(c *gin.Context) {
	gender := c.DefaultQuery("gender", "M")
	league := c.DefaultQuery("league", "AMATEUR TOUR")
	season := c.DefaultQuery("season", strconv.Itoa(time.Now().Year()))

	client := volleynet.DefaultClient()
	games, err := client.AllTournaments(gender, league, season)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	jsonn(c, http.StatusOK, games, "")
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

	client := volleynet.DefaultClient()
	t, err := client.GetTournament(tournamentID)

	if err == scores.ErrorNotFound {
		c.AbortWithError(http.StatusNotFound, err)
	} else if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		jsonn(c, http.StatusOK, t, "")
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
