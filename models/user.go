package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Email           string `json:"email"`
	Player          Player `json:"player"`
	PlayerID        uint   `json:"playerId"`
	ProfileImageURL string `json:"profileImageUrl"`
}

type Users []User

func (u *User) GetUserByEmail(db *gorm.DB, email string) {
	db.Where(&User{Email: email}).First(&u)
}

func (u *User) UpdateUser(db *gorm.DB) {
	db.Save(&u)
}
