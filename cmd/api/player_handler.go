package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/raphi011/scores"
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
		responseBadRequest(c)
		return
	} else {
		player, err := h.playerService.Create(&scores.Player{Name: newPlayer.Name})

		if err != nil {
			responseErr(c, err)
			return
		}

		response(c, http.StatusCreated, player)
	}
}

func (h *playerHandler) getStatistics(c *gin.Context) {
	filter := c.DefaultQuery("filter", "all")
	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		responseBadRequest(c)
		return
	}

	statistic, err := h.statisticService.Player(uint(playerID), filter)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, statistic)
}

func (h *playerHandler) getTeamStatistics(c *gin.Context) {
	filter := c.DefaultQuery("filter", "all")
	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		responseBadRequest(c)
		return
	}

	statistics, err := h.statisticService.PlayerTeams(uint(playerID), filter)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, statistics)
}

func (h *playerHandler) getMatches(c *gin.Context) {
	var err error
	after := time.Now()
	count := uint(25)

	if afterParam := c.Query("after"); afterParam != "" {
		after, err = time.Parse(time.RFC3339, afterParam)

		if err != nil {
			responseBadRequest(c)
			return
		}
	}

	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		responseBadRequest(c)
		return
	}

	_, err = h.playerService.Get(uint(playerID))

	if err != nil {
		responseErr(c, err)
		return
	}

	matches, err := h.matchService.ByPlayer(uint(playerID), after, count)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, matches)
}

func (h *playerHandler) getPlayer(c *gin.Context) {
	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		responseBadRequest(c)
		return
	}

	player, err := h.playerService.Get(uint(playerID))

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, player)
}
