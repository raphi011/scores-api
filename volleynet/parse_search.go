package volleynet

import "io"

// PlayerInfo contains all player information that the search player api returns
type PlayerInfo struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
}

func parsePlayers(html io.Reader) ([]PlayerInfo, error) {
	players := []PlayerInfo{}
	doc, err := parseHTML(html)

	if err != nil {
		return nil, err
	}

	rows := doc.Find("tr")

	for i := range rows.Nodes {
		playerFound := false

		r := rows.Eq(i)

		player := PlayerInfo{}

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
				player.Birthday = c.Text()
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
