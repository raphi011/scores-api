package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/raphi011/scores/cmd/api/logger"
	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/services"
	"github.com/raphi011/scores/volleynet/client"
)

type volleynetHandler struct {
	volleynetService *services.Volleynet
	userService      *services.User
}

func (h *volleynetHandler) getTournaments(c *gin.Context) {
	season := c.QueryArray("seasons")
	gender := c.QueryArray("genders")
	league := c.QueryArray("leagues")

	filters := h.volleynetService.SetDefaultFilters(repo.TournamentFilter{
		Seasons: season,
		Leagues: league,
		Genders: gender,
	})

	tournaments, err := h.volleynetService.SearchTournaments(filters)

	if err != nil {
		responseErr(c, err)
		return
	}

	session := sessions.Default(c)

	if userID, ok := session.Get("user-id").(int); ok {
		err := h.userService.UpdateTournamentFilter(userID, filters)

		if err != nil {
			logger.Get(c).Warnf("could not update user settings %v", err)
		}
	}

	response(c, http.StatusOK, tournaments)
}

func (h *volleynetHandler) getFilterOptions(c *gin.Context) {
	filters, err := h.volleynetService.TournamentFilterOptions()

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, filters)
}

func (h *volleynetHandler) getTournament(c *gin.Context) {
	tournamentID, err := strconv.Atoi(c.Param("tournamentID"))

	if err != nil {
		responseBadRequest(c)
		return
	}

	tournament, err := h.volleynetService.TournamentInfo(tournamentID)

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
		userID := session.Get("user-id").(int)
		user, err := h.userService.ByID(userID)

		if err != nil {
			logger.Get(c).Warnf("loading user by email: %s failed", userID)
		}

		if user != nil && user.VolleynetUser != su.Username ||
			user.VolleynetUserID != loginData.ID {

			err = h.userService.SetVolleynetLogin(userID, loginData.ID, su.Username)

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
