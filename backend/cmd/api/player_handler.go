package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/raphi011/scores/volleynet/client"
)

func (h *volleynetHandler) getLadder(c *gin.Context) {
	gender := c.DefaultQuery("gender", "M")

	if !h.volleynetService.ValidGender(gender) {
		responseBadRequest(c)
		return
	}

	ladder, err := h.volleynetService.Ladder(gender)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, ladder)
}

func (h *volleynetHandler) getSearchPlayers(c *gin.Context) {
	vnClient := client.DefaultClient()
	firstName := c.Query("fname")
	lastName := c.Query("lname")
	birthday := c.Query("bday")

	players, err := vnClient.SearchPlayers(firstName, lastName, birthday)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, players)
}