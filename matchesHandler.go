package main

import (
	"net/http"
	"scores-backend/models"
	"strconv"

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

func (a *App) matchIndex(c *gin.Context) {
	matches := models.GetMatches(a.Db)

	JSONN(c, http.StatusOK, matches, "")
}

func (a *App) matchCreate(c *gin.Context) {
	var newMatch createMatchDto
	userID := c.GetString("userID")

	if err := c.ShouldBindJSON(&newMatch); err != nil {
		JSONN(c, http.StatusBadRequest, nil, "Bad request")
	} else {
		match := &models.Match{}
		match.CreateMatch(
			a.Db,
			newMatch.Player1ID,
			newMatch.Player2ID,
			newMatch.Player3ID,
			newMatch.Player4ID,
			newMatch.ScoreTeam1,
			newMatch.ScoreTeam2,
			userID,
		)

		JSONN(c, http.StatusCreated, match, "")
	}
}

func (a *App) matchShow(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("matchID"))
	match := &models.Match{}

	if err != nil {
		JSONN(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	match.GetMatch(a.Db, uint(matchID))

	if match.ID == 0 {
		JSONN(c, http.StatusNotFound, nil, "Match not found")
		return
	}

	JSONN(c, http.StatusOK, match, "")
}

func (a *App) matchDelete(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("matchID"))
	userID := c.GetString("userID")

	if err != nil {
		JSONN(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	match := &models.Match{}
	match.GetMatch(a.Db, uint(matchID))

	if match.ID == 0 {
		JSONN(c, http.StatusNotFound, nil, "Match not found")
		return
	}

	user := &models.User{}
	user.GetUserByEmail(a.Db, userID)

	if user.ID != match.CreatedByID {
		JSONN(c, http.StatusForbidden, nil, "Match was not created by you")
		return
	}

	match.DeleteMatch(a.Db)

	JSONN(c, http.StatusOK, nil, "")
}
