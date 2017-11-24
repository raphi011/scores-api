package main

import (
	"go-test/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func initDb() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open("sqlite3", "/tmp/gorm.db")

	db.AutoMigrate(&models.Player{})
	db.AutoMigrate(&models.Match{})
	db.AutoMigrate(&models.Team{})

	var count int
	db.First(&models.Player{}).Count(&count)

	if count == 0 {
		player1 := models.Player{Name: "Steve"}
		player2 := models.Player{Name: "Tom"}
		player3 := models.Player{Name: "Nick"}
		player4 := models.Player{Name: "Phil"}

		team1 := models.Team{Players: []models.Player{player1, player2}}
		team2 := models.Team{Players: []models.Player{player3, player4}}

		match := models.Match{
			Team1:      team1,
			Team2:      team2,
			ScoreTeam1: 21,
			ScoreTeam2: 18,
		}

		db.Create(&match)
	}

	return db, err
}

func getMatches() []models.Match {
	var matches []models.Match
	db.Preload("Team1.Players").Preload("Team2.Players").Find(&matches)
	return matches
}

func getMatch(id int) models.Match {
	var match models.Match
	db.Preload("Team1.Players").Preload("Team2.Players").First(&match, id)

	return match
}

func createMatch(
	player1ID,
	player2ID,
	player3ID,
	player4ID,
	scoreTeam1,
	scoreTeam2 int) (models.Match, error) {
	// TODO: get (or create) real team ids from players
	team1ID := 1
	team2ID := 2

	match := models.Match{
		Team1ID:    team1ID,
		Team2ID:    team2ID,
		ScoreTeam1: scoreTeam1,
		ScoreTeam2: scoreTeam2,
	}

	db.Create(&match)

	return match, nil
}

func getPlayers() []models.Player {
	var players []models.Player
	db.Find(&players)
	return players
}

func createPlayer(name string) (models.Player, error) {
	player := models.Player{
		Name: name,
	}
	db.Create(&player)

	return player, nil
}
