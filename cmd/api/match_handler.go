package main

import (
	"net/http"
	"strconv"

	"github.com/raphi011/scores"

	"github.com/gin-gonic/gin"
)

type matchHandler struct {
	userService  *scores.UserService
	matchService *scores.MatchService
}

func (h *matchHandler) getMatch(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("matchID"))

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	match, err := h.matchService.Get(uint(matchID))

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Match not found")
		return
	}

	jsonn(c, http.StatusOK, match, "")
}

func (h *matchHandler) deleteMatch(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("matchID"))
	userID := c.GetString("user-id")

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	match, err := h.matchService.Get(uint(matchID))

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Match not found")
		return
	}

	user, err := h.userService.ByEmail(userID)

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "User not found")
		return
	}

	if user.ID != match.CreatedByUserID {
		jsonn(c, http.StatusForbidden, nil, "Match was not created by you")
		return
	}

	h.matchService.Delete(match.ID)

	jsonn(c, http.StatusOK, match, "")
}
