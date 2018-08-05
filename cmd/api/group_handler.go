package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/raphi011/scores"

	"github.com/gin-gonic/gin"
)

type groupHandler struct {
	playerService    scores.PlayerService
	groupService     scores.GroupService
	matchService     scores.MatchService
	statisticService scores.StatisticService
}

func (h *groupHandler) index(c *gin.Context) {

}

func (h *groupHandler) groupShow(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("groupID"))

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	group, err := h.groupService.Group(uint(groupID))

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Group not found")
		return
	}

	var pErr, mErr error

	group.Players, pErr = h.playerService.ByGroup(group.ID)
	group.Matches, mErr = h.matchService.GroupMatches(group.ID, time.Now(), 25)

	if pErr != nil || mErr != nil {
		jsonn(c, http.StatusNotFound, nil, "Group not found")
		return
	}

	jsonn(c, http.StatusOK, group, "")
}
