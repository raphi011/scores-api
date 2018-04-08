package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/sqlite"
)

var dbPath = flag.String("db", "scores.db", "Path to sqlite db")

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Possible commands: set-pw, seed, migrate")
		os.Exit(1)
	}

	cmd := args[0]

	switch cmd {
	case "migrate":
		migrate()
	case "seed":
		seedDb()
	case "set-pw":
		setNewPassword()
	default:
		flag.PrintDefaults()
	}
}

func getUser(us sqlite.UserService) *scores.User {
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
	db, err := sqlite.Open(*dbPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	us := sqlite.UserService{DB: db, PW: &scores.PBKDF2PasswordService{
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

func migrate() {
	db, err := sqlite.Open(*dbPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = sqlite.Migrate(db)

	if err != nil {
		fmt.Printf("Error migrating %v", err)
	}
}

func seedDb() {
	db, err := sqlite.Open(*dbPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	ps := sqlite.PlayerService{DB: db}

	ps.Create(&scores.Player{Name: "Raphi"})
	ps.Create(&scores.Player{Name: "Richie"})
	ps.Create(&scores.Player{Name: "Dominik"})
	ps.Create(&scores.Player{Name: "Lukas"})
}
