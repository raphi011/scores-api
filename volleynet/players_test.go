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
