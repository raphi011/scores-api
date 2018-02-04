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

func scanUser(scanner scan) (*scores.User, error) {
	u := scores.User{}
	var profileImageURL sql.NullString
	err := scanner.Scan(&u.ID, &u.Email, &profileImageURL)

	if err != nil {
		return nil, err
	}
	if profileImageURL.Valid {
		u.ProfileImageURL = profileImageURL.String
	}

	return &u, nil
}

const (
	usersSelectSQL = `
		SELECT
			id,
			email,
			profile_image_url
		FROM users 
		WHERE deleted_at is null
	`

	userByIDSelectSQL    = usersSelectSQL + " and id = $1"
	userByEmailSelectSQL = usersSelectSQL + " and email = $1"
)

func (s *UserService) Users() (scores.Users, error) {
	users := scores.Users{}

	rows, err := s.DB.Query(usersSelectSQL)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		u, err := scanUser(rows)

		if err != nil {
			return nil, err
		}

		users = append(users, *u)
	}

	return users, nil
}

func (s *UserService) User(userID uint) (*scores.User, error) {
	row := s.DB.QueryRow(userByIDSelectSQL, userID)

	user, err := scanUser(row)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) ByEmail(email string) (*scores.User, error) {
	row := s.DB.QueryRow(userByEmailSelectSQL, email)

	user, err := scanUser(row)

	if err != nil {
		return nil, err
	}

	return user, nil
}
