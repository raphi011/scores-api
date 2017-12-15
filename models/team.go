package models

import (
	"github.com/jinzhu/gorm"
)

type Team struct {
	gorm.Model
	Name      string
	Player1   Player
	Player1ID uint
	Player2   Player
	Player2ID uint
}

type Teams []Team

func (t *Team) GetTeam(db *gorm.DB, player1ID, player2ID uint) {
	if player1ID > player2ID {
		player1ID, player2ID = player2ID, player1ID
	}

	db.Where(Team{Player1ID: player1ID, Player2ID: player2ID}).FirstOrCreate(&t)
}
