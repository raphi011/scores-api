package main

import (
	"fmt"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/sqlite"
)

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
