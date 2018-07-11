package volleynet

import "io"

type PlayerInfo struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Login     string `json:"login"`
	Birthday  string `json:"birthday"`
}

func parsePlayers(html io.Reader) ([]PlayerInfo, error) {
	players := []PlayerInfo{}
	doc, err := parseHtml(html)

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
				player.FirstName, player.LastName, player.Login = parsePlayerName(c)

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
