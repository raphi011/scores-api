package parse

import (
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/raphi011/scores"
	"github.com/raphi011/scores/volleynet"
)

func TestParseLadderTest(t *testing.T) {
	response, _ := os.Open("../testdata/ladder-men.html")

	expected := []*volleynet.Player{
		&volleynet.Player{
			PlayerInfo: volleynet.PlayerInfo{
				FirstName: "Alexander",
				LastName:  "Horst",
				TrackedModel: scores.TrackedModel{ Model: scores.Model{ID: 246 }},
				Birthday:  time.Time{},
			},
			Rank:         1,
			Club:         "Badener Beach Volleyballverein",
			CountryUnion: "NÖVV",
			TotalPoints:  2575,
			Gender:       "M",
		},
		&volleynet.Player{
			PlayerInfo: volleynet.PlayerInfo{
				FirstName: "Clemens",
				LastName:  "Doppler",
				TrackedModel: scores.TrackedModel{ Model: scores.Model{ID: 1050 }},
				Birthday:  time.Time{},
			},
			Rank:         1,
			Club:         "Badener Beach Volleyballverein",
			CountryUnion: "NÖVV",
			TotalPoints:  2575,
			Gender:       "M",
		},
		&volleynet.Player{
			PlayerInfo: volleynet.PlayerInfo{
				FirstName: "Robin Valentin",
				LastName:  "Seidl",
				TrackedModel: scores.TrackedModel{ Model: scores.Model{ID: 5626 }},
				Birthday:  time.Time{},
			},
			Rank:         3,
			Club:         "ABC Wörthersee",
			CountryUnion: "KVV",
			TotalPoints:  1900,
			Gender:       "M",
		},
		&volleynet.Player{
			PlayerInfo: volleynet.PlayerInfo{
				FirstName: "Martin",
				LastName:  "Ermacora",
				TrackedModel: scores.TrackedModel{ Model: scores.Model{ID: 6656 }},
				Birthday:  time.Time{},
			},
			Rank:         4,
			Club:         "My BeachEvent",
			CountryUnion: "ÖVV",
			TotalPoints:  1710,
			Gender:       "M",
		},
	}

	ladder, err := Ladder(response)

	if err != nil {
		t.Errorf("Ladder() err: %s", err)
	}

	if !cmp.Equal(ladder, expected) {
		t.Errorf("Ladder(): %s", cmp.Diff(expected, ladder))
	}
}
