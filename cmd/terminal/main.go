package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var dbPath = flag.String("db", "scores.db", "Path to sqlite db")

var scanner = bufio.NewScanner(os.Stdin)

var version = "undefined"

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Possible commands: set-pw, seed, migrate, scrape")
		os.Exit(1)
	}

	cmd := args[0]

	switch cmd {
	case "scrape":
		scrape()
	case "migrate":
		migrate()
	case "seed":
		seedDb()
	case "set-pw":
		setNewPassword()
	default:
		flag.PrintDefaults()
	}
}
