package main

import (
	"fmt"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/db/sqlite"
)

func getUser(us sqlite.UserRepository) *scores.User {
	fmt.Print("Enter the users email: ")
	for scanner.Scan() {
		email := scanner.Text()
		user, err := us.ByEmail(email)

		if err == nil {
			return user
		}

		fmt.Printf("User %s not found\n", email)
		fmt.Print("Enter the users email: ")
	}

	return nil
}

func getPassword() string {
	fmt.Print("Enter a new password: ")
	scanner.Scan()
	password := scanner.Text()
	return password
}

func setNewPassword() {
	db, err := sqlite.Open(*dbPath, "")

	if err != nil {
		fmt.Println(err)
		return
	}

	us := sqlite.UserRepository{DB: db, PW: &scores.PBKDF2PasswordRepository{
		SaltBytes:  16,
		Iterations: 10000,
	}}

	user := getUser(us)

	if user == nil {
		return
	}

	password := getPassword()

	hashedPassword, err := us.PW.HashPassword([]byte(password))

	if err != nil {
		fmt.Printf("Error hashing password %v", err)
		return
	}

	err = us.UpdatePasswordAuthentication(user.ID, hashedPassword)

	if err != nil {
		fmt.Printf("Error updating password %v", err)
	} else {
		fmt.Printf("Password has been updated")
	}
}
