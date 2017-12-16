package main

import (
	"os"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := App{Production: os.Getenv("APP_ENV") == "production"}
	app.Initialize()
	app.Run()
}
