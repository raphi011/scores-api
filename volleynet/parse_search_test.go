package volleynet

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParsePlayers(t *testing.T) {
	response, _ := os.Open("testdata/search.html")

	expected := []PlayerInfo{
		PlayerInfo{
			FirstName: "Cristopher",
			LastName:  "Pfau",
			ID:        34822,
			Birthday:  "28.01.2000",
		},
		PlayerInfo{
			FirstName: "Hannes",
			LastName:  "Pfau",
			ID:        50427,
			Birthday:  "22.10.2002",
		},
		PlayerInfo{
			FirstName: "Jennifer",
			LastName:  "Pfau",
			ID:        42378,
			Birthday:  "26.08.1988",
		},
	}

	players, err := parsePlayers(response)

	if err != nil {
		t.Errorf("parsePlayers() err: %s", err)
	}

	if !cmp.Equal(players, expected) {
		t.Errorf("parsePlayers(): %s", cmp.Diff(expected, players))
	}
}
