package volleynet

import "testing"

func Test_upcoming_games(t *testing.T) {
	c := DefaultClient()
	_, err := c.UpcomingTournaments()

	if err != nil {
		t.Error(err)
	}
}
