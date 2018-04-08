package main

import (
	"net/http"
	"strconv"

	"github.com/raphi011/scores"

	"github.com/gin-gonic/gin"
)

type groupHandler struct {
	playerService scores.PlayerService
	groupService  scores.GroupService
}

func (h *groupHandler) index(c *gin.Context) {

}

func (h *groupHandler) groupShow(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("groupId"))

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	group, err := h.groupService.Group(uint(groupID))

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Group not found")
		return
	}

	group.Players, err = h.playerService.ByGroup(group.ID)

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Group not found")
		return
	}

	jsonn(c, http.StatusOK, group, "")
}
