package route

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/oauth2"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/cmd/api/auth"
	"github.com/raphi011/scores/cmd/api/logger"
	"github.com/raphi011/scores/services"
)

type loginRouteOrUserDto struct {
	LoginRoute string       `json:"loginRoute,omitempty"`
	User       *scores.User `json:"user" `
}

type userDto struct {
	ID              int    `json:"id"`
	Email           string `json:"email"`
	PlayerID        int    `json:"playerId"`
	ProfileImageURL string `json:"profileImageUrl"`
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func AuthHandler(userService *services.User, passwordService services.Password, conf *oauth2.Config) Auth {
	return Auth{
		userService:     userService,
		passwordService: passwordService,
		conf:            conf,
	}
}

type Auth struct {
	userService     *services.User
	passwordService services.Password

	conf *oauth2.Config
}

func (a *Auth) PostPasswordAuthenticate(c *gin.Context) {
	session := sessions.Default(c)
	var credentials auth.PasswordCredentials

	if err := c.ShouldBindWith(&credentials, binding.JSON); err != nil {
		responseBadRequest(c)
		return
	}

	user, err := a.userService.ByEmail(credentials.Email)

	if err != nil {
		response(c, http.StatusUnauthorized, nil)
		return
	}

	if !a.passwordService.Compare([]byte(credentials.Password), &user.PasswordInfo) {
		response(c, http.StatusUnauthorized, nil)
		return
	}

	logger.Get(c).Infof("user %q authenticated via password", user.Email)

	successfullLogin(session, user)

	response(c, http.StatusOK, loginRouteOrUserDto{User: user})
}

func (a *Auth) GetGoogleAuthenticate(c *gin.Context) {
	if a.conf == nil {
		// google oauth config is missing, only password
		// authentication is available
		response(c, http.StatusNotImplemented, nil)
		return
	}

	// Handle the exchange code to initiate a transport.
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	originalState := c.Query("state")
	if retrievedState != originalState {
		response(c, http.StatusUnauthorized, nil)
		return
	}

	tok, err := a.conf.Exchange(oauth2.NoContext, c.Query("code"))

	if err != nil {
		responseBadRequest(c)
		return
	}

	client := a.conf.Client(oauth2.NoContext, tok)
	userinfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

	if err != nil {
		responseBadRequest(c)
		return
	}

	defer userinfo.Body.Close()

	// since the client successfully logged in reset the state token
	newStateToken(session)

	googleUser := auth.OauthUser{}

	data, _ := ioutil.ReadAll(userinfo.Body)

	if err := json.Unmarshal(data, &googleUser); err != nil {
		responseBadRequest(c)
		return
	}

	user, err := a.userService.ByEmail(googleUser.Email)

	if err != nil {
		c.Redirect(http.StatusFound, "/login?error=USER_NOT_FOUND")
		return
	}

	if user.ProfileImageURL != googleUser.Picture {
		err := a.userService.SetProfileImage(user.ID, googleUser.Picture)
		if err != nil {
			logger.Get(c).Warningf("setting profile image %v", err)
		}
	}

	logger.Get(c).Infof("user %q authenticated via google", googleUser.Email)

	successfullLogin(session, user)

	c.Redirect(http.StatusFound, "/")
}

func successfullLogin(session sessions.Session, user *scores.User) {
	session.Set("user-id", user.ID)
	session.Save()
}

func (a *Auth) GetLoginRouteOrUser(c *gin.Context) {
	session := sessions.Default(c)

	if userID, ok := session.Get("user-id").(int); ok {
		user, err := a.userService.ByID(userID)

		if err != nil {
			session.Delete("user-id")
		} else {
			response(c, http.StatusOK, loginRouteOrUserDto{User: user})
			return
		}
	}

	loginRoute := ""

	if a.conf != nil {
		state := ""
		var ok bool

		if state, ok = session.Get("state").(string); !ok {
			state = newStateToken(session)
		}

		loginRoute = a.conf.AuthCodeURL(state)
	}

	response(c, http.StatusOK, loginRouteOrUserDto{LoginRoute: loginRoute})
}

func newStateToken(session sessions.Session) string {
	state := randToken()
	session.Set("state", state)
	session.Save()

	return state
}

func (a *Auth) PostLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()

	loginRoute := ""

	if a.conf != nil {
		state := newStateToken(session)
		loginRoute = a.conf.AuthCodeURL(state)
	}

	logger.Get(c).Info("user logged out")

	response(c, http.StatusOK, loginRouteOrUserDto{LoginRoute: loginRoute})
}
