package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/raphi011/scores"

	"github.com/gin-gonic/gin"
)

type groupHandler struct {
	playerRepository    scores.PlayerRepository
	groupRepository     scores.GroupRepository
	matchRepository     scores.MatchRepository
	statisticRepository scores.StatisticRepository
}

func (h *groupHandler) index(c *gin.Context) {

}

func (h *groupHandler) groupShow(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("groupID"))

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	group, err := h.groupRepository.Group(uint(groupID))

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Group not found")
		return
	}

	var pErr, mErr error

	group.Players, pErr = h.playerRepository.ByGroup(group.ID)
	group.Matches, mErr = h.matchRepository.GroupMatches(group.ID, time.Now(), 25)

	if pErr != nil || mErr != nil {
		jsonn(c, http.StatusNotFound, nil, "Group not found")
		return
	}

	jsonn(c, http.StatusOK, group, "")
}
