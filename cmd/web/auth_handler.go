package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/sqlite"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type loginRouteOrUserDto struct {
	LoginRoute string       `json:"loginRoute"`
	User       *scores.User `json:"user"`
}

type userDto struct {
	ID              uint          `json:"id"`
	Email           string        `json:"email"`
	Player          scores.Player `json:"player"`
	PlayerID        uint          `json:"playerId"`
	ProfileImageURL string        `json:"profileImageUrl"`
}

type credentialsDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type oauthUser struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

var state string

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

type authHandler struct {
	userService   *sqlite.UserService
	playerService *sqlite.PlayerService
	groupService  *sqlite.GroupService
	conf          *oauth2.Config
}

func (a *authHandler) passwordAuthenticate(c *gin.Context) {
	session := sessions.Default(c)
	var credentials credentialsDto

	if err := c.ShouldBindWith(&credentials, binding.JSON); err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	user, err := a.getUser(credentials.Email)

	if err != nil {
		jsonn(c, http.StatusUnauthorized, nil, "")
		return
	}

	if !a.userService.PW.ComparePassword([]byte(credentials.Password), &user.PasswordInfo) {
		jsonn(c, http.StatusUnauthorized, nil, "")
		return
	}

	response := loginRouteOrUserDto{User: user}

	successfullLogin(c, session, user.Email)
	jsonn(c, http.StatusOK, response, "")
}

func (a *authHandler) googleAuthenticate(c *gin.Context) {
	if a.conf == nil {
		// google oauth config is missing, only password
		// authentication is available
		jsonn(c, http.StatusNotImplemented, nil, "")
		return
	}
	// Handle the exchange code to initiate a transport.
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	originalState := c.Query("state")
	if retrievedState != originalState {
		jsonn(c, http.StatusUnauthorized, nil, "")
		return
	}

	tok, err := a.conf.Exchange(oauth2.NoContext, c.Query("code"))

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "")
		return
	}

	client := a.conf.Client(oauth2.NoContext, tok)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

	if err != nil {
		jsonn(c, http.StatusBadRequest, nil, "")
		return
	}
	defer email.Body.Close()

	googleUser := oauthUser{}

	data, _ := ioutil.ReadAll(email.Body)

	if err := json.Unmarshal(data, &googleUser); err != nil {
		jsonn(c, http.StatusBadRequest, nil, "")
		return
	}

	user, err := a.getUser(googleUser.Email)

	if err != nil {
		c.Redirect(http.StatusFound, "/login?error=USER_NOT_FOUND")
	} else {
		if user.ProfileImageURL != googleUser.Picture {
			user.ProfileImageURL = googleUser.Picture
			err := a.userService.Update(user)
			if err != nil {
				log.Printf("Error updating user profile image, id: %d", user.ID)
			}
		}

		successfullLogin(c, session, user.Email)
		c.Redirect(http.StatusFound, "/")
	}
}

func (a *authHandler) getUser(email string) (*scores.User, error) {
	user, err := a.userService.ByEmail(email)

	if err != nil {
		return nil, err
	}

	if user.PlayerID > 0 {
		var player *scores.Player
		player, err = a.playerService.Player(user.PlayerID)

		if err != nil {
			return nil, err
		}

		user.Player = player
	}

	return user, nil
}

func successfullLogin(c *gin.Context, session sessions.Session, email string) {
	session.Set("user-id", email)
	session.Save()
}

func (a *authHandler) loginRouteOrUser(c *gin.Context) {
	response := loginRouteOrUserDto{}
	session := sessions.Default(c)

	if userID := session.Get("user-id"); userID != nil {
		user, err := a.getUser(userID.(string))

		if err != nil {
			session.Delete("user-id")
		} else {
			response.User = user
			jsonn(c, http.StatusOK, response, "")
			return
		}
	}

	state = randToken()
	session.Set("state", state)
	session.Save()

	if a.conf != nil {
		response.LoginRoute = a.conf.AuthCodeURL(state)
	}

	jsonn(c, http.StatusOK, response, "")
}

func (a *authHandler) logout(c *gin.Context) {
	response := loginRouteOrUserDto{}
	state = randToken()
	session := sessions.Default(c)
	session.Delete("user-id")
	session.Set("state", state)
	session.Save()

	if a.conf != nil {
		response.LoginRoute = a.conf.AuthCodeURL(state)
	}

	jsonn(c, http.StatusOK, response, "")
}
