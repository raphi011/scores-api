package services

import (
	"time"

	"github.com/pkg/errors"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/repo"
)

// User allows loading / mutation of user data
type User struct {
	Repo        repo.UserRepository
	PlayerRepo  repo.PlayerRepository
	SettingRepo repo.SettingRepository

	Password Password
}

// HasRole verifies if a user has a certain role
func (s *User) HasRole(userID int, roleName string) bool {
	user, err := s.Repo.ByID(userID)

	if err != nil {
		return false
	}

	return user.Role == roleName
}

// UpdateSettings updates settings for a user
func (s *User) UpdateSettings(userID int, settings ...*scores.Setting) error {
	currentSettings, err := s.loadSettings(userID)

	if err != nil {
		return err
	}

	for _, setting := range settings {
		var err error

		if setting.Value == "" {
			continue
		}
		if previousSetting, ok := currentSettings[setting.Key]; ok && previousSetting.Value == setting.Value {
			continue
		} else if ok {
			setting.CreatedAt = previousSetting.CreatedAt
			setting.UpdatedAt = time.Now()

			err = s.SettingRepo.Update(setting)
		} else {
			_, err = s.SettingRepo.Create(setting)
		}

		if err != nil {
			return errors.Wrapf(err, "updating user's %q setting key %q", userID, setting.Key)
		}
	}

	return nil
}

// New creates a new user
func (s *User) New(email, password string, role string) (*scores.User, error) {
	passwordInfo, err := s.Password.Hash([]byte(password))

	if err != nil {
		return nil, errors.Wrap(err, "hashing password")
	}

	user, err := s.Repo.New(&scores.User{
		Email:        email,
		PasswordInfo: *passwordInfo,
		Role:         role,
	})

	if err != nil {
		return nil, errors.Wrap(err, "creating user")
	}

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

	user.Settings, err = s.loadSettingsDictionary(user.ID)

	return user, err
}

// ByID retrieves a user by ID
func (s *User) ByID(userID int) (*scores.User, error) {
	user, err := s.Repo.ByID(userID)

	if err != nil {
		return nil, errors.Wrapf(err, "could not load user by ID %d", userID)
	}

	user.Settings, err = s.loadSettingsDictionary(userID)

	return user, err
}

func (s *User) loadSettingsDictionary(userID int) (scores.Settings, error) {
	settings, err := s.SettingRepo.ByUserID(userID)

	if err != nil {
		return nil, errors.Wrapf(err, "load settings for user %q", userID)
	}

	return scores.ToSettingsDictionary(settings), nil
}

type settingsMap map[string]*scores.Setting

func (s *User) loadSettings(userID int) (settingsMap, error) {
	settings, err := s.SettingRepo.ByUserID(userID)

	if err != nil {
		return nil, errors.Wrapf(err, "load settings for user %q", userID)
	}

	settingsMap := map[string]*scores.Setting{}

	for _, setting := range settings {
		settingsMap[setting.Key] = setting
	}

	return settingsMap, nil
}

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
