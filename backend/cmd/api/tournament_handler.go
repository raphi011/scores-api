package main

import (
	"net/http"
	"time"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"

	"github.com/raphi011/scores/services"
	"github.com/raphi011/scores/volleynet/client"
	"github.com/raphi011/scores/cmd/api/logger"
)

type volleynetHandler struct {
	volleynetService *services.Volleynet
	userService      *services.User
}


func (h *volleynetHandler) getTournaments(c *gin.Context) {
	season := c.DefaultQuery("season", strconv.Itoa(time.Now().Year()))
	gender := c.QueryArray("gender")
	league := c.QueryArray("league")

	// if !h.volleynetService.ValidGender(gender) {
	// 	responseBadRequest(c)
	// 	return
	// }

	seasonNumber, err := strconv.Atoi(season)

	// if err != nil {
	// 	responseBadRequest(c)
	// 	return
	// }

	tournaments, err := h.volleynetService.GetTournaments(
		[]int{seasonNumber},
		gender,
		league,
	)

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

	vnClient := client.DefaultClient()
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

		if user != nil && user.VolleynetUser != su.Username ||
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
