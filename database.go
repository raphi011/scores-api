package main

import (
	"errors"
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
	db.AutoMigrate(&models.User{})

	var count int
	db.First(&models.User{}).Count(&count)

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

		db.Create(&user1)
		db.Create(&user2)
		db.Create(&user3)
		db.Create(&user4)
		db.Create(&user5)
		db.Create(&user6)
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

func getMatch(id uint) models.Match {
	var match models.Match
	db.
		Preload("Team1.Player1").
		Preload("Team1.Player2").
		Preload("Team2.Player1").
		Preload("Team2.Player2").
		Preload("CreatedBy").
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

func deleteMatch(matchID uint, userEmail string) error {
	user := getUserByEmail(userEmail)
	match := getMatch(matchID)

	if user.ID != match.CreatedByID {
		return errors.New("Match was not created by you")
	}

	db.Delete(&match)

	return nil
}

func getUserByEmail(email string) models.User {
	var user models.User

	db.Where(&User{Email: email}).First(&user)

	return user
}

func createMatch(
	player1ID uint,
	player2ID uint,
	player3ID uint,
	player4ID uint,
	scoreTeam1 int,
	scoreTeam2 int,
	userEmail string,
) (models.Match, error) {
	user := getUserByEmail(userEmail)
	team1 := getTeam(player1ID, player2ID)
	team2 := getTeam(player3ID, player4ID)

	match := models.Match{
		Team1:       team1,
		Team2:       team2,
		ScoreTeam1:  scoreTeam1,
		ScoreTeam2:  scoreTeam2,
		CreatedByID: user.ID,
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
