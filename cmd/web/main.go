package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/raphi011/scores/sqlite"
)

type app struct {
	db         *sql.DB
	conf       *oauth2.Config
	production bool
}

func main() {
	dbPath := flag.String("db", "./scores.db", "Path to sqlite db")
	gSecret := flag.String("goauth", "./client_secret.json", "Path to google oauth secret")
	flag.Parse()

	production := os.Getenv("APP_ENV") == "production"

	db, err := sqlite.Open(*dbPath)

	if err != nil {
		log.Printf("Could not open DB: %v\n", err)
		os.Exit(1)
	}

	defer db.Close()

	googleOAuth, err := googleOAuthConfig(*gSecret, production)

	if err != nil {
		log.Printf("Could not read google secret: %v, continuing without google oauth\n", err)
	}

	app := app{
		production: production,
		conf:       googleOAuth,
		db:         db,
	}

	router := initRouter(app)
	router.Run()
}

type credentials struct {
	ClientID    string `json:"client_id"`
	CientSecret string `json:"client_secret"`
}

func googleOAuthConfig(configPath string, production bool) (*oauth2.Config, error) {
	var credentials credentials
	var redirectURL string
	file, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	if production {
		redirectURL = "https://scores.raphi011.com/api/auth"
	} else {
		redirectURL = "http://localhost:3000/api/auth"
	}

	json.Unmarshal(file, &credentials)
	config := &oauth2.Config{
		ClientID:     credentials.ClientID,
		ClientSecret: credentials.CientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	return config, nil
}
