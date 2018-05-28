package volleynet

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PlayerInfo struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Login     string `json:"login"`
	ID        int    `json:"id"`
	Birthday  string `json:"birthday"`
}

type Player struct {
	PlayerInfo
	Gender       string `json:"gender"`
	TotalPoints  int    `json:"totalPoints"`
	Rank         int    `json:"rank"`
	Club         string `json:"club"`
	CountryUnion string `json:"countryUnion"`
	License      string `json:"license"`
}

var lastNameRegex = regexp.MustCompile("\\p{Lu}+\\b")
var firstNameRegex = regexp.MustCompile("\\p{Lu}\\p{Ll}+\\b")

func parsePlayerName(c *goquery.Selection) (string, string, string) {
	playerName := c.Text()

	lastName := strings.Join(lastNameRegex.FindAllString(playerName, -1), " ")
	firstName := strings.Join(firstNameRegex.FindAllString(playerName, -1), " ")

	return strings.Title(firstName), strings.Title(lastName), strings.Join([]string{firstName, lastName}, ".")
}

func parsePlayerID(s *goquery.Selection) (int, error) {
	href := parseHref(s)

	if href == "" {
		return -1, errors.New("ID not found")
	}

	re, err := regexp.Compile("\\(([0-9]*),")

	if err != nil {
		return -1, err
	}

	id := re.FindStringSubmatch(href)

	return strconv.Atoi(id[1])
}

func parsePlayers(html io.Reader) ([]PlayerInfo, error) {
	players := []PlayerInfo{}
	doc, err := goquery.NewDocumentFromReader(html)

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

func (c *Client) SearchPlayers(firstName, lastName, birthday string) ([]PlayerInfo, error) {
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
