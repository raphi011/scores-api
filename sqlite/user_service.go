package sqlite

import (
	"database/sql"
	"scores-backend"
)

var _ scores.UserService = &UserService{}

type UserService struct {
	DB *sql.DB
}

func (s *UserService) Create(User *scores.User) (*scores.User, error) {

	return &scores.User{}, nil
}

func (s *UserService) Update(user *scores.User) error {

	return nil
}

func (s *UserService) User(userID uint) (*scores.User, error) {

	return &scores.User{}, nil
}

func (u *UserService) ByEmail(email string) (*scores.User, error) {
	// db.Where(&User{Email: email}).First(&u)
	return &scores.User{}, nil
}

// func (u *UserService) UpdateUser(scores.User) {
// 	// db.Save(&u)
// }
