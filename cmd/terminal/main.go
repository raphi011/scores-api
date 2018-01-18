package main

import (
	"flag"
	"fmt"
	"os"
	"scores-backend/sqlite"
)

func main() {
	// cmd := flag.String("cmd", "", "command to execute, available: 'createdb'")
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Possible commands: 'createdb'")
		os.Exit(1)
	}

	cmd := args[0]

	switch cmd {
	case "createdb":
		createDb()
	}
}

func createDb() {
	db, err := sqlite.Open("test")

	if err != nil {
		fmt.Println(err)
		return
	}
	db.Close()
}
