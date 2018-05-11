package volleynet

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Tournament struct {
	StartDate string
	EndDate   string
	Name      string
	League    string
	Link      string
	EntryLink string
	ID        string
}

type FullTournament struct {
	Tournament
	Players []Player
	Status  string
	Notes   string
}

func (c *Client) TournamentEntry(playerID, tournamentID string) error {
	form := url.Values{}

	form.Add("action", "Beach/Profile/TurnierAnmeldung")
	form.Add("XX_unique_write_XXBeach/Profile/TurnierAnmeldung", "0.93754300 1525810822")
	form.Add("parent", "21617")
	form.Add("prev", "0")
	form.Add("next", "0")
	form.Add("cur", tournamentID)
	// form.Add("name_b", playerName)
	form.Add("bte_per_id_b", playerID) // 18068
	form.Add("submit", "Anmelden")

	req, err := http.NewRequest("POST", c.PostUrl, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", c.Cookie)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return errors.New("entry did not work")
	}

	return nil
}

func (c *Client) GetTournament(tournamentID string) (*FullTournament, error) {
	// resp, err := http.Get("")

	return nil, nil
}

func (c *Client) UpcomingTournaments() ([]Tournament, error) {
	resp, err := http.Get(c.ApiUrl + c.AmateurPath)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return parseTournaments(resp.Body)
}

func trimmedText(s *goquery.Selection) string {
	return strings.TrimSpace(s.Text())
}

func parseHref(anchor *goquery.Selection) string {
	href, _ := anchor.Attr("href")

	return href
}

func parseDates(s *goquery.Selection) (string, string) {
	a := trimmedText(s)

	dates := strings.Split(a, "-")

	if len(dates) != 2 {
		return "", ""
	}

	return strings.TrimSpace(dates[0]), strings.TrimSpace(dates[1])
}

func parseEntryLink(href string) (string, error) {

	index := strings.LastIndex(href, "/")

	if index >= 0 {
		idPart := href[index+1:]
		ids := strings.Split(idPart, "-")

		if len(ids) != 3 {
			return "", errors.New("Malformed entry link")
		}

		return ids[1], nil
	}

	return "", errors.New("Malformed entry link")
}

func parseTournaments(html io.Reader) ([]Tournament, error) {
	doc, err := goquery.NewDocumentFromReader(html)
	text, _ := doc.Html()
	log.Print(text)

	if err != nil {
		return nil, err
	}

	tournaments := []Tournament{}

	rows := doc.Find("tbody>tr")

	for i := range rows.Nodes {
		r := rows.Eq(i)

		tournament := Tournament{}

		columns := r.Find("td")

		if len(columns.Nodes) != 5 {
			continue
		}

		for j := range columns.Nodes {
			c := columns.Eq(j)

			switch j {
			case 1:
				tournament.StartDate, tournament.EndDate = parseDates(c)
			case 2:
				tournament.Link = parseHref(c.Find("a"))
				tournament.Name = trimmedText(c)
			case 3:
				tournament.League = trimmedText(c)
			case 4:
				tournament.EntryLink = parseHref(c.Find("a"))
				tournament.ID, err = parseEntryLink(tournament.EntryLink)

				if err != nil {
					log.Printf("Error parsing ID, err: %v\n", err)
					break
				}
			}
		}

		if err == nil {
			tournaments = append(tournaments, tournament)
		} else {
			err = nil
		}
	}

	return tournaments, nil
}
