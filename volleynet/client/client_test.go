package client

import (
	"os"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func Test_upcoming_games(t *testing.T) {
	t.Skip()

	c := Default()
	tournaments, err := c.Tournaments("M", "AMATEUR TOUR", 2018)

	if err != nil {
		t.Error(err)
	} else if len(tournaments) <= 0 {
		t.Error(errors.New("tournaments didn't return anything"))
	}
}

func Test_searchPlayers(t *testing.T) {
	c := Default()
	players, err := c.SearchPlayers("Lukas", "Wimmer", "")

	if err != nil {
		t.Error(err)
	}

	if len(players) <= 0 {
		t.Errorf("searchPlayers(), want len(players) > 0, got %v", len(players))
	}
}

func Test_login(t *testing.T) {
	user := os.Getenv("VOLLEYNET_LOGIN_USER")
	password := os.Getenv("VOLLEYNET_LOGIN_PASSWORD")

	if user == "" || password == "" {
		t.Skip("must set VOLLEYNET_LOGIN_USER and VOLLEYNET_LOGIN_PASSWORD to run this test")
	}

	c := Default()
	result, err := c.Login(user, password)

	if err != nil {
		t.Error(err)
	}

	if result.FirstName == "" || !strings.Contains(user, result.FirstName) {
		t.Error("login(), should return the logged in user")
	}
}
