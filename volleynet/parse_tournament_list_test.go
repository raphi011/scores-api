package volleynet

import (
	"os"
	"testing"
)

func Test_tournament_list(t *testing.T) {
	response, _ := os.Open("testdata/tournament-list.html")

	tournaments, err := parseTournamentList(response, "http://example.com")

	if err != nil {
		t.Errorf("parseTournaments() err: %s", err)
	}

	if len(tournaments) != 54 {
		t.Errorf("parseTournaments(), want len(tournaments) = 54, got %v", len(tournaments))
	}
}
