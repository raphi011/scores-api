package sqlite

import (
	"testing"

	"github.com/raphi011/scores"
)

func TestCreateUser(t *testing.T) {
	s := createRepositories(t)

	email := "test@test.com"

	user, err := s.User.New(&scores.User{
		Email:           email,
		ProfileImageURL: "image.url",
	})

	if err != nil {
		t.Errorf("userRepository.Create() err: %s", err)
	}

	if user.ID == 0 {
		t.Errorf("userRepository.Create(), want ID != 0, got 0")
	}

	userID := user.ID

	user, _ = s.User.ByEmail(email)

	if user.ID != userID {
		t.Errorf("userRepository.Create(), user not persisted")
	}
}

func TestUsers(t *testing.T) {
	s := createRepositories(t)

	_, err := s.User.New(&scores.User{
		Email: "test@test.at",
	})

	_, err = s.User.New(&scores.User{
		Email:           "test2@test.at",
		ProfileImageURL: "image.url",
	})

	users, err := s.User.All()

	if err != nil {
		t.Errorf("UserRepository.Users() err: %s", err)
	}

	userCount := len(users)
	if userCount != 2 {
		t.Errorf("len(UserRepository.Users()), want 2, got %d", userCount)
	}
}

func TestUpdateUser(t *testing.T) {
	s := createRepositories(t)

	email := "test@test.com"
	newEmail := "test2@test.com"

	user, _ := s.User.New(&scores.User{
		Email:           email,
		ProfileImageURL: "image.url",
	})

	user.Email = newEmail

	err := s.User.Update(user)

	if err != nil {
		t.Errorf("UserRepository.Update() err: %s", err)
	}

	user, err = s.User.ByEmail(newEmail)

	if err != nil || user.Email != newEmail {
		t.Error("UserRepository.Update(), user not updated")
	}
}
