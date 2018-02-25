package main

import (
	"net/http"
	"strconv"

	"github.com/raphi011/scores"

	"github.com/gin-gonic/gin"
)

type createPlayerDto struct {
	Name string `json:"name"`
}

type playerHandler struct {
	playerService scores.PlayerService
}

func (h *playerHandler) playerCreate(c *gin.Context) {
	var newPlayer createPlayerDto

	if err := c.ShouldBindJSON(&newPlayer); err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
	} else {
		player, err := h.playerService.Create(&scores.Player{Name: newPlayer.Name})

		if err != nil {
			jsonn(c, http.StatusBadRequest, nil, "Bad request")
			return
		}

		jsonn(c, http.StatusCreated, player, "")
	}
}

func (h *playerHandler) playerIndex(c *gin.Context) {
	players, err := h.playerService.Players()

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	jsonn(c, http.StatusOK, players, "")
}

func (h *playerHandler) playerShow(c *gin.Context) {
	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	player, err := h.playerService.Player(uint(playerID))

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Player not found")
		return
	}

	jsonn(c, http.StatusOK, player, "")
}
