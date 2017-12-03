package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email           string
	Player          Player
	PlayerID        uint
	ProfileImageURL string
}

type Users []User
