package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"scores-backend/dtos"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Credentials which stores google ids.
type Credentials struct {
	Cid     string `json:"client_id"`
	Csecret string `json:"client_secret"`
}

// User is a retrieved and authentiacted user.
type User struct {
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

var cred Credentials
var conf *oauth2.Config
var state string

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func initAuth() {
	var redirectURL string
	env := os.Getenv("APP_ENV")
	file, err := ioutil.ReadFile("./client_secret.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &cred)

	if env == "production" {
		redirectURL = "https://scores.raphi011.com/api/auth"
	} else {
		redirectURL = "http://localhost:3000/api/auth"
	}

	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func index(c *gin.Context) {
	c.String(200, "Welcome!")
}

func matchShow(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("matchID"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	match := getMatch(uint(matchID))

	c.JSON(http.StatusOK, match)
}

func matchDelete(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("matchID"))
	userID := c.GetString("userID")

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := deleteMatch(uint(matchID), userID); err != nil {
		c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.Status(http.StatusNoContent)
	}
}

func matchIndex(c *gin.Context) {
	matches := getMatches()

	c.JSON(http.StatusOK, matches)
}

func matchCreate(c *gin.Context) {
	var newMatch dtos.CreateMatchDto
	userID := c.GetString("userID")

	if err := c.ShouldBindJSON(&newMatch); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		match, _ := createMatch(
			newMatch.Player1ID,
			newMatch.Player2ID,
			newMatch.Player3ID,
			newMatch.Player4ID,
			newMatch.ScoreTeam1,
			newMatch.ScoreTeam2,
			userID,
		)

		c.JSON(http.StatusCreated, match)
	}
}

func playerCreate(c *gin.Context) {
	var newPlayer dtos.CreatePlayerDto

	if err := c.ShouldBindJSON(&newPlayer); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		player, _ := createPlayer(newPlayer.Name)
		c.JSON(http.StatusCreated, player)
	}
}

func playerIndex(c *gin.Context) {
	players := getPlayers()

	c.JSON(http.StatusOK, players)
}

func playerStatistic(c *gin.Context) {
	var statistic dtos.Statistic

	c.JSON(http.StatusOK, statistic)
}

func authHandler(c *gin.Context) {
	// Handle the exchange code to initiate a transport.
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	originalState := c.Query("state")
	if retrievedState != originalState {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
		return
	}

	tok, err := conf.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := conf.Client(oauth2.NoContext, tok)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer email.Body.Close()

	user := User{}

	data, _ := ioutil.ReadAll(email.Body)

	if err := json.Unmarshal(data, &user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dbUser := getUserByEmail(user.Email)

	if dbUser.ID == 0 {
		c.Redirect(http.StatusFound, "/loggedIn?error=USER_NOT_FOUND")
	} else {
		if dbUser.ProfileImageURL != user.Picture {
			dbUser.ProfileImageURL = user.Picture
			updateUser(dbUser)
		}

		session.Set("user-id", user.Email)
		session.Save()
		c.Redirect(http.StatusFound, "/loggedIn?username="+user.Email)
	}
}

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

func loginHandler(c *gin.Context) {
	response := dtos.LoginRouteOrUserDto{}
	session := sessions.Default(c)

	state = randToken()

	if userID := session.Get("user-id"); userID != nil {
		user := getUserByEmail(userID.(string))
		response.User = &user
	} else {
		session.Set("state", state)
		session.Save()

		response.LoginRoute = getLoginURL(state)
	}

	c.JSON(http.StatusOK, response)
}

func logoutHandler(c *gin.Context) {
	response := dtos.LoginRouteOrUserDto{}
	state = randToken()
	session := sessions.Default(c)
	session.Delete("user-id")
	session.Set("state", state)
	session.Save()

	response.LoginRoute = getLoginURL(state)

	c.JSON(http.StatusOK, response)
}
