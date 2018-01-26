package main

import (
	"net/http"
	"strconv"

	"scores-backend"

	"github.com/gin-gonic/gin"
)

type createMatchDto struct {
	Player1ID   uint `json:"player1Id"`
	Player2ID   uint `json:"player2Id"`
	Player3ID   uint `json:"player3Id"`
	Player4ID   uint `json:"player4Id"`
	ScoreTeam1  int  `json:"scoreTeam1"`
	ScoreTeam2  int  `json:"scoreTeam2"`
	TargetScore int  `json:"targetScore"`
}

type matchHandler struct {
	playerService scores.PlayerService
	matchService  scores.MatchService
	teamService   scores.TeamService
	userService   scores.UserService
}

func (h *matchHandler) index(c *gin.Context) {
	matches, err := h.matchService.Matches()

	if err != nil {
		jsonn(c, http.StatusInternalServerError, nil, "Unknown error")
		return
	}

	jsonn(c, http.StatusOK, matches, "")
}

func (h *matchHandler) byPlayer(c *gin.Context) {
	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	_, err = h.playerService.Player(uint(playerID))

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Player not found")
		return
	}

	matches, err := h.matchService.PlayerMatches(uint(playerID))

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Match not found")
		return
	}

	jsonn(c, http.StatusOK, matches, "")
}

func (h *matchHandler) matchCreate(c *gin.Context) {
	var newMatch createMatchDto
	userEmail := c.GetString("userID")

	if err := c.ShouldBindJSON(&newMatch); err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	_, pErr1 := h.playerService.Player(newMatch.Player1ID)
	_, pErr2 := h.playerService.Player(newMatch.Player2ID)
	_, pErr3 := h.playerService.Player(newMatch.Player3ID)
	_, pErr4 := h.playerService.Player(newMatch.Player4ID)
	user, uErr := h.userService.ByEmail(userEmail)

	if pErr1 != nil || pErr2 != nil || pErr3 != nil || pErr4 != nil || uErr != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	team1, tErr1 := h.teamService.GetOrCreate(newMatch.Player1ID, newMatch.Player2ID)
	team2, tErr2 := h.teamService.GetOrCreate(newMatch.Player3ID, newMatch.Player4ID)

	if tErr1 != nil || tErr2 != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	// TODO: additional score validation

	match, err := h.matchService.Create(&scores.Match{
		Team1:       team1,
		Team2:       team2,
		ScoreTeam1:  newMatch.ScoreTeam1,
		ScoreTeam2:  newMatch.ScoreTeam2,
		TargetScore: newMatch.TargetScore,
		CreatedBy:   user,
	})

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	jsonn(c, http.StatusCreated, match, "")
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

func (h *matchHandler) matchDelete(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("matchID"))
	userID := c.GetString("userID")

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	match, err := h.matchService.Match(uint(matchID))

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Match not found")
		return
	}

	user, err := h.userService.ByEmail(userID)

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "User not found")
		return
	}

	if user.ID != match.CreatedBy.ID {
		jsonn(c, http.StatusForbidden, nil, "Match was not created by you")
		return
	}

	h.matchService.Delete(match.ID)

	jsonn(c, http.StatusOK, nil, "")
}
