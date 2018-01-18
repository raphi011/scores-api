package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	scores "scores-backend"
	"scores-backend/sqlite"
)

func main() {
	dbPath := flag.String("db", "scores.db", "Path to sqlite db")
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Possible commands: 'createdb'")
		os.Exit(1)
	}

	cmd := args[0]

	switch cmd {
	case "createdb":
		db, err := getDb(*dbPath)
		if err != nil {
			fmt.Println(err)
		}
		db.Close()
	case "seed":
		seedDb(*dbPath)
	}
}

func seedDb(path string) {
	db, err := getDb(path)

	if err != nil {
		fmt.Println(err)
	}

	ps := sqlite.PlayerService{DB: db}

	ps.Create(&scores.Player{Name: "Raphi"})
	ps.Create(&scores.Player{Name: "Richie"})
	ps.Create(&scores.Player{Name: "Dominik"})
	ps.Create(&scores.Player{Name: "Lukas"})
}

func getDb(path string) (*sql.DB, error) {
	db, err := sqlite.Open(path)

	if err != nil {
		return nil, err
	}

	return db, nil
}
