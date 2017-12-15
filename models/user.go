package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email           string
	Player          Player
	PlayerID        uint
	ProfileImageURL string
}

type Users []User

func (u *User) GetUserByEmail(db *gorm.DB, email string) {
	db.Where(&User{Email: email}).First(&u)
}

func (u *User) UpdateUser(db *gorm.DB) {
	db.Save(&u)
}
