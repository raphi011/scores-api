package models

import (
	"github.com/jinzhu/gorm"
)

type Player struct {
	gorm.Model
	Name string
}

type Players []Player
