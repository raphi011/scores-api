package client

import (
	"testing"

	"github.com/pkg/errors"
)

func Test_upcoming_games(t *testing.T) {
	t.Skip()

	c := Default()
	tournaments, err := c.AllTournaments("M", "AMATEUR TOUR", 2018)

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
	t.Skip() // Add Testing credentials to env

	c := Default()
	_, err := c.Login("", "")

	if err != nil {
		t.Error(err)
	}
}
