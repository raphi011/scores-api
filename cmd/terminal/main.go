package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/sqlite"
)

var dbPath = flag.String("db", "scores.db", "Path to sqlite db")

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Possible commands: createdb, seed")
		os.Exit(1)
	}

	cmd := args[0]

	switch cmd {
	case "migrate":
		migrate()
	case "seed":
		seedDb()
	default:
		flag.PrintDefaults()
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
