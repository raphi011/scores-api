package main

import (
	"fmt"

	"github.com/raphi011/scores/db/sqlite"
)

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
