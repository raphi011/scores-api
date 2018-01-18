package main

import (
	"net/http"
	"strconv"

	"scores-backend"

	"github.com/gin-gonic/gin"
)

type createMatchDto struct {
	Player1ID  uint `json:"player1Id"`
	Player2ID  uint `json:"player2Id"`
	Player3ID  uint `json:"player3Id"`
	Player4ID  uint `json:"player4Id"`
	ScoreTeam1 int  `json:"scoreTeam1"`
	ScoreTeam2 int  `json:"scoreTeam2"`
}

type matchHandler struct {
	matchService scores.MatchService
	teamService  scores.TeamService
	userService  scores.UserService
}

func (a *matchHandler) index(c *gin.Context) {
	matches, err := a.matchService.Matches()

	if err != nil {
		jsonn(c, http.StatusInternalServerError, nil, "Unknown error")
		return
	}

	jsonn(c, http.StatusOK, matches, "")
}

func (h *matchHandler) matchCreate(c *gin.Context) {
	var newMatch createMatchDto
	userID, _ := strconv.Atoi(c.GetString("userID"))

	if err := c.ShouldBindJSON(&newMatch); err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
	} else {
		team1, err1 := h.teamService.ByPlayers(newMatch.Player1ID, newMatch.Player2ID)
		team2, err2 := h.teamService.ByPlayers(newMatch.Player3ID, newMatch.Player4ID)
		user, err3 := h.userService.User(uint(userID))

		if err1 != nil || err2 != nil || err3 != nil {
			jsonn(c, http.StatusBadRequest, nil, "Bad request")
		}

		// TODO: additional score validation

		match, err := h.matchService.Create(&scores.Match{
			Team1:      team1,
			Team2:      team2,
			ScoreTeam1: newMatch.ScoreTeam1,
			ScoreTeam2: newMatch.ScoreTeam2,
			CreatedBy:  user,
		})

		if err != nil {
			jsonn(c, http.StatusBadRequest, nil, "Bad request")
			return
		}

		jsonn(c, http.StatusCreated, match, "")
	}
}

func (a *matchHandler) matchShow(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("matchID"))

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	match, err := a.matchService.Match(uint(matchID))

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Match not found")
		return
	}

	jsonn(c, http.StatusOK, match, "")
}

func (a *matchHandler) matchDelete(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("matchID"))
	userID := c.GetString("userID")

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	match, err := a.matchService.Match(uint(matchID))

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Match not found")
		return
	}

	user, err := a.userService.ByEmail(userID)

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "User not found")
		return
	}

	if user.ID != match.CreatedBy.ID {
		jsonn(c, http.StatusForbidden, nil, "Match was not created by you")
		return
	}

	a.matchService.Delete(match.ID)

	jsonn(c, http.StatusOK, nil, "")
}
