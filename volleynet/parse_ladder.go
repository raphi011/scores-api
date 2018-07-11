package volleynet

import (
	"io"
)

func parseLadder(html io.Reader) ([]Player, error) {
	doc, err := parseHtml(html)

	if err != nil {
		return nil, err
	}

	players := []Player{}
	rows := doc.Find("tbody>tr")
	genderTitle := doc.Find("h2").Text()
	var gender string

	if genderTitle == "Herren" {
		gender = "M"
	} else {
		gender = "W"
	}

	for i := range rows.Nodes {
		r := rows.Eq(i)

		columns := r.Find("td")

		if len(columns.Nodes) != 7 {
			continue
		}

		p := Player{}
		p.Gender = gender

		for j := range columns.Nodes {
			c := columns.Eq(j)

			switch j {
			case 1:
				p.Rank = findInt(c.Text())
			case 2:
				p.FirstName, p.LastName, p.Login = parsePlayerName(c)
				p.ID, err = parsePlayerIDFromSteckbrief(c.Find("a"))
			case 3:
				break
			case 4:
				p.CountryUnion = trimmSelectionText(c)
			case 5:
				p.Club = trimmSelectionText(c)
			case 6:
				p.TotalPoints = findInt(c.Text())
			}
		}

		players = append(players, p)

	}

	return players, nil
}
