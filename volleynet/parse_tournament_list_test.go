package volleynet

import (
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func mustParseDate(date string) time.Time {
	result, err := time.Parse("02.01.2006", date)

	if err != nil {
		panic("date could not be parsed")
	}

	return result
}

func Test_tournament_list(t *testing.T) {
	response, _ := os.Open("testdata/tournament-list-amateur.html")

	tournaments, err := parseTournamentList(response, "http://example.com")

	if err != nil {
		t.Errorf("parseTournaments() err: %s", err)
	}

	if !cmp.Equal(tournaments, tournamentListAmateur) {
		t.Errorf("parseTournamentList() err: mismatch of tournament list")
	}
}

var tournamentListAmateur = []Tournament{
	Tournament{
		Name:             "Herren Beachvolley Wien Summer Opening",
		Start:            mustParseDate("21.04.2018"),
		End:              mustParseDate("21.04.2018"),
		ID:               21908,
		Season:           2018,
		RegistrationOpen: false,
		Gender:           "M",
		Status:           "upcoming",
		League:           "AMATEUR TOUR",
		Phase:            "ABV Tour AMATEUR 1",
		Link:             "http://example.com/beach/bewerbe/AMATEUR TOUR/phase/ABV Tour AMATEUR 1/sex/M/saison/2018/cup/21908",
	},
	Tournament{
		Name:             "Herren Graz",
		Start:            mustParseDate("21.04.2018"),
		End:              mustParseDate("22.04.2018"),
		ID:               21880,
		Season:           2018,
		RegistrationOpen: false,
		Gender:           "M",
		Status:           "upcoming",
		League:           "AMATEUR TOUR",
		Phase:            "ABV Tour AMATEUR 1",
		Link:             "http://example.com/beach/bewerbe/AMATEUR TOUR/phase/ABV Tour AMATEUR 1/sex/M/saison/2018/cup/21880",
	},
	Tournament{
		Name:             "Herren Beachvolley Grieskirchen",
		Start:            mustParseDate("06.05.2018"),
		End:              mustParseDate("06.05.2018"),
		ID:               22055,
		Season:           2018,
		RegistrationOpen: true,
		Gender:           "M",
		Status:           "upcoming",
		League:           "AMATEUR TOUR",
		Phase:            "ABV Tour AMATEUR 1",
		Link:             "http://example.com/beach/bewerbe/AMATEUR TOUR/phase/ABV Tour AMATEUR 1/sex/M/saison/2018/cup/22055",
		EntryLink:        "https://beach.volleynet.at/Anmelden/21617-22055-00",
	},
	Tournament{
		Name:             "Herren Innsbruck",
		Start:            mustParseDate("01.05.2018"),
		End:              mustParseDate("01.05.2018"),
		ID:               21938,
		Season:           2018,
		RegistrationOpen: false,
		Gender:           "M",
		Status:           "canceled",
		League:           "AMATEUR TOUR",
		Phase:            "ABV Tour AMATEUR 1",
		Link:             "http://example.com/beach/bewerbe/AMATEUR TOUR/phase/ABV Tour AMATEUR 1/sex/M/saison/2018/cup/21938",
	},
}
