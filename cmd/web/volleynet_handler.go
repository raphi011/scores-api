package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin/binding"

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
	PartnerName  string `json:"partnerName"`
	TournamentID string `json:"tournamentId"`
}

func getTournamentLink(id string) (string, error) {
	client := volleynet.DefaultClient()
	games, err := client.UpcomingTournaments()

	if err != nil {
		return "", err
	}

	for _, t := range games {
		if t.ID == id {
			return client.ApiUrl + t.Link, nil
		}
	}

	return "", errors.New("Not found")
}

func (h *volleynetHandler) tournament(c *gin.Context) {
	tournamentID := c.Param("tournamentID")

	link, err := getTournamentLink(tournamentID)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	log.Print(link)

	client := volleynet.DefaultClient()
	t, err := client.GetTournament(link)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, t)
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

	c.JSON(http.StatusOK, players)
}
