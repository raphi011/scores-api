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
		responseBadRequest(c)
		return
	}

	match, err := h.matchService.Get(uint(matchID))

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, match)
}

func (h *matchHandler) deleteMatch(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("matchID"))

	if err != nil {
		responseBadRequest(c)
		return
	}

	match, err := h.matchService.Get(uint(matchID))

	if err != nil {
		responseErr(c, err)
		return
	}

	userID := c.GetString("user-id")
	user, err := h.userService.ByEmail(userID)

	if err != nil {
		responseErr(c, err)
		return
	}

	if user.ID != match.CreatedByUserID {
		response(c, http.StatusForbidden, nil)
		return
	}

	h.matchService.Delete(match.ID)

	response(c, http.StatusOK, match)
}
