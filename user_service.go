package scores

import (
	"github.com/pkg/errors"
)

// UserService allows loading / mutation of user data
type UserService struct {
	Repository       UserRepository
	PlayerRepository PlayerRepository
	Password         Password
}

// HasRole verifies if a user has a certain role
func (s *UserService) HasRole(userID uint, roleName string) bool {
	user, err := s.Repository.ByID(userID)

	if err != nil {
		return false
	}

	return user.Role == roleName
}

// New creates a new user
func (s *UserService) New(email, password string) (*User, error) {
	passwordInfo, err := s.Password.Hash([]byte(password))

	if err != nil {
		return nil, errors.Wrap(err, "hashing password")
	}

	user, err := s.Repository.New(&User{
		Email:        email,
		PasswordInfo: *passwordInfo,
		Role:         "user",
	})

	if err != nil {
		return nil, errors.Wrap(err, "creating user")
	}

	user.Player, err = s.PlayerRepository.Create(&Player{
		Name:   "",
		UserID: user.ID,
	})

	if err != nil {
		return nil, errors.Wrap(err, "creating user player")
	}

	return user, nil
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
		return errors.Wrap(err, "hashing password")
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

// All returns all users
func (s *UserService) All() ([]User, error) {
	return s.Repository.All()
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
