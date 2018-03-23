package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin/binding"

	"github.com/raphi011/scores"

	"github.com/gin-gonic/gin"
)

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

type matchQueryDto struct {
	After time.Time `json:"after"`
	count uint      `json:"count"`
}

type matchHandler struct {
	playerService scores.PlayerService
	groupService  scores.GroupService
	matchService  scores.MatchService
	teamService   scores.TeamService
	userService   scores.UserService
}

func (h *matchHandler) index(c *gin.Context) {
	var err error
	after := time.Now()
	count := uint(25)

	if afterParam := c.Query("after"); afterParam != "" {
		after, err = time.Parse(time.RFC3339, afterParam)

		if err != nil {
			jsonn(c, http.StatusBadRequest, nil, "Bad request")
			return
		}
	}

	matches, err := h.matchService.Matches(after, count)

	if err != nil {
		jsonn(c, http.StatusInternalServerError, nil, "Unknown error")
		return
	}

	jsonn(c, http.StatusOK, matches, "")
}

func (h *matchHandler) byPlayer(c *gin.Context) {
	var err error
	after := time.Now()
	count := uint(25)

	if afterParam := c.Query("after"); afterParam != "" {
		after, err = time.Parse(time.RFC3339, afterParam)

		if err != nil {
			jsonn(c, http.StatusBadRequest, nil, "Bad request")
			return
		}
	}

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

	matches, err := h.matchService.PlayerMatches(uint(playerID), after, count)

	if err != nil {
		jsonn(c, http.StatusNotFound, nil, "Match not found")
		return
	}

	jsonn(c, http.StatusOK, matches, "")
}

func (h *matchHandler) matchCreate(c *gin.Context) {
	var newMatch createMatchDto
	userEmail := c.GetString("userID")

	if err := c.ShouldBindWith(&newMatch, binding.JSON); err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	group, gErr1 := h.groupService.Group(newMatch.GroupID)
	_, pErr1 := h.playerService.Player(newMatch.Player1ID)
	_, pErr2 := h.playerService.Player(newMatch.Player2ID)
	_, pErr3 := h.playerService.Player(newMatch.Player3ID)
	_, pErr4 := h.playerService.Player(newMatch.Player4ID)
	user, uErr := h.userService.ByEmail(userEmail)

	if gErr1 != nil || pErr1 != nil || pErr2 != nil || pErr3 != nil || pErr4 != nil || uErr != nil {
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
		Group:       group,
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

	jsonn(c, http.StatusOK, match, "")
}
