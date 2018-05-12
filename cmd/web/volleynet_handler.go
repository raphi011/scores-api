package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/raphi011/scores/volleynet"
)

type volleynetHandler struct{}

func (h *volleynetHandler) allTournaments(c *gin.Context) {
	client := volleynet.DefaultClient()
	games, _ := client.UpcomingTournaments()

	c.JSON(http.StatusOK, games)
}

type signupForm struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	PartnerID    string `json:"partnerId"`
	TournamentID string `json:"tournamentId"`
}

func (h *volleynetHandler) tournament(c *gin.Context) {
	tournamentID := c.Param("tournamentID")

	/* testing */
	tournamentID = "http://www.volleynet.at/beach/bewerbe/AMATEUR%20TOUR/phase/ABV%20Tour%20AMATEUR%201/sex/M/saison/2018/cup/22108"

	client := volleynet.DefaultClient()
	t, err := client.GetTournament(tournamentID)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, t)
}

func (h *volleynetHandler) signup(c *gin.Context) {
	su := signupForm{}
	if err := c.BindJSON(su); err != nil {
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

	err = client.TournamentEntry(su.PartnerID, su.TournamentID)

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

	c.JSON(http.StatusOK, players)
}
