package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/raphi011/scores/db/sqlite"
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
	host := os.Getenv("BACKEND_URL")

	if host == "" {
		host = "http://localhost:3000"
	}

	db, err := sqlite.Open(*dbPath)

	if err != nil {
		panic(fmt.Sprintf("Could not open DB: %v\n", err))
	}

	defer db.Close()

	googleOAuth, err := googleOAuthConfig(*gSecret, host)

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

func googleOAuthConfig(configPath, host string) (*oauth2.Config, error) {
	var credentials credentials
	file, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	redirectURL := host + "/api/auth"

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
