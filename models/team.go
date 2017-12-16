package models

import (
	"github.com/jinzhu/gorm"
)

type Team struct {
	Model
	Name      string `json:"name"`
	Player1   Player `json:"player1"`
	Player1ID uint   `json:"player1Id"`
	Player2   Player `json:"player2"`
	Player2ID uint   `json:"player2Id"`
}

type Teams []Team

func (t *Team) GetTeam(db *gorm.DB, player1ID, player2ID uint) {
	if player1ID > player2ID {
		player1ID, player2ID = player2ID, player1ID
	}

	db.Where(Team{Player1ID: player1ID, Player2ID: player2ID}).FirstOrCreate(&t)
}
