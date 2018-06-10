package volleynet

import (
	"testing"
)

func Test_searchPlayers(t *testing.T) {
	c := DefaultClient()
	_, err := c.SearchPlayers("Lukas", "Wimmer", "")

	if err != nil {
		t.Error(err)
		return
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
