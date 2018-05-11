package volleynet

import "testing"

func Test_upcoming_games(t *testing.T) {
	c := DefaultClient()
	_, err := c.UpcomingTournaments()

	if err != nil {
		t.Error(err)
	}
}

func Test_full_tournament(t *testing.T) {
	c := DefaultClient()
	_, err := c.GetTournament("http://www.volleynet.at/beach/bewerbe/AMATEUR%20TOUR/phase/ABV%20Tour%20AMATEUR%201/sex/M/saison/2018/cup/22177")

	if err != nil {
		t.Error(err)
	}
}
