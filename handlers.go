package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
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

type Credentials struct {
	Cid     string `json:"client_id"`
	Csecret string `json:"client_secret"`
}

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

func JSONN(c *gin.Context, code int, data interface{}, message string) {
	c.JSON(code, gin.H{
		"status":  code,
		"message": message,
		"data":    data,
	})
}

func playerShow(c *gin.Context) {
	playerID, err := strconv.Atoi(c.Param("playerID"))

	if err != nil {
		JSONN(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	player := getPlayer(uint(playerID))

	if player.ID == 0 {
		JSONN(c, http.StatusNotFound, nil, "Player not found")
		return
	}

	JSONN(c, http.StatusOK, player, "")
}

func matchShow(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("matchID"))

	if err != nil {
		JSONN(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	match := getMatch(uint(matchID))

	if match.ID == 0 {
		JSONN(c, http.StatusNotFound, nil, "Match not found")
		return
	}

	JSONN(c, http.StatusOK, match, "")
}

func matchDelete(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("matchID"))
	userID := c.GetString("userID")

	if err != nil {
		JSONN(c, http.StatusBadRequest, nil, "Bad request")
		return
	}

	match := getMatch(uint(matchID))

	if match.ID == 0 {
		JSONN(c, http.StatusNotFound, nil, "Match not found")
		return
	}

	user := getUserByEmail(userID)

	if user.ID != match.CreatedByID {
		JSONN(c, http.StatusForbidden, nil, "Match was not created by you")
		return
	}

	deleteMatch(match)

	JSONN(c, http.StatusOK, nil, "")
}

func matchIndex(c *gin.Context) {
	matches := getMatches()

	JSONN(c, http.StatusOK, matches, "")
}

func matchCreate(c *gin.Context) {
	var newMatch dtos.CreateMatchDto
	userID := c.GetString("userID")

	if err := c.ShouldBindJSON(&newMatch); err != nil {
		JSONN(c, http.StatusBadRequest, nil, "Bad request")
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

		JSONN(c, http.StatusCreated, match, "")
	}
}

func playerCreate(c *gin.Context) {
	var newPlayer dtos.CreatePlayerDto

	if err := c.ShouldBindJSON(&newPlayer); err != nil {
		JSONN(c, http.StatusBadRequest, nil, "Bad request")
	} else {
		player, _ := createPlayer(newPlayer.Name)
		JSONN(c, http.StatusCreated, player, "")
	}
}

func playerIndex(c *gin.Context) {
	players := getPlayers()

	JSONN(c, http.StatusOK, players, "")
}

type player struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func playerStatisticIndex(c *gin.Context) {
	filter := c.DefaultQuery("filter", "all")

	statistics := playersStatistic(filter)

	JSONN(c, http.StatusOK, statistics, "")
}

func authHandler(c *gin.Context) {
	// Handle the exchange code to initiate a transport.
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	originalState := c.Query("state")
	if retrievedState != originalState {
		JSONN(c, http.StatusUnauthorized, nil, "")
		return
	}

	tok, err := conf.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		JSONN(c, http.StatusBadRequest, nil, "")
		return
	}

	client := conf.Client(oauth2.NoContext, tok)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

	if err != nil {
		JSONN(c, http.StatusBadRequest, nil, "")
		return
	}
	defer email.Body.Close()

	user := User{}

	data, _ := ioutil.ReadAll(email.Body)

	if err := json.Unmarshal(data, &user); err != nil {
		JSONN(c, http.StatusBadRequest, nil, "")
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

	JSONN(c, http.StatusOK, response, "")
}

func logoutHandler(c *gin.Context) {
	response := dtos.LoginRouteOrUserDto{}
	state = randToken()
	session := sessions.Default(c)
	session.Delete("user-id")
	session.Set("state", state)
	session.Save()

	response.LoginRoute = getLoginURL(state)

	JSONN(c, http.StatusOK, response, "")
}
