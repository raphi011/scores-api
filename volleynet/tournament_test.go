package volleynet

import (
	"errors"
	"testing"
)

func Test_upcoming_games(t *testing.T) {
	c := DefaultClient()
	tournaments, err := c.AllTournaments("M", "AMATEUR TOUR", "2018")

	if err != nil {
		t.Error(err)
	} else if len(tournaments) <= 0 {
		t.Error(errors.New("tournaments didn't return anything"))
	}
}

func Test_full_tournament(t *testing.T) {
	c := DefaultClient()
	_, err := c.GetTournament("http://www.volleynet.at/api//beach/bewerbe/AMATEUR%20TOUR/phase/ABV%20Tour%20AMATEUR%201/sex/M/saison/2018/cup/22127")

	if err != nil {
		t.Error(err)
	}
}
