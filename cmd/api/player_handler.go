package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin/binding"

	"github.com/raphi011/scores"

	"github.com/gin-gonic/gin"
)

type createPlayerDto struct {
	Name string `json:"name"`
}

type playerHandler struct {
	playerService    *scores.PlayerService
	statisticService *scores.StatisticService
	matchService     *scores.MatchService
}

func (h *playerHandler) postPlayer(c *gin.Context) {
	var newPlayer createPlayerDto

	if err := c.ShouldBindWith(&newPlayer, binding.JSON); err != nil {
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

func (h *playerHandler) getStatistics(c *gin.Context) {
	filter := c.DefaultQuery("filter", "all")
	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	statistic, err := h.statisticService.Player(uint(playerID), filter)

	if err != nil {
		// TODO: check if not found or other error
		jsonn(c, http.StatusNotFound, nil, "Statistic not found")
		return
	}

	jsonn(c, http.StatusOK, statistic, "")
}

func (h *playerHandler) getTeamStatistics(c *gin.Context) {
	filter := c.DefaultQuery("filter", "all")
	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	statistics, err := h.statisticService.PlayerTeams(uint(playerID), filter)

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	jsonn(c, http.StatusOK, statistics, "")
}

func (h *playerHandler) getMatches(c *gin.Context) {
	var err error
	after := time.Now()
	count := uint(25)

	if afterParam := c.Query("after"); afterParam != "" {
		after, err = time.Parse(time.RFC3339, afterParam)

		if err != nil {
			jsonn(c, http.StatusBadRequest, nil, "Bad request")
			return
		}
	}

	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	_, err = h.playerService.Get(uint(playerID))

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Player not found")
		return
	}

	matches, err := h.matchService.ByPlayer(uint(playerID), after, count)

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Match not found")
		return
	}

	jsonn(c, http.StatusOK, matches, "")
}

func (h *playerHandler) getPlayer(c *gin.Context) {
	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	player, err := h.playerService.Get(uint(playerID))

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Player not found")
		return
	}

	jsonn(c, http.StatusOK, player, "")
}
