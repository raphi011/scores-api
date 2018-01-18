package sqlite

import (
	"scores-backend"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, _ := Open("file::memory:?mode=memory&cache=shared")
	defer ClearTables(db)

	email := "test@test.com"

	userService := UserService{DB: db}
	user, err := userService.Create(&scores.User{
		Email:           email,
		ProfileImageURL: "image.url",
	})

	if err != nil {
		t.Errorf("userService.Create() err: %s", err)
	}

	if user.ID == 0 {
		t.Errorf("userService.Create(), want ID != 0, got 0")
	}

	userID := user.ID

	user, _ = userService.ByEmail(email)

	if user.ID != userID {
		t.Errorf("userService.Create(), user not persisted")
	}
}

func TestUsers(t *testing.T) {
	db, _ := Open("file::memory:?mode=memory&cache=shared")
	defer ClearTables(db)

	userService := UserService{DB: db}
	userService.Create(&scores.User{
		Email: "test@test.at",
	})
	userService.Create(&scores.User{
		Email:           "test2@test.at",
		ProfileImageURL: "image.url",
	})

	users, err := userService.Users()

	if err != nil {
		t.Errorf("UserService.Users() err: %s", err)
	}

	userCount := len(users)
	if userCount != 2 {
		t.Errorf("len(UserService.Users()), want 2, got %d", userCount)
	}

}

func TestUpdateUser(t *testing.T) {
	db, _ := Open("file::memory:?mode=memory&cache=shared")
	defer ClearTables(db)

	email := "test@test.com"
	newEmail := "test2@test.com"

	userService := UserService{DB: db}
	user, _ := userService.Create(&scores.User{
		Email:           email,
		ProfileImageURL: "image.url",
	})

	user.Email = newEmail

	err := userService.Update(user)

	if err != nil {
		t.Errorf("UserService.Update() err: %s", err)
	}

	user, err = userService.ByEmail(newEmail)

	if err != nil || user.Email != newEmail {
		t.Error("UserService.Update(), user not updated")
	}
}
