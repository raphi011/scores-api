package sqlite

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/raphi011/scores"
)

var _ scores.UserRepository = &UserRepository{}

type UserRepository struct {
	DB *sql.DB
}

// Create creates persists user and assigns a new id
func (s *UserRepository) Create(user *scores.User) (*scores.User, error) {
	result, err := s.DB.Exec(query("user/insert"),
		user.Email,
		user.ProfileImageURL,
		user.VolleynetUserId,
		user.VolleynetLogin,
		user.Role,
	)

	if err != nil {
		return nil, err
	}

	ID, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	user.ID = uint(ID)

	return user, nil
}

func (s *UserRepository) UpdatePasswordAuthentication(
	userID uint,
	auth *scores.PasswordInfo,
) error {
	result, err := s.DB.Exec(
		query("user/update-password"),
		auth.Salt,
		auth.Hash,
		auth.Iterations,
		userID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("User not found")
	}

	return nil
}

func (s *UserRepository) Update(user *scores.User) error {
	if user == nil || user.ID == 0 {
		return errors.New("User must exist")
	}

	result, err := s.DB.Exec(query("user/update"),
		user.ProfileImageURL,
		user.Email,
		user.VolleynetUserId,
		user.VolleynetLogin,
		user.Role,
		user.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("User not found")
	}

	return nil
}

func (s *UserRepository) Users() (scores.Users, error) {
	users := scores.Users{}

	rows, err := s.DB.Query(query("user/select-all"))

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

func (s *UserRepository) User(userID uint) (*scores.User, error) {
	row := s.DB.QueryRow(query("user/select-by-id"), userID)

	return scanUser(row)
}

func (s *UserRepository) ByEmail(email string) (*scores.User, error) {
	row := s.DB.QueryRow(query("user/select-by-email"), email)

	return scanUser(row)
}

func scanUser(scanner scan) (*scores.User, error) {
	u := &scores.User{}

	err := scanner.Scan(
		&u.ID,
		&u.Email,
		&u.ProfileImageURL,
		&u.PlayerID,
		&u.CreatedAt,
		&u.PasswordInfo.Salt,
		&u.PasswordInfo.Hash,
		&u.PasswordInfo.Iterations,
		&u.VolleynetUserId,
		&u.VolleynetLogin,
		&u.Role,
	)

	return u, err
}
