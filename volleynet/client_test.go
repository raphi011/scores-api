package volleynet

import (
	"testing"
)

func Test_searchPlayers(t *testing.T) {
	c := DefaultClient()
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

	c := DefaultClient()
	err := c.Login("", "")

	if err != nil {
		t.Error(err)
	}
}
