package sql

import (
	"github.com/google/go-cmp/cmp"
	"testing"

	"github.com/raphi011/scores"
)

func TestCreateUser(t *testing.T) {
	db := setupDB(t)
	userRepo := &UserRepository{DB: db}

	email := "test@test.com"

	user, err := userRepo.New(&scores.User{
		Email:           email,
		ProfileImageURL: "image.url",
	})

	if err != nil {
		t.Fatalf("userRepository.Create() err: %s", err)
	}

	if user.ID == 0 {
		t.Fatalf("userRepository.Create(), want ID != 0, got 0")
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

func TestUsers(t *testing.T) {
	db := setupDB(t)
	userRepo := &UserRepository{DB: db}

	_, err := userRepo.New(&scores.User{
		Email: "test@test.at",
	})

	_, err = userRepo.New(&scores.User{
		Email:           "test2@test.at",
		ProfileImageURL: "image.url",
	})

	users, err := userRepo.All()

	if err != nil {
		t.Errorf("UserRepository.Users() err: %s", err)
	}

	userCount := len(users)
	if userCount != 2 {
		t.Errorf("len(UserRepository.Users()), want 2, got %d", userCount)
	}
}

func TestUpdateUser(t *testing.T) {
	db := setupDB(t)
	userRepo := &UserRepository{DB: db}

	email := "test@test.com"
	newEmail := "test2@test.com"

	user, _ := userRepo.New(&scores.User{
		Email:           email,
		ProfileImageURL: "image.url",
	})

	user.Email = newEmail

	err := userRepo.Update(user)

	if err != nil {
		t.Errorf("UserRepository.Update() err: %s", err)
	}

	user, err = userRepo.ByEmail(newEmail)

	if err != nil || user.Email != newEmail {
		t.Error("UserRepository.Update(), user not updated")
	}
}
