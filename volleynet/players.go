package volleynet

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Player struct {
	FirstName    string
	LastName     string
	ID           string
	Birthday     string
	TotalPoints  string
	CountryUnion string
}

func parsePlayerName(c *goquery.Selection) (string, string) {
	fullName := strings.TrimSpace(c.Text())

	parts := strings.Split(fullName, " ")

	if len(parts) != 2 {
		return "", ""
	}

	return strings.Title(parts[0]), strings.Title(parts[1])
}

func parsePlayerID(s *goquery.Selection) (string, error) {
	href := parseHref(s)

	if href == "" {
		return "", errors.New("ID not found")
	}

	re, err := regexp.Compile("\\(([0-9]*),")

	if err != nil {
		return "", err
	}

	id := re.FindStringSubmatch(href)

	return id[1], nil
}

func parsePlayers(html io.Reader) ([]Player, error) {
	players := []Player{}
	doc, err := goquery.NewDocumentFromReader(html)
	text, _ := doc.Html()
	log.Print(text)

	if err != nil {
		return nil, err
	}

	rows := doc.Find("tr")

	for i := range rows.Nodes {
		playerFound := false

		r := rows.Eq(i)

		player := Player{}

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

func (c *Client) SearchPlayers(firstName, lastName, birthday string) ([]Player, error) {
	form := url.Values{}

	form.Add("XX_unique_write_XXAdmin/Search", "0.50981600 1525795371")
	form.Add("popup", "1")
	form.Add("add", "")
	form.Add("target", "bte_per_id_b")
	form.Add("txm_language", "de")
	form.Add("sai_id", "")
	form.Add("action", "Admin/Search")
	form.Add("submit", "Suchen")
	form.Add("search", "Person")
	form.Add("per_name", lastName)
	form.Add("per_vorname", firstName)
	form.Add("per_geburtsdatum", birthday)
	form.Add("doit", "1")
	form.Add("text", "0")

	response, err := http.PostForm(c.PostUrl, form)

	if err != nil {
		return nil, err
	}

	return parsePlayers(response.Body)
}
