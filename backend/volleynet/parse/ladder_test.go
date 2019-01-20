package parse

import (
	"os"
	"testing"
	"time"

	"github.com/raphi011/scores/test"
	"github.com/raphi011/scores/volleynet"
)

func TestParseLadderTest(t *testing.T) {
	response, _ := os.Open("../testdata/ladder-men.html")

	expected := []*volleynet.Player{
		&volleynet.Player{
			ID: 246,
			FirstName: "Alexander",
			LastName:  "Horst",
			Birthday:  time.Time{},
			Rank:         1,
			Club:         "Badener Beach Volleyballverein",
			CountryUnion: "NÖVV",
			TotalPoints:  2575,
			Gender:       "M",
		},
		&volleynet.Player{
			ID: 1050,
			FirstName: "Clemens",
			LastName:  "Doppler",
			Birthday:  time.Time{},
			Rank:         1,
			Club:         "Badener Beach Volleyballverein",
			CountryUnion: "NÖVV",
			TotalPoints:  2575,
			Gender:       "M",
		},
		&volleynet.Player{
			ID: 5626,
			FirstName: "Robin Valentin",
			LastName:  "Seidl",
			Birthday:  time.Time{},
			Rank:         3,
			Club:         "ABC Wörthersee",
			CountryUnion: "KVV",
			TotalPoints:  1900,
			Gender:       "M",
		},
		&volleynet.Player{
			ID: 6656,
			FirstName: "Martin",
			LastName:  "Ermacora",
			Birthday:  time.Time{},
			Rank:         4,
			Club:         "My BeachEvent",
			CountryUnion: "ÖVV",
			TotalPoints:  1710,
			Gender:       "M",
		},
	}

	ladder, err := Ladder(response)

	test.Check(t, "Ladder() err: %v", err)
	test.Compare(t, "Ladder(): %s", ladder, expected)
}
