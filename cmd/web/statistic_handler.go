package main

import (
	"net/http"
	"scores-backend"
	"strconv"

	"github.com/gin-gonic/gin"
)

type statisticHandler struct {
	statisticService scores.StatisticService
}

func (h *statisticHandler) playerTeams(c *gin.Context) {
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

func (h *statisticHandler) players(c *gin.Context) {
	filter := c.DefaultQuery("filter", "all")

	statistics, err := h.statisticService.Players(filter)

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	jsonn(c, http.StatusOK, statistics, "")
}

func (h *statisticHandler) player(c *gin.Context) {
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
