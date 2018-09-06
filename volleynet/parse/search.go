package parse

import (
	"io"

	"github.com/raphi011/scores/volleynet"
)

func Players(html io.Reader) ([]volleynet.PlayerInfo, error) {
	players := []volleynet.PlayerInfo{}
	doc, err := parseHTML(html)

	if err != nil {
		return nil, err
	}

	rows := doc.Find("tr")

	for i := range rows.Nodes {
		playerFound := false

		r := rows.Eq(i)

		player := volleynet.PlayerInfo{}

		columns := r.Find("td")

		if len(columns.Nodes) != 4 {
			continue
		}

		for j := range columns.Nodes {
			c := columns.Eq(j)

			switch j {
			case 1:
				player.FirstName, player.LastName = parsePlayerName(c)

				player.ID, err = parsePlayerID(c.Find("a"))
				if err == nil {
					playerFound = true
				}
			case 2:
				// player.Birthday = c.Text() // TODO
			}
		}

		if playerFound {
			players = append(players, player)
		} else {
			err = nil
		}
	}

	return players, nil
}
