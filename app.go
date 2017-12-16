package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"scores-backend/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	playerStatisticsView = `
		create view if not exists playerStatistics as
		select 
		p.id,
		p.name,
		m.created_at,
		case when 
			(t1.player1_id = p.id or t1.player2_id = p.id) and (m.score_team1 > m.score_team2) or 
			(t2.player1_id = p.id or t2.player2_id = p.id) and (m.score_team2 > m.score_team1)
		 then 1 else 0 end as won,
		case when t1.player1_id = p.id or t1.player2_id = p.id then m.score_team1 else m.score_team2 end as pointsWon,
		case when t1.player1_id = p.id or t1.player2_id = p.id then m.score_team2 else m.score_team1 end as pointsLost
		from matches m	
		join teams t1 on m.team1_id = t1.id
		join teams t2 on m.team2_id = t2.id
		join players p on t1.player1_id = p.id or t1.player2_id = p.id or t2.player1_id = p.id or t2.player2_id = p.id
		where m.deleted_at is null
	`
)

type credentials struct {
	Cid     string `json:"client_id"`
	Csecret string `json:"client_secret"`
}

type App struct {
	Router     *gin.Engine
	Db         *gorm.DB
	Conf       *oauth2.Config
	Production bool
}

func (a *App) initDb() {
	var err error
	a.Db, err = gorm.Open("sqlite3", "/tmp/gorm.db")

	if err != nil {
		panic("Could not open database")
	}

	a.Db.Exec(playerStatisticsView)
	a.Db.AutoMigrate(&models.Player{})
	a.Db.AutoMigrate(&models.Match{})
	a.Db.AutoMigrate(&models.Team{})
	a.Db.AutoMigrate(&models.User{})

	var count int
	a.Db.First(&models.User{}).Count(&count)

	if count == 0 {
		player1 := models.Player{Name: "Raphi"}
		player2 := models.Player{Name: "Robert"}
		player3 := models.Player{Name: "Lukas"}
		player4 := models.Player{Name: "Richie"}
		player5 := models.Player{Name: "Dominik"}
		player6 := models.Player{Name: "Roman"}

		user1 := models.User{Email: "raphi011@gmail.com", Player: player1}
		user2 := models.User{Email: "", Player: player2}
		user3 := models.User{Email: "", Player: player3}
		user4 := models.User{Email: "Rb1@outlook.at", Player: player4}
		user5 := models.User{Email: "Rieder.dominik@gmail.com", Player: player5}
		user6 := models.User{Email: "", Player: player6}

		a.Db.Create(&user1)
		a.Db.Create(&user2)
		a.Db.Create(&user3)
		a.Db.Create(&user4)
		a.Db.Create(&user5)
		a.Db.Create(&user6)
	}
}

func (a *App) initRouter() {
	if a.Production {
		a.Router = gin.Default()
	} else {
		gin.SetMode(gin.TestMode)
		a.Router = gin.New()
		a.Router.Use(gin.Recovery())
	}

	a.Router.Use(sessions.Sessions("goquestsession", store))

	a.Router.GET("/matches", a.matchIndex)
	a.Router.GET("/matches/:matchID", a.matchShow)
	a.Router.GET("/players", a.playerIndex)
	a.Router.GET("/players/:playerID", a.playerShow)
	a.Router.GET("/statistics", a.playerStatisticIndex)
	a.Router.GET("/statistics/:playerID", a.statisticShow)
	a.Router.GET("/userOrLoginRoute", a.loginHandler)
	a.Router.GET("/auth", a.authHandler)
	a.Router.POST("/logout", a.logoutHandler)

	auth := a.Router.Group("/")
	auth.Use(authRequired())
	{
		auth.DELETE("/matches/:matchID", a.matchDelete)
		auth.POST("/players", a.playerCreate)
		auth.POST("/matches", a.matchCreate)
	}
}

func (a *App) initAuth() {
	var redirectURL string
	var cred credentials
	file, err := ioutil.ReadFile("./client_secret.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &cred)

	if a.Production {
		redirectURL = "https://scores.raphi011.com/api/auth"
	} else {
		redirectURL = "http://localhost:3000/api/auth"
	}

	a.Conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func (a *App) Initialize() {
	a.initDb()
	a.initAuth()
	a.initRouter()
}

func (a *App) Run() {
	a.Router.Run()
	defer a.Db.Close()
}
