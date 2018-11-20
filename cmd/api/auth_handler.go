package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/raphi011/scores"

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

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

type authHandler struct {
	userService  *scores.UserService

	conf     *oauth2.Config
	password scores.Password
}

func (a *authHandler) passwordAuthenticate(c *gin.Context) {
	session := sessions.Default(c)
	var credentials credentialsDto

	if err := c.ShouldBindWith(&credentials, binding.JSON); err != nil {
		jsonn(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	user, err := a.userService.ByEmail(credentials.Email)

	if err != nil {
		jsonn(c, http.StatusUnauthorized, nil, "")
		return
	}

	if !a.password.Compare([]byte(credentials.Password), &user.PasswordInfo) {
		jsonn(c, http.StatusUnauthorized, nil, "")
		return
	}

	response := loginRouteOrUserDto{User: user}

	logger(c).Infof("user %q authenticated via password", user.Email)

	successfullLogin(c, session, user)
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

	user, err := a.userService.ByEmail(googleUser.Email)

	if err != nil {
		c.Redirect(http.StatusFound, "/login?error=USER_NOT_FOUND")
	} else {
		if user.ProfileImageURL != googleUser.Picture {
			err := a.userService.SetProfileImage(user.ID, googleUser.Picture)
			if err != nil {
				logger(c).Errorf("error setting profile image %q", err)
			}
		}

		logger(c).Infof("user %q authenticated via google", googleUser.Email)

		successfullLogin(c, session, user)

		c.Redirect(http.StatusFound, "/")
	}
}

func successfullLogin(c *gin.Context, session sessions.Session, user *scores.User) {
	session.Set("user-id", user.ID)
	session.Save()
}

func (a *authHandler) loginRouteOrUser(c *gin.Context) {
	response := loginRouteOrUserDto{}
	session := sessions.Default(c)

	if userID := session.Get("user-id"); userID != nil {
		user, err := a.userService.ByID(userID.(uint))

		if err != nil {
			session.Delete("user-id")
		} else {
			response.User = user
			jsonn(c, http.StatusOK, response, "")
			return
		}
	}

	state := randToken()
	session.Set("state", state)
	session.Save()

	if a.conf != nil {
		response.LoginRoute = a.conf.AuthCodeURL(state)
	}

	jsonn(c, http.StatusOK, response, "")
}

func (a *authHandler) logout(c *gin.Context) {
	response := loginRouteOrUserDto{}
	state := randToken()
	session := sessions.Default(c)
	session.Delete("user-id")
	session.Set("state", state)
	session.Save()

	if a.conf != nil {
		response.LoginRoute = a.conf.AuthCodeURL(state)
	}

	logger(c).Info("User logged out")

	jsonn(c, http.StatusOK, response, "")
}
