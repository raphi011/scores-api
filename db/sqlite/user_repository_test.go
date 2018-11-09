package sqlite

import (
	"testing"

	"github.com/raphi011/scores"
)

func TestSetPassword(t *testing.T) {
	s := createRepositories(t)

	email := "test@test.com"

	user, err := s.userRepository.Create(&scores.User{
		Email:           email,
		ProfileImageURL: "image.url",
	})

	pw := []byte("password")

	info, _ := s.pwRepository.HashPassword(pw)

	err = s.userRepository.UpdatePasswordAuthentication(user.ID, info)

	if err != nil {
		t.Errorf("userRepository.UpdatePasswordAuthentication(), err: %s", err)
	}

	user, err = s.userRepository.User(user.ID)

	if err != nil {
		t.Errorf("userRepository.User(), err: %s", err)
	}

	if !s.pwRepository.ComparePassword(pw, &user.PasswordInfo) {
		t.Error("PasswordRepository.ComparePassword(), want true, got false")
	}
}

func TestCreateUser(t *testing.T) {
	s := createRepositories(t)

	email := "test@test.com"

	user, err := s.userRepository.Create(&scores.User{
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

	user, _ = s.userRepository.ByEmail(email)

	if user.ID != userID {
		t.Errorf("userRepository.Create(), user not persisted")
	}
}

func TestUsers(t *testing.T) {
	s := createRepositories(t)

	_, err := s.userRepository.Create(&scores.User{
		Email: "test@test.at",
	})

	_, err = s.userRepository.Create(&scores.User{
		Email:           "test2@test.at",
		ProfileImageURL: "image.url",
	})

	users, err := s.userRepository.Users()

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

	user, _ := s.userRepository.Create(&scores.User{
		Email:           email,
		ProfileImageURL: "image.url",
	})

	user.Email = newEmail

	err := s.userRepository.Update(user)

	if err != nil {
		t.Errorf("UserRepository.Update() err: %s", err)
	}

	user, err = s.userRepository.ByEmail(newEmail)

	if err != nil || user.Email != newEmail {
		t.Error("UserRepository.Update(), user not updated")
	}
}
