package scores

import (
	"github.com/pkg/errors"
)

type UserService struct {
	Repository       UserRepository
	PlayerRepository PlayerRepository
	Password         Password
}

// SetPassword sets a new password for a user
func (s *UserService) SetPassword(
	userID uint,
	password string,
) error {
	user, err := s.Repository.ByID(userID)

	if err != nil {
		return err
	}

	passwordInfo, err := s.Password.Hash([]byte(password))

	if err != nil {
		return errors.Wrap(err, "error hashing password")
	}

	user.PasswordInfo = *passwordInfo

	err = s.Repository.Update(user)

	return errors.Wrap(err, "could not update user password")
}

// ByEmail retrieves a user by email
func (s *UserService) ByEmail(email string) (*User, error) {
	user, err := s.Repository.ByEmail(email)

	if err != nil {
		return nil, errors.Wrapf(err, "could not load user by email %s", email)
	}

	return s.complementUser(user)
}

// ByID retrieves a user by ID
func (s *UserService) ByID(userID uint) (*User, error) {
	user, err := s.Repository.ByID(userID)

	if err != nil {
		return nil, errors.Wrapf(err, "could not load user by ID %d", userID)
	}

	return s.complementUser(user)
}

func (s *UserService) complementUser(user *User) (*User, error) {
	var err error

	if user.PlayerID == 0 {
		return user, nil
	}

	user.Player, err = s.PlayerRepository.Get(user.PlayerID)

	return user, errors.Wrapf(err, "could not load user player %d", user.PlayerID)
}

// SetProfileImage updates a users profile image
func (s *UserService) SetProfileImage(userID uint, imageURL string) error {
	user, err := s.Repository.ByID(userID)

	if err != nil {
		return err
	}

	user.ProfileImageURL = imageURL

	err = s.Repository.Update(user)

	return errors.Wrap(err, "updating profile image")
}

// SetVolleynetLogin updates the users volleynet login
func (s *UserService) SetVolleynetLogin(loginName string, userID int) error {
	user, err := s.Repository.ByID(uint(userID))

	if err != nil {
		return err
	}

	user.VolleynetLogin = loginName
	user.VolleynetUserID = userID

	err = s.Repository.Update(user)

	return errors.Wrap(err, "updatin volleynet login")
}
