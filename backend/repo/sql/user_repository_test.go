package sql

import (
	"github.com/pkg/errors"
	"github.com/google/go-cmp/cmp"
	"testing"

	"github.com/raphi011/scores"
)

func TestCreateUser(t *testing.T) {
	db := setupDB(t)
	userRepo := &userRepository{DB: db}

	email := "test@test.com"

	user, err := userRepo.New(&scores.User{
		Email:           email,
		ProfileImageURL: "image.url",
	})

	if err != nil {
		t.Fatalf("userRepository.New() err: %s", err)
	}

	if user.ID == 0 {
		t.Fatalf("userRepository.New(), want ID != 0, got 0")
	}

	userByEmail, err := userRepo.ByEmail(email)

	if err != nil {
		t.Fatalf("userRepository.ByEmail() err: %s", err)
	}

	userByID, err := userRepo.ByID(user.ID)

	if err != nil {
		t.Fatalf("userRepository.ByID() err: %s", err)
	}

	if !cmp.Equal(userByEmail, userByID) {
		t.Fatal("user byID and byEmail is not equal")
	}
}

func TestUserNotFound(t *testing.T) {
	db := setupDB(t)
	userRepo := &userRepository{DB: db}

	_, err := userRepo.ByID(1)

	if errors.Cause(err) != scores.ErrNotFound {
		t.Errorf("userRepository.ByID(), want err = ErrNotFound, got: %v", err)
	}
}

func TestUsers(t *testing.T) {
	db := setupDB(t)
	userRepo := &userRepository{DB: db}

	_, err := userRepo.New(&scores.User{
		Email: "test@test.at",
	})

	_, err = userRepo.New(&scores.User{
		Email:           "test2@test.at",
		ProfileImageURL: "image.url",
	})

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
	db := setupDB(t)
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
