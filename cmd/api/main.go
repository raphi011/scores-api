package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"golang.org/x/oauth2"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/repo"
	"github.com/raphi011/scores/repo/sqlite"
)

type app struct {
	conf       *oauth2.Config
	services   *scores.Services
	production bool
}

var version = "undefined"

func main() {
	dbProvider := flag.String("provider", "sqlite3", "DB Driver (sqlite3 or mysql)")
	connectionString := flag.String("connection", "./scores.db", "Path to sqlite db")
	gSecret := flag.String("goauth", "./client_secret.json", "Path to google oauth secret")

	flag.Parse()

	production := os.Getenv("APP_ENV") == "production"
	host := os.Getenv("BACKEND_URL")

	if host == "" {
		host = "http://localhost:3000"
	}

	services, repoClose, err := createServices(*dbProvider, *connectionString)

	if err != nil {
		log.Fatal("Could not initialize services: %s", err)
	}

	defer repoClose()

	googleOAuth, err := googleOAuthConfig(*gSecret, host)

	if err != nil {
		log.Printf("Could not read google secret: %v, continuing without google oauth\n", err)
	}

	app := app{
		production: production,
		services:   services,
		conf:       googleOAuth,
	}

	router := initRouter(app)
	router.Run()
}

func createServices(provider string, connectionString string) (*scores.Services, func(), error) {
	var repos *repo.Repositories
	var services *scores.Services
	var closerFunc func()
	var err error

	switch provider {
	case "sqlite3":
		fallthrough
	case "mysql":
		repos, closerFunc, err = sqlite.Create(provider, connectionString)
	default:
		return nil, nil, fmt.Errorf("invalid repo provider %q", provider)
	}

	if err != nil {
		return nil, nil, err
	}

	passwordService := &scores.PBKDF2PasswordService{
		SaltBytes:  16,
		Iterations: 10000,
	}

	groupService := scores.NewGroupService(
		repos.Group,
		repos.Match,
		repos.Player,
		repos.Statistic,
	)

	userService := &scores.UserService{
		Repository: repos.User,
		Password:   passwordService,
	}

	services = &scores.Services{
		Group:    groupService,
		Password: passwordService,
		User:     userService,

		Match:     repos.Match,
		Statistic: repos.Statistic,
		Team:      repos.Team,
		Player:    repos.Player,
		Volleynet: repos.Volleynet,
	}

	return services, closerFunc, nil
}
