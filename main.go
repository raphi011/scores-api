package main

import (
	"os"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := App{}
	app.Initialize(os.Getenv("APP_ENV"))
	app.Run()
}
