package scrape

import (
	"os"
	"testing"

	"github.com/raphi011/scores-backend/test"
)

func TestParsePlayers(t *testing.T) {
	response, _ := os.Open("../testdata/search.html")

	expected := []*PlayerInfo{
		{
			ID:        34822,
			FirstName: "Cristopher",
			LastName:  "Pfau",
			Birthday:  test.MustParseDate("28.01.2000"),
		},
		{
			ID:        50427,
			FirstName: "Hannes",
			LastName:  "Pfau",
			Birthday:  test.MustParseDate("22.10.2002"),
		},
		{
			ID:        42378,
			FirstName: "Jennifer",
			LastName:  "Pfau",
			Birthday:  test.MustParseDate("26.08.1988"),
		},
	}

	players, err := Players(response)

	test.Check(t, "Players() err: %v", err)
	test.Compare(t, "Players():\n%s", players, expected)
}
