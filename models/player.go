package models

import (
	"github.com/jinzhu/gorm"
)

type Player struct {
	gorm.Model
	Name string
}

type Players []Player

func GetPlayers(db *gorm.DB) Players {
	var players []Player
	db.Order("name").Find(&players)
	return players
}

func (p *Player) CreatePlayer(db *gorm.DB) {
	db.Create(&p)
}

func (p *Player) GetPlayer(db *gorm.DB, ID uint) {
	db.First(&p, ID)
}
