package scrape

import (
	"io"
	"time"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"

	"github.com/raphi011/scores-backend/volleynet"
)

// Ladder parses players from the ladder page.
func Ladder(html io.Reader) ([]*volleynet.Player, error) {
	doc, err := parseHTML(html)

	if err != nil {
		return nil, err
	}

	players := []*volleynet.Player{}
	rows := doc.Find("tbody>tr")
	genderTitle := doc.Find("h2").Text()
	var gender string

	if genderTitle == "Herren" {
		gender = "M"
	} else {
		gender = "W"
	}

	for i := range rows.Nodes {
		row := rows.Eq(i)

		player := parseLadderRow(row)

		if player == nil {
			continue
		}

		player.Gender = gender

		players = append(players, player)
	}

	return players, nil
}

func parseLadderRow(row *goquery.Selection) *volleynet.Player {
	columns := row.Find("td")

	if len(columns.Nodes) != 7 {
		return nil
	}

	p := &volleynet.Player{}

	for j := range columns.Nodes {
		c := columns.Eq(j)

		switch j {
		case 1:
			p.LadderRank, _ = findInt(c.Text())
		case 2:
			p.FirstName, p.LastName = parsePlayerName(c)
			p.ID, _ = parsePlayerIDFromSteckbrief(c.Find("a"))
		case 3:
			// since we only know the birth year add a 'magic' time, so that we
			// can update it as soon as we have the information.
			birthday, err := time.Parse("2006 15:04", trimmSelectionText(c)+" 13:37")

			if err != nil {
				log.Debugf("parsing birthday of player in ladder: %q", trimmSelectionText(c))
			} else {
				p.Birthday = &birthday
			}
		case 4:
			p.CountryUnion = trimmSelectionText(c)
		case 5:
			p.Club = trimmSelectionText(c)
		case 6:
			p.TotalPoints, _ = findInt(c.Text())
		}
	}

	return p
}
