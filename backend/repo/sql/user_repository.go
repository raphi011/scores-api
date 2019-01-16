package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/repo"
)

var _ repo.UserRepository = &UserRepository{}

// UserRepository stores users.
type UserRepository struct {
	DB *sqlx.DB
}

// New persists a user and assigns a new id.
func (s *UserRepository) New(user *scores.User) (*scores.User, error) {
	err := insertSetID(s.DB, "user/insert", user)

	return user, errors.Wrap(err, "new user")
}

// Update updates a user.
func (s *UserRepository) Update(user *scores.User) error {
	err := update(s.DB, "user/update", user)

	return errors.Wrap(err, "update user")
}

// All returns all user's, this is used mainly for testing.
func (s *UserRepository) All() ([]*scores.User, error) {
	users, err := s.scan("user/select-all")

	return users, errors.Wrap(err, "all users")
}

// ByID retrieves a user by his/her ID.
func (s *UserRepository) ByID(userID int) (*scores.User, error) {
	user, err := s.scanOne("user/select-by-id", userID)

	return user, errors.Wrap(err, "byID user")
}

// ByEmail retrieves a user by his/her email.
func (s *UserRepository) ByEmail(email string) (*scores.User, error) {
	user, err := s.scanOne("user/select-by-email", email)

	return user, errors.Wrap(err, "byEmail user")
}

func (s *UserRepository) scan(queryName string, args ...interface{}) (
	[]*scores.User, error) {

	users := []*scores.User{}

	q := query(s.DB, queryName)

	rows, err := s.DB.Query(q, args...)

	if err != nil {
		return nil, mapError(err)
	}

	defer rows.Close()

	for rows.Next() {
		u := &scores.User{}

		err := rows.Scan(
			&u.ID,
			&u.CreatedAt,
			&u.Email,
			&u.PasswordInfo.Hash,
			&u.PasswordInfo.Iterations,
			&u.ProfileImageURL,
			&u.Role,
			&u.PasswordInfo.Salt,
			&u.VolleynetUser,
			&u.VolleynetUserID,
		)

		if err != nil {
			return nil, mapError(err)
		}

		users = append(users, u)
	}

	return users, nil
}

func (s *UserRepository) scanOne(query string, args ...interface{}) (
	*scores.User, error) {

	users, err := s.scan(query, args...)

	if err != nil {
		return nil, err
	}

	if len(users) >= 1 {
		return users[0], nil
	}

	return nil, scores.ErrNotFound
}