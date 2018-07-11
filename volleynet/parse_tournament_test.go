package volleynet

import (
	"os"
	"testing"

	"github.com/pkg/errors"
)

func Test_upcoming_games(t *testing.T) {
	t.Skip()

	c := DefaultClient()
	tournaments, err := c.AllTournaments("M", "AMATEUR TOUR", "2018")

	if err != nil {
		t.Error(err)
	} else if len(tournaments) <= 0 {
		t.Error(errors.New("tournaments didn't return anything"))
	}
}

func Test_done_tournament(t *testing.T) {
	response, _ := os.Open("testdata/done.html")

	_, err := parseFullTournament(response, Tournament{Gender: "M", Status: "upcoming"})

	if err != nil {
		t.Errorf("parseFullTournament() err: %s", err)
	}
}

func Test_upcoming_tournament(t *testing.T) {
	response, _ := os.Open("testdata/upcoming.html")

	_, err := parseFullTournament(response, Tournament{Gender: "M", Status: "upcoming"})

	if err != nil {
		t.Errorf("parseFullTournament() err: %s", err)
	}
}
