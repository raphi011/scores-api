package route

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/raphi011/scores-backend/repo"
	"github.com/raphi011/scores-backend/services"
	"github.com/raphi011/scores-backend/volleynet/client"
)

// PlayerHandler is the constructor for the player routes handler.
func PlayerHandler(volleynetService *services.Volleynet, userService *services.User) Player {
	return Player{
		volleynetService: volleynetService,
		userService:      userService,
	}
}

// Player wraps the dependencies of the PlayerHandler.
type Player struct {
	volleynetService *services.Volleynet
	userService      *services.User
}

// GetLadder returns the ladder of a gender.
// Default gender is "M"
func (h *Player) GetLadder(c *gin.Context) {
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

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// PostLogin logs in to volleynet and if successfull saves the
// username of the player.
func (h *Player) PostLogin(c *gin.Context) {
	login := loginForm{}

	if err := c.ShouldBindWith(&login, binding.JSON); err != nil {
		responseBadRequest(c)
		return
	}

	if login.Username == "" ||
		login.Password == "" {
		responseBadRequest(c)
		return
	}

	vnClient := client.DefaultClient()
	loginData, err := vnClient.Login(login.Username, login.Password)

	if err != nil {
		response(c, http.StatusUnauthorized, nil)
		return
	}

	session := sessions.Default(c)
	userID := session.Get("user-id").(int)
	user, err := h.userService.ByID(userID)

	if err != nil {
		responseErr(c, err)
		return
	}

	err = h.userService.SetVolleynetLogin(userID, loginData.ID, login.Username)

	if err != nil {
		responseErr(c, err)
		return
	}

	user, err = h.userService.ByID(userID)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, loginRouteOrUserDto{User: user})
}

// GetPartners returns all previous tournament partners of a player.
func (h *Player) GetPartners(c *gin.Context) {
	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		responseBadRequest(c)
		return
	}

	partners, err := h.volleynetService.PreviousPartners(playerID)

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, partners)
}

// GetSearchPlayers searches all players of a gender.
func (h *Player) GetSearchPlayers(c *gin.Context) {
	firstName := c.Query("fname")
	lastName := c.Query("lname")
	gender := c.Query("gender")

	players, err := h.volleynetService.SearchPlayers(repo.PlayerFilter{
		FirstName: firstName,
		LastName:  lastName,
		Gender:    gender,
	})

	if err != nil {
		responseErr(c, err)
		return
	}

	response(c, http.StatusOK, players)
}
