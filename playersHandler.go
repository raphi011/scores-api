package main

import (
	"net/http"
	"scores-backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type player struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type createPlayerDto struct {
	Name string `json:"name"`
}

func (a *App) playerCreate(c *gin.Context) {
	var newPlayer createPlayerDto

	if err := c.ShouldBindJSON(&newPlayer); err != nil {
		JSONN(c, http.StatusBadRequest, nil, "Bad request")
	} else {
		player := &models.Player{Name: newPlayer.Name}
		player.CreatePlayer(a.Db)
		JSONN(c, http.StatusCreated, player, "")
	}
}

func (a *App) playerIndex(c *gin.Context) {
	players := models.GetPlayers(a.Db)

	JSONN(c, http.StatusOK, players, "")
}

func (a *App) playerShow(c *gin.Context) {
	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		JSONN(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	player := &models.Player{}
	player.GetPlayer(a.Db, uint(playerID))

	if player.ID == 0 {
		JSONN(c, http.StatusNotFound, nil, "Player not found")
		return
	}

	JSONN(c, http.StatusOK, player, "")
}
