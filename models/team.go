package models

import "github.com/jinzhu/gorm"

type Team struct {
	gorm.Model
	Name    string
	Players []Player `gorm:"many2many:team_players;"`
}

type Teams []Team
