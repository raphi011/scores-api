package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"scores-backend/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type loginRouteOrUserDto struct {
	LoginRoute string       `json:"loginRoute"`
	User       *models.User `json:"user"`
}

type userDto struct {
	ID              uint          `json:"id"`
	Email           string        `json:"email"`
	Player          models.Player `json:"player"`
	PlayerID        uint          `json:"playerId"`
	ProfileImageURL string        `json:"profileImageUrl"`
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

func (a *App) authHandler(c *gin.Context) {
	// Handle the exchange code to initiate a transport.
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	originalState := c.Query("state")
	if retrievedState != originalState {
		JSONN(c, http.StatusUnauthorized, nil, "")
		return
	}

	tok, err := a.Conf.Exchange(oauth2.NoContext, c.Query("code"))

	if err != nil {
		JSONN(c, http.StatusBadRequest, nil, "")
		return
	}

	client := a.Conf.Client(oauth2.NoContext, tok)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

	if err != nil {
		JSONN(c, http.StatusBadRequest, nil, "")
		return
	}
	defer email.Body.Close()

	googleUser := oauthUser{}

	data, _ := ioutil.ReadAll(email.Body)

	if err := json.Unmarshal(data, &googleUser); err != nil {
		JSONN(c, http.StatusBadRequest, nil, "")
		return
	}

	user := &models.User{}
	user.GetUserByEmail(a.Db, googleUser.Email)

	if user.ID == 0 {
		c.Redirect(http.StatusFound, "/loggedIn?error=USER_NOT_FOUND")
	} else {
		if user.ProfileImageURL != googleUser.Picture {
			user.ProfileImageURL = googleUser.Picture
			user.UpdateUser(a.Db)
		}

		session.Set("user-id", user.Email)
		session.Save()
		c.Redirect(http.StatusFound, "/loggedIn?username="+user.Email)
	}
}

func (a *App) loginHandler(c *gin.Context) {
	response := loginRouteOrUserDto{}
	session := sessions.Default(c)

	if userID := session.Get("user-id"); userID != nil {
		user := &models.User{}
		user.GetUserByEmail(a.Db, userID.(string))

		if user.ID == 0 {
			session.Delete("user-id")
		} else {
			JSONN(c, http.StatusOK, response, "")
			return
		}
	}

	state = randToken()
	session.Set("state", state)
	session.Save()

	response.LoginRoute = a.Conf.AuthCodeURL(state)

	JSONN(c, http.StatusOK, response, "")
}

func (a *App) logoutHandler(c *gin.Context) {
	response := loginRouteOrUserDto{}
	state = randToken()
	session := sessions.Default(c)
	session.Delete("user-id")
	session.Set("state", state)
	session.Save()

	response.LoginRoute = a.Conf.AuthCodeURL(state)

	JSONN(c, http.StatusOK, response, "")
}
