package models

import "github.com/jinzhu/gorm"

type Team struct {
	gorm.Model
	Name      string
	Player1   Player
	Player1ID uint
	Player2   Player
	Player2ID uint
}

type Teams []Team
