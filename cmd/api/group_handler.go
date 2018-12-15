package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/raphi011/scores"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type groupHandler struct {
	service          *scores.GroupService
	playerService    *scores.PlayerService
	matchService     *scores.MatchService
	statisticService *scores.StatisticService
}

type createMatchDto struct {
	GroupID     uint `json:"groupId"`
	Player1ID   uint `json:"player1Id"`
	Player2ID   uint `json:"player2Id"`
	Player3ID   uint `json:"player3Id"`
	Player4ID   uint `json:"player4Id"`
	ScoreTeam1  int  `json:"scoreTeam1"`
	ScoreTeam2  int  `json:"scoreTeam2"`
	TargetScore int  `json:"targetScore"`
}

func (h *groupHandler) postMatch(c *gin.Context) {
	var newMatch createMatchDto
	userID := c.GetInt("user-id")

	if err := c.ShouldBindWith(&newMatch, binding.JSON); err != nil {
		responseBadRequest(c)
		return
	}

	match, err := h.matchService.Create(&scores.Match{
		CreatedByUserID: uint(userID),
	})

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusCreated, match)
}

func (h *groupHandler) getGroup(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("groupID"))

	if err != nil {
		responseBadRequest(c)
		return
	}

	group, err := h.service.Group(uint(groupID))

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, group)
}

func (h *groupHandler) getMatches(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("groupID"))

	if err != nil {
		responseBadRequest(c)
		return
	}

	after := time.Now()
	count := uint(25)

	if afterParam := c.Query("after"); afterParam != "" {
		after, err = time.Parse(time.RFC3339, afterParam)

		if err != nil {
			responseBadRequest(c)
			return
		}
	}

	matches, err := h.matchService.ByGroup(uint(groupID), after, count)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, matches)
}

func (h *groupHandler) getPlayers(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("groupID"))

	if groupID < 0 {
		responseBadRequest(c)
		return
	}

	players, err := h.playerService.ByGroup(uint(groupID))

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, players)
}

func (h *groupHandler) getPlayerStatistics(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	filter := c.DefaultQuery("filter", "all")

	statistics, err := h.statisticService.PlayersByGroup(uint(groupID), filter)

	if err != nil {
		responseBadRequest(c)
		return
	}

	response(c, http.StatusOK, statistics)
}
