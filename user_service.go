package scores

import (
	"github.com/pkg/errors"
)

type UserService struct {
	Repository UserRepository
	Password   PasswordService
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
	return s.Repository.ByEmail(email)
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
