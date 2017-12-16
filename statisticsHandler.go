package main

import (
	"net/http"
	"scores-backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a *App) playerStatisticIndex(c *gin.Context) {
	filter := c.DefaultQuery("filter", "all")

	statistics := models.GetStatistics(a.Db, filter)

	JSONN(c, http.StatusOK, statistics, "")
}

func (a *App) statisticShow(c *gin.Context) {
	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		JSONN(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	statistic := &models.Statistic{}
	statistic.GetStatistic(a.Db, uint(playerID))

	if statistic.ID == 0 {
		JSONN(c, http.StatusNotFound, nil, "Statistic not found")
		return
	}

	JSONN(c, http.StatusOK, statistic, "")
}
