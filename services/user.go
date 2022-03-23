package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/raphi011/scores-api"
	"github.com/raphi011/scores-api/password"
	"github.com/raphi011/scores-api/repo"
)

// User allows loading / mutation of user data
type User struct {
	Repo        repo.UserRepository
	PlayerRepo  repo.PlayerRepository
	SettingRepo repo.SettingRepository

	Password *password.PBKDF2
}

// HasRole verifies if a user has a certain role
func (s *User) HasRole(userID uuid.UUID, roleName string) bool {
	user, err := s.Repo.ByID(userID)

	if err != nil {
		return false
	}

	return user.Role == roleName
}

// UpdateTournamentFilter updates a users tournament filter
func (s *User) UpdateTournamentFilter(userID uuid.UUID, filter repo.TournamentFilter) error {
	return s.UpdateSettings(
		userID,
		&scores.Setting{UserID: userID, Key: "tournament-filter-league", Type: "strings", Value: scores.ListToString(filter.Leagues)},
		&scores.Setting{UserID: userID, Key: "tournament-filter-gender", Type: "strings", Value: scores.ListToString(filter.Genders)},
		&scores.Setting{UserID: userID, Key: "tournament-filter-season", Type: "strings", Value: scores.ListToString(filter.Seasons)},
	)
}

// UpdateSettings updates settings for a user
func (s *User) UpdateSettings(userID uuid.UUID, settings ...*scores.Setting) error {
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
			err = s.SettingRepo.Update(setting)
		} else {
			_, err = s.SettingRepo.Create(setting)
		}

		if err != nil {
			return fmt.Errorf("updating user's %q setting key %q %w", userID, setting.Key, err)
		}
	}

	return nil
}

// New creates a new user
func (s *User) New(email, password string, role string) (*scores.User, error) {
	passwordInfo, err := s.Password.Hash([]byte(password))

	if err != nil {
		return nil, fmt.Errorf("hashing password: %w", err)
	}

	user, err := s.Repo.New(&scores.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordInfo: *passwordInfo,
		Role:         role,
	})

	if err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("creating user player: %w", err)
	}

	return user, nil
}

// SetPassword sets a new password for a user
func (s *User) SetPassword(
	userID uuid.UUID,
	password string,
) error {
	user, err := s.Repo.ByID(userID)

	if err != nil {
		return err
	}

	passwordInfo, err := s.Password.Hash([]byte(password))

	if err != nil {
		return fmt.Errorf("hashing password: %w", err)
	}

	user.PasswordInfo = *passwordInfo

	err = s.Repo.Update(user)

	return fmt.Errorf("could not update user password: %w", err)
}

// ByEmail retrieves a user by email
func (s *User) ByEmail(email string) (*scores.User, error) {
	user, err := s.Repo.ByEmail(email)

	if err != nil {
		return nil, fmt.Errorf("could not load user by email %s %w", email, err)
	}

	user.Settings, err = s.loadSettingsDictionary(user.ID)

	return user, err
}

// ByID retrieves a user by ID
func (s *User) ByID(userID uuid.UUID) (*scores.User, error) {
	user, err := s.Repo.ByID(userID)

	if err != nil {
		return nil, fmt.Errorf("could not load user by ID %d %w", userID, err)
	}

	user.Settings, err = s.loadSettingsDictionary(userID)

	return user, err
}

func (s *User) loadSettingsDictionary(userID uuid.UUID) (scores.Settings, error) {
	settings, err := s.SettingRepo.ByUserID(userID)

	if err != nil {
		return nil, fmt.Errorf("load settings for user %q %w", userID, err)
	}

	return scores.ToSettingsDictionary(settings), nil
}

type settingsMap map[string]*scores.Setting

func (s *User) loadSettings(userID uuid.UUID) (settingsMap, error) {
	settings, err := s.SettingRepo.ByUserID(userID)

	if err != nil {
		return nil, fmt.Errorf("load settings for user %q %w", userID, err)
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
func (s *User) SetProfileImage(userID uuid.UUID, imageURL string) error {
	user, err := s.Repo.ByID(userID)

	if err != nil {
		return err
	}

	user.ProfileImageURL = imageURL

	err = s.Repo.Update(user)

	return fmt.Errorf("updating profile image: %w", err)
}

// SetVolleynetLogin updates the users volleynet login
func (s *User) SetVolleynetLogin(userID uuid.UUID, playerID int, playerLogin string) error {
	user, err := s.Repo.ByID(userID)

	if err != nil {
		return err
	}

	user.PlayerLogin = playerLogin
	user.PlayerID = playerID

	err = s.Repo.Update(user)

	return fmt.Errorf("updatin volleynet login: %w", err)
}
