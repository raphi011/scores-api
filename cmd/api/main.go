package main

import (
	"flag"

	"github.com/raphi011/scores-api/cmd/api/router"
)

var version = "undefined"

func main() {
	dbProvider := flag.String("provider", "sqlite3", "DB Driver (sqlite3, mysql or postgres)")
	connectionString := flag.String("connection", "./scores.db", "provider specific connectionstring")
	gSecret := flag.String("gauth", "./client_secret.json", "Path to google oauth secret")
	mode := flag.String("mode", "production", "debug or production")
	host := flag.String("backendurl", "https://localhost", "backend url")

	flag.Parse()

	r := router.New(
		router.WithVersion(version),
		router.WithMode(*mode),
		router.WithRepository(*dbProvider, *connectionString),
		router.WithOAuth(*gSecret, *host),
		router.WithEventQueue(),
	)

	r.Run()
}
