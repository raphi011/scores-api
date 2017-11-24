package main

import (
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, err := initDb()
	if err != nil {
		panic("Error creating db")
	}

	defer db.Close()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
