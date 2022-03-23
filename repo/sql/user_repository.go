package sql

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/raphi011/scores-api"
	"github.com/raphi011/scores-api/repo"
	"github.com/raphi011/scores-api/repo/sql/crud"
)

var _ repo.UserRepository = &userRepository{}

type userRepository struct {
	DB *sqlx.DB
}

// New persists a user and assigns a new id.
func (s *userRepository) New(user *scores.User) (*scores.User, error) {
	err := crud.Create(s.DB, "user/insert", user)

	return user, fmt.Errorf("new user: %w", err)
}

// Update updates a user.
func (s *userRepository) Update(user *scores.User) error {
	err := crud.Update(s.DB, "user/update", user)

	return fmt.Errorf("update user: %w", err)
}

// All returns all user's, this is used mainly for testing.
func (s *userRepository) All() ([]*scores.User, error) {

	users := []*scores.User{}
	err := crud.Read(s.DB, "user/select-all", &users)

	return users, fmt.Errorf("all users: %w", err)
}

// ByID retrieves a user by his/her ID.
func (s *userRepository) ByID(userID uuid.UUID) (*scores.User, error) {

	user := &scores.User{}
	err := crud.ReadOne(s.DB, "user/select-by-id", user, userID)

	return user, fmt.Errorf("byID user: %w", err)
}

// ByEmail retrieves a user by his/her email.
func (s *userRepository) ByEmail(email string) (*scores.User, error) {
	user := &scores.User{}
	err := crud.ReadOne(s.DB, "user/select-by-email", user, email)

	return user, fmt.Errorf("byEmail user: %w", err)
}
