package main

import (
	"net/http"
	"strconv"

	"github.com/raphi011/scores"

	"github.com/gin-gonic/gin"
)

type groupHandler struct {
	service *scores.GroupService
}

func (h *groupHandler) groupShow(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("groupID"))

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	group, err := h.service.Group(uint(groupID))

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Group not found")
		return
	}

	jsonn(c, http.StatusOK, group, "")
}
