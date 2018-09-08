package parse

import (
	"fmt"
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
		r := rows.Eq(i)

		player := volleynet.PlayerInfo{}

		columns := r.Find("td")

		if len(columns.Nodes) != 4 {
			continue
		}

		column := columns.Eq(1)
		player.FirstName, player.LastName = parsePlayerName(column)

		player.ID, err = parsePlayerID(column.Find("a"))

		if err != nil {
			continue
		}

		column = columns.Eq(2)
		dateString := column.Text()
		player.Birthday, err = parseDate(dateString)

		if err != nil {
			return nil, fmt.Errorf("unable to parse date '%s'", dateString)
		}

		players = append(players, player)
	}

	return players, nil
}
