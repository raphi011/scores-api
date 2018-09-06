package parse

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/raphi011/scores/volleynet"
)

func TestParsePlayers(t *testing.T) {
	response, _ := os.Open("../testdata/search.html")

	expected := []volleynet.PlayerInfo{
		volleynet.PlayerInfo{
			FirstName: "Cristopher",
			LastName:  "Pfau",
			ID:        34822,
			Birthday:  mustParseDate("28.01.2000"),
		},
		volleynet.PlayerInfo{
			FirstName: "Hannes",
			LastName:  "Pfau",
			ID:        50427,
			Birthday:  mustParseDate("22.10.2002"),
		},
		volleynet.PlayerInfo{
			FirstName: "Jennifer",
			LastName:  "Pfau",
			ID:        42378,
			Birthday:  mustParseDate("26.08.1988"),
		},
	}

	players, err := Players(response)

	if err != nil {
		t.Errorf("Players() err: %s", err)
	}

	if !cmp.Equal(players, expected) {
		t.Errorf("Players(): %s", cmp.Diff(expected, players))
	}
}
