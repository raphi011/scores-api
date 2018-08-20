package volleynet

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseLadderTest(t *testing.T) {
	response, _ := os.Open("testdata/ladder-men.html")

	expected := []Player{
		Player{
			PlayerInfo: PlayerInfo{
				FirstName: "Alexander",
				LastName:  "Horst",
				ID:        246,
				Birthday:  "",
			},
			Rank:         1,
			Club:         "Badener Beach Volleyballverein",
			CountryUnion: "NÖVV",
			TotalPoints:  2575,
			Gender:       "M",
		},
		Player{
			PlayerInfo: PlayerInfo{
				FirstName: "Clemens",
				LastName:  "Doppler",
				ID:        1050,
				Birthday:  "",
			},
			Rank:         1,
			Club:         "Badener Beach Volleyballverein",
			CountryUnion: "NÖVV",
			TotalPoints:  2575,
			Gender:       "M",
		},
		Player{
			PlayerInfo: PlayerInfo{
				FirstName: "Robin Valentin",
				LastName:  "Seidl",
				ID:        5626,
				Birthday:  "",
			},
			Rank:         3,
			Club:         "ABC Wörthersee",
			CountryUnion: "KVV",
			TotalPoints:  1900,
			Gender:       "M",
		},
		Player{
			PlayerInfo: PlayerInfo{
				FirstName: "Martin",
				LastName:  "Ermacora",
				ID:        6656,
				Birthday:  "",
			},
			Rank:         4,
			Club:         "My BeachEvent",
			CountryUnion: "ÖVV",
			TotalPoints:  1710,
			Gender:       "M",
		},
	}

	ladder, err := parseLadder(response)

	if err != nil {
		t.Errorf("parsePlayers() err: %s", err)
	}

	if !cmp.Equal(ladder, expected) {
		t.Errorf("parsePlayers(): %s", cmp.Diff(expected, ladder))
	}
}
