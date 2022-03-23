//go:build repository
// +build repository

package sql

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/raphi011/scores-api"
	"github.com/raphi011/scores-api/test"
)

func TestCreateUser(t *testing.T) {
	db := SetupDB(t)
	userRepo := &userRepository{DB: db}

	email := "test@test.com"
	id := uuid.New()

	user, err := userRepo.New(&scores.User{
		ID:              id,
		Email:           email,
		ProfileImageURL: "image.url",
	})

	test.Check(t, "userRepository.New() err: %v", err)
	test.Assert(t, "userRepository.New(): want ID != 0, got 0", user.ID == id)
	userByEmail, err := userRepo.ByEmail(email)

	test.Check(t, "userRepository.ByEmail() err: %v", err)

	userByID, err := userRepo.ByID(user.ID)

	test.Check(t, "userRepository.ByID() err: %v", err)
	test.Compare(t, "user byID and byEmail is not equal\n%s", userByEmail, userByID)
}

func TestUserNotFound(t *testing.T) {
	db := SetupDB(t)
	userRepo := &userRepository{DB: db}

	id := uuid.New()

	_, err := userRepo.ByID(id)

	if !errors.Is(err, scores.ErrNotFound) {
		t.Errorf("userRepository.ByID(), want err = ErrNotFound, got: %v", err)
	}
}

func TestUsers(t *testing.T) {
	db := SetupDB(t)
	userRepo := &userRepository{DB: db}

	_, err := userRepo.New(&scores.User{
		ID:    uuid.New(),
		Email: "test@test.at",
	})

	if err != nil {
		t.Errorf("userRepository.New() err: %s", err)
	}

	_, err = userRepo.New(&scores.User{
		ID:              uuid.New(),
		Email:           "test2@test.at",
		ProfileImageURL: "image.url",
	})

	if err != nil {
		t.Errorf("userRepository.New() err: %s", err)
	}

	users, err := userRepo.All()

	if err != nil {
		t.Errorf("userRepository.Users() err: %s", err)
	}

	userCount := len(users)
	if userCount != 2 {
		t.Errorf("len(userRepository.Users()), want 2, got %d", userCount)
	}
}

func TestUpdateUser(t *testing.T) {
	db := SetupDB(t)
	userRepo := &userRepository{DB: db}

	email := "test@test.com"
	newEmail := "test2@test.com"

	user, _ := userRepo.New(&scores.User{
		Email:           email,
		ProfileImageURL: "image.url",
	})

	user.Email = newEmail

	err := userRepo.Update(user)

	if err != nil {
		t.Errorf("userRepository.Update() err: %s", err)
	}

	user, err = userRepo.ByEmail(newEmail)

	if err != nil || user.Email != newEmail {
		t.Error("userRepository.Update(), user not updated")
	}
}
