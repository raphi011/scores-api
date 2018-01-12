package sqlite

import (
	"database/sql"
	"errors"
	"scores-backend"
)

var _ scores.UserService = &UserService{}

type UserService struct {
	DB *sql.DB
}

const userInsertSQL = `
	INSERT INTO users (created_at, email, profile_image_url)
	VALUES (CURRENT_TIMESTAMP, $1, $2)
`

func (s *UserService) Create(user *scores.User) (*scores.User, error) {
	result, err := s.DB.Exec(userInsertSQL, user.Email, user.ProfileImageURL)

	if err != nil {
		return nil, err
	}

	ID, _ := result.LastInsertId()

	user.ID = uint(ID)

	return user, nil
}

const userUpdateSQL = `
	UPDATE users
	SET profile_image_url = $1, email = $2
	WHERE id = $3
`

func (s *UserService) Update(user *scores.User) error {
	if user == nil || user.ID == 0 {
		return errors.New("User must exist")
	}

	result, err := s.DB.Exec(userUpdateSQL, user.ProfileImageURL, user.Email, user.ID)

	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected != 1 {
		return errors.New("User not found")
	}

	return nil
}

const userSelectSQL = `
	SELECT email, profile_image_url
	FROM users 
	WHERE id = $1
`

func (s *UserService) User(userID uint) (*scores.User, error) {
	user := &scores.User{}
	user.ID = userID

	err := s.DB.QueryRow(userSelectSQL, userID).Scan(&user.Email, &user.ProfileImageURL)

	if err != nil {
		return nil, err
	}

	return user, nil
}

const usersSelectSQL = `
		SELECT
			id,
			email,
			profile_image_url
		FROM users 
		WHERE deleted_at is null
`

func (s *UserService) Users() (scores.Users, error) {
	users := scores.Users{}

	rows, err := s.DB.Query(usersSelectSQL)

	for rows.Next() {
		u := scores.User{}
		err = rows.Scan(&u.ID, &u.Email, &u.ProfileImageURL)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

const userByEmailSelectSQL = `
	SELECT id, profile_image_url
	FROM users 
	WHERE email = $1
`

func (s *UserService) ByEmail(email string) (*scores.User, error) {
	user := &scores.User{Email: email}

	err := s.DB.QueryRow(userByEmailSelectSQL, email).Scan(&user.ID, &user.ProfileImageURL)

	if err != nil {
		return nil, err
	}

	return user, nil
}
