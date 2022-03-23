package route

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/raphi011/scores-api"
	"github.com/raphi011/scores-api/cmd/api/auth"
	"github.com/raphi011/scores-api/cmd/api/logger"
	"github.com/raphi011/scores-api/password"
	"github.com/raphi011/scores-api/services"
	"golang.org/x/oauth2"
)

func init() {
	// needed for cookie storage marshalling
	gob.Register(&uuid.UUID{})
}

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

// AuthHandler is the constructor for the Auth routes handler.
func AuthHandler(userService *services.User, passwordService *password.PBKDF2, conf *oauth2.Config) Auth {
	return Auth{
		userService:     userService,
		passwordService: passwordService,
		conf:            conf,
	}
}

// Auth handles the authentication routes.
type Auth struct {
	userService     *services.User
	passwordService *password.PBKDF2

	conf *oauth2.Config
}

// PostPasswordAuthenticate handles the password authentication route.
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

// GetGoogleAuthenticate handles is the oauth endpoint for a successfull
// oauth authentication from google.
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
			logger.Get(c).Warnf("setting profile image %v", err)
		}
	}

	logger.Get(c).Infof("user %q authenticated via google", googleUser.Email)

	successfullLogin(session, user)

	c.Redirect(http.StatusFound, "/")
}

func successfullLogin(session sessions.Session, user *scores.User) {
	session.Set("user-id", &user.ID)
	err := session.Save()

	if err != nil {
		panic(err)
	}
}

// GetLoginRouteOrUser returns the oauth login route or the user
// if logged in.
func (a *Auth) GetLoginRouteOrUser(c *gin.Context) {
	session := sessions.Default(c)

	if userID, ok := session.Get("user-id").(*uuid.UUID); ok {
		user, err := a.userService.ByID(*userID)

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

// PostLogout handles the logout route.
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
