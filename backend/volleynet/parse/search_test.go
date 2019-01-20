package parse

import (
	"os"
	"testing"

	"github.com/raphi011/scores/test"
	"github.com/raphi011/scores/volleynet"
)

func TestParsePlayers(t *testing.T) {
	response, _ := os.Open("../testdata/search.html")

	expected := []*volleynet.PlayerInfo{
		&volleynet.PlayerInfo{
			ID: 34822,
			FirstName: "Cristopher",
			LastName:  "Pfau",
			Birthday:  mustParseDate("28.01.2000"),
		},
		&volleynet.PlayerInfo{
			ID: 50427,
			FirstName: "Hannes",
			LastName:  "Pfau",
			Birthday:  mustParseDate("22.10.2002"),
		},
		&volleynet.PlayerInfo{
			ID: 42378,
			FirstName: "Jennifer",
			LastName:  "Pfau",
			Birthday:  mustParseDate("26.08.1988"),
		},
	}

	players, err := Players(response)

	test.Check(t, "Players() err: %v", err)
	test.Compare(t, "Players():\n%s", players, expected)
}
