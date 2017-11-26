package main

import (
	"scores-backend/models"

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
		player1 := models.Player{Name: "Raphi"}
		player2 := models.Player{Name: "Robert"}
		player3 := models.Player{Name: "Lukas"}
		player4 := models.Player{Name: "Richie"}
		player5 := models.Player{Name: "Dominik"}
		player6 := models.Player{Name: "Roman"}

		db.Create(&player1)
		db.Create(&player2)
		db.Create(&player3)
		db.Create(&player4)
		db.Create(&player5)
		db.Create(&player6)
	}

	return db, err
}

func getMatches() []models.Match {
	var matches []models.Match
	db.
		Preload("Team1.Player1").
		Preload("Team1.Player2").
		Preload("Team2.Player1").
		Preload("Team2.Player2").
		Find(&matches)

	return matches
}

func getMatch(id int) models.Match {
	var match models.Match
	db.
		Preload("Team1.Player1").
		Preload("Team1.Player2").
		Preload("Team2.Player1").
		Preload("Team2.Player2").
		First(&match, id)

	return match
}

func getTeam(player1ID, player2ID uint) models.Team {
	if player1ID > player2ID {
		player1ID, player2ID = player2ID, player1ID
	}

	var team models.Team

	db.Where(models.Team{Player1ID: player1ID, Player2ID: player2ID}).FirstOrCreate(&team)

	return team
}

func createMatch(
	player1ID uint,
	player2ID uint,
	player3ID uint,
	player4ID uint,
	scoreTeam1 int,
	scoreTeam2 int) (models.Match, error) {
	team1 := getTeam(player1ID, player2ID)
	team2 := getTeam(player3ID, player4ID)

	match := models.Match{
		Team1:      team1,
		Team2:      team2,
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
