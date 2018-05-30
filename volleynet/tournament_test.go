package volleynet

import (
	"errors"
	"os"
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

func Test_done_tournament(t *testing.T) {
	response, _ := os.Open("testdata/done.html")

	tournament, err := parseFullTournament(response)

	if err != nil {
		t.Errorf("parseFullTournament() err: %s", err)
	}

	if tournament.Status != "done" {
		t.Errorf("parseFullTournament(), want .Status = 'done', got %s", tournament.Status)
	}
}

func Test_upcoming_tournament(t *testing.T) {
	response, _ := os.Open("testdata/upcoming.html")

	tournament, err := parseFullTournament(response)

	if err != nil {
		t.Errorf("parseFullTournament() err: %s", err)
	}

	if tournament.Status != "upcoming" {
		t.Errorf("parseFullTournament(), want .Status = 'upcoming', got %s", tournament.Status)
	}
}

func Test_tournament_list(t *testing.T) {
	response, _ := os.Open("testdata/tournament-list.html")

	tournaments, err := parseTournaments(response)

	if err != nil {
		t.Errorf("parseTournaments() err: %s", err)
	}

	if len(tournaments) != 54 {
		t.Errorf("parseTournaments(), want len(tournaments) = 54, got %v", len(tournaments))
	}
}
