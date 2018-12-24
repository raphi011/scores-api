package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/volleynet/client"
	"github.com/raphi011/scores/cmd/api/logger"
)

type volleynetHandler struct {
	volleynetService *scores.VolleynetService
	userService      *scores.UserService
}

func (h *volleynetHandler) getLadder(c *gin.Context) {
	gender := c.DefaultQuery("gender", "M")

	if !h.volleynetService.ValidGender(gender) {
		responseBadRequest(c)
		return
	}

	ladder, err := h.volleynetService.Ladder(gender)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, ladder)
}

func (h *volleynetHandler) getTournaments(c *gin.Context) {
	gender := c.DefaultQuery("gender", "M")
	league := c.DefaultQuery("league", "AMATEUR TOUR")
	season := c.DefaultQuery("season", strconv.Itoa(time.Now().Year()))

	if !h.volleynetService.ValidGender(gender) {
		responseBadRequest(c)
		return
	}

	seasonNumber, err := strconv.Atoi(season)

	if err != nil {
		responseBadRequest(c)
		return
	}

	tournaments, err := h.volleynetService.GetTournaments(gender, league, seasonNumber)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, tournaments)
}

func (h *volleynetHandler) getTournament(c *gin.Context) {
	tournamentID, err := strconv.Atoi(c.Param("tournamentID"))

	if err != nil {
		responseBadRequest(c)
		return
	}

	tournament, err := h.volleynetService.Tournament(tournamentID)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, tournament)
}

type signupForm struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	PartnerID    int    `json:"partnerId"`
	PartnerName  string `json:"partnerName"`
	TournamentID int    `json:"tournamentId"`
	RememberMe   bool   `json:"rememberMe"`
}

func (h *volleynetHandler) postSignup(c *gin.Context) {
	su := signupForm{}

	if err := c.ShouldBindWith(&su, binding.JSON); err != nil {
		responseBadRequest(c)
		return
	}

	if su.Username == "" ||
		su.Password == "" ||
		su.PartnerID <= 0 ||
		su.TournamentID <= 0 {

		responseBadRequest(c)
		return
	}

	vnClient := client.Default()
	loginData, err := vnClient.Login(su.Username, su.Password)

	if err != nil {
		response(c, http.StatusUnauthorized, nil)
		return
	}

	if su.RememberMe {
		session := sessions.Default(c)
		userID := session.Get("user-id")
		user, err := h.userService.ByEmail(userID.(string))

		if err != nil {
			logger.Get(c).Warnf("loading user by email: %s failed", userID.(string))
		}

		if user != nil && user.VolleynetLogin != su.Username ||
			user.VolleynetUserID != loginData.ID {

			err = h.userService.SetVolleynetLogin(su.Username, loginData.ID)

			if err != nil {
				logger.Get(c).Warnf("updating volleynet user information failed for userID: %d", user.ID)
			}
		}
	}

	err = vnClient.TournamentEntry(su.PartnerName, su.PartnerID, su.TournamentID)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, nil)
}

func (h *volleynetHandler) getSearchPlayers(c *gin.Context) {
	vnClient := client.Default()
	firstName := c.Query("fname")
	lastName := c.Query("lname")
	birthday := c.Query("bday")

	players, err := vnClient.SearchPlayers(firstName, lastName, birthday)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, players)
}
