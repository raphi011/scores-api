package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/raphi011/scores/repo"
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

func (h *volleynetHandler) getPartners(c *gin.Context) {
	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		responseBadRequest(c)
		return
	}

	partners, err := h.volleynetService.PreviousPartners(playerID)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, partners)
}

func (h *volleynetHandler) getSearchPlayers(c *gin.Context) {
	firstName := c.Query("fname")
	lastName := c.Query("lname")
	gender := c.Query("gender")

	players, err := h.volleynetService.SearchPlayers(repo.PlayerFilter{
		FirstName: firstName,
		LastName:  lastName,
		Gender:    gender,
	})

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, players)
}
