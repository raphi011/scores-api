package scrape

import (
	"os"
	"testing"
	"time"

	"github.com/raphi011/scores-backend/test"
	"github.com/raphi011/scores-backend/volleynet"
)

func ParseMagicYear(timeString string) *time.Time {
	t := test.MustParseTimeFormat("2006 15:04", timeString+" 13:37")
	return &t
}

func TestParseLadderTest(t *testing.T) {
	response, _ := os.Open("../testdata/ladder-men.html")

	expected := []*volleynet.Player{
		{
			ID:           246,
			Birthday:     ParseMagicYear("1982"),
			FirstName:    "Alexander",
			LastName:     "Horst",
			LadderRank:   1,
			Club:         "Badener Beach Volleyballverein",
			CountryUnion: "NÖVV",
			TotalPoints:  2575,
			Gender:       "M",
		},
		{
			ID:           1050,
			Birthday:     ParseMagicYear("1980"),
			FirstName:    "Clemens",
			LastName:     "Doppler",
			LadderRank:   1,
			Club:         "Badener Beach Volleyballverein",
			CountryUnion: "NÖVV",
			TotalPoints:  2575,
			Gender:       "M",
		},
		{
			ID:           5626,
			Birthday:     ParseMagicYear("1990"),
			FirstName:    "Robin Valentin",
			LastName:     "Seidl",
			LadderRank:   3,
			Club:         "ABC Wörthersee",
			CountryUnion: "KVV",
			TotalPoints:  1900,
			Gender:       "M",
		},
		{
			ID:           6656,
			Birthday:     ParseMagicYear("1994"),
			FirstName:    "Martin",
			LastName:     "Ermacora",
			LadderRank:   4,
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
