package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	err := initDb()
	initAuth()

	if err != nil {
		panic("Error creating db")
	}

	defer db.Close()

	router := newRouter()
	log.Fatal(router.Run())
}
