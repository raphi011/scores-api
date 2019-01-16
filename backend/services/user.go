package services

import (
	"github.com/pkg/errors"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/repo"
)

// scores.User allows loading / mutation of user data
type User struct {
	Repo       repo.UserRepository
	PlayerRepo repo.PlayerRepository

	Password         Password
}

// HasRole verifies if a user has a certain role
func (s *User) HasRole(userID int, roleName string) bool {
	user, err := s.Repo.ByID(userID)

	if err != nil {
		return false
	}

	return user.Role == roleName
}

// New creates a new user
func (s *User) New(email, password string) (*scores.User, error) {
	passwordInfo, err := s.Password.Hash([]byte(password))

	if err != nil {
		return nil, errors.Wrap(err, "hashing password")
	}

	user, err := s.Repo.New(&scores.User{
		Email:        email,
		PasswordInfo: *passwordInfo,
		Role:         "user",
	})

	if err != nil {
		return nil, errors.Wrap(err, "creating user")
	}

	// user.Player, err = s.PlayerRepository.Create(&Player{
	// 	Name:   "",
	// 	UserID: user.ID,
	// })

	if err != nil {
		return nil, errors.Wrap(err, "creating user player")
	}

	return user, nil
}

// SetPassword sets a new password for a user
func (s *User) SetPassword(
	userID int,
	password string,
) error {
	user, err := s.Repo.ByID(userID)

	if err != nil {
		return err
	}

	passwordInfo, err := s.Password.Hash([]byte(password))

	if err != nil {
		return errors.Wrap(err, "hashing password")
	}

	user.PasswordInfo = *passwordInfo

	err = s.Repo.Update(user)

	return errors.Wrap(err, "could not update user password")
}

// ByEmail retrieves a user by email
func (s *User) ByEmail(email string) (*scores.User, error) {
	user, err := s.Repo.ByEmail(email)

	if err != nil {
		return nil, errors.Wrapf(err, "could not load user by email %s", email)
	}

	return user, nil
	// return s.complementUser(user)
}

// ByID retrieves a user by ID
func (s *User) ByID(userID int) (*scores.User, error) {
	user, err := s.Repo.ByID(userID)

	if err != nil {
		return nil, errors.Wrapf(err, "could not load user by ID %d", userID)
	}

	return user, nil
	// return s.complementUser(user)
}

// func (s *User) complementUser(user *scores.User) (*scores.User, error) {
// 	var err error

// 	if user.PlayerID == 0 {
// 		return user, nil
// 	}

// 	user.Player, err = s.PlayerRepo.Get(user.PlayerID)

// 	return user, errors.Wrapf(err, "could not load user player %d", user.PlayerID)
// }

// All returns all users
func (s *User) All() ([]*scores.User, error) {
	return s.Repo.All()
}

// SetProfileImage updates a users profile image
func (s *User) SetProfileImage(userID int, imageURL string) error {
	user, err := s.Repo.ByID(userID)

	if err != nil {
		return err
	}

	user.ProfileImageURL = imageURL

	err = s.Repo.Update(user)

	return errors.Wrap(err, "updating profile image")
}

// SetVolleynetLogin updates the users volleynet login
func (s *User) SetVolleynetLogin(loginName string, userID int) error {
	user, err := s.Repo.ByID(userID)

	if err != nil {
		return err
	}

	user.VolleynetUser = loginName
	user.VolleynetUserID = userID

	err = s.Repo.Update(user)

	return errors.Wrap(err, "updatin volleynet login")
}