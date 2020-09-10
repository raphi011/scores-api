package main

import (
	"flag"

	"github.com/raphi011/scores-api/cmd/api/app"
)

var version = "undefined"

func main() {
	dbProvider := flag.String("provider", "sqlite3", "DB Driver (sqlite3 or postgres)")
	connectionString := flag.String("connection", "./scores.db", "provider specific connectionstring")
	gSecret := flag.String("gauth", "./client_secret.json", "Path to google oauth secret")
	mode := flag.String("mode", "production", "debug or production")
	host := flag.String("backendurl", "https://localhost", "backend url")

	flag.Parse()

	r := app.New(
		app.WithVersion(version),
		app.WithMode(*mode),
		app.WithRepository(*dbProvider, *connectionString),
		app.WithCron(),
		app.WithOAuth(*gSecret, *host),
		app.WithEventQueue(),
	)

	r.Run()
}
