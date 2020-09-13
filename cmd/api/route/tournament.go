package route

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"

	"github.com/raphi011/scores-api/cmd/api/logger"
	"github.com/raphi011/scores-api/repo"
	"github.com/raphi011/scores-api/services"
	"github.com/raphi011/scores-api/volleynet/client"
	volleynet_client "github.com/raphi011/scores-api/volleynet/client"
)

//TournamentHandler is the constructor for the tournament routes handler.
func TournamentHandler(
	volleynetService *services.Volleynet,
	volleynetClient volleynet_client.Client,
	userService *services.User,
) Tournament {
	return Tournament{
		volleynetService: volleynetService,
		volleynetClient:  volleynetClient,
		userService:      userService,
	}
}

// Tournament wraps the depdencies of the TournamentHandler.
type Tournament struct {
	volleynetService *services.Volleynet
	volleynetClient  volleynet_client.Client
	userService      *services.User
}

// GetTournaments queries all available tournaments.
func (h *Tournament) GetTournaments(c *gin.Context) {
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

	if userID, ok := session.Get("user-id").(*uuid.UUID); ok {
		err := h.userService.UpdateTournamentFilter(*userID, filters)

		if err != nil {
			logger.Get(c).Warnf("could not update user settings %v", err)
		}
	}

	response(c, http.StatusOK, tournaments)
}

// GetFilterOptions returns the possible tournament filter values.
func (h *Tournament) GetFilterOptions(c *gin.Context) {
	filters, err := h.volleynetService.TournamentFilterOptions()

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, filters)
}

// GetTournament loads a tournament.
func (h *Tournament) GetTournament(c *gin.Context) {
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
	TournamentID int    `json:"tournamentId"`
	RememberMe   bool   `json:"rememberMe"`
}

// PostSignup allows a player to signup for a tournament.
func (h *Tournament) PostSignup(c *gin.Context) {
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
		h.rememberMe(c, su.Username, loginData.ID)
	}

	err = h.volleynetService.EnterTournament(su.PartnerID, su.TournamentID)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, nil)
}

func (h *Tournament) rememberMe(
	c *gin.Context,
	userName string,
	playerID int,
) {
	session := sessions.Default(c)
	userID := session.Get("user-id").(*uuid.UUID)
	user, err := h.userService.ByID(*userID)

	if err != nil {
		logger.Get(c).Warnf("loading user by id: %s failed", userID)
	}

	if user != nil &&
		user.PlayerLogin != userName ||
		user.PlayerID != playerID {

		err = h.userService.SetVolleynetLogin(*userID, playerID, userName)

		if err != nil {
			logger.Get(c).Warnf("updating volleynet user information failed for userID: %d", user.ID)
		}
	}
}
