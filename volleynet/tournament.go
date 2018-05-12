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
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Name      string `json:"name"`
	League    string `json:"league"`
	Link      string `json:"link"`
	EntryLink string `json:"entryLink"`
	ID        string `json:"id"`
}

type FullTournament struct {
	Tournament
	Players         []Player `json:"players"`
	Status          string   `json:"status"`
	Location        string   `json:"location"`
	HTMLNotes       string   `json:"htmlNotes"`
	Mode            string   `json:"mode"`
	MinTeams        string   `json:"minTeams"`
	MaxPoints       string   `json:"maxPoints"`
	EndRegistration string   `json:"endRegistration"`
	Organiser       string   `json:"organiser"`
	Phone           string   `json:"phone"`
	Email           string   `json:"email"`
	Web             string   `json:"web"`
	CurrentPoints   string   `json:"currentPoints"`
	LivescoringLink string   `json:"livescoringLink"`
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

func parseIDFromSteckbrief(s *goquery.Selection) string {
	href := parseHref(s)

	dashIndex := strings.Index(href, "-")

	return href[dashIndex+1:]
}

func parseFullTournamentPlayers(doc *goquery.Document) ([]Player, error) {
	players := []Player{}

	tables := doc.Find("tbody")

	for i := range tables.Nodes {
		table := tables.Eq(i)
		rows := table.Find("tr")

		if rows.First().Children().Eq(0).Text() == "Nr." {
			player := Player{}
			for j := range rows.Nodes {
				if j == 0 {
					continue
				}
				row := rows.Eq(j)
				columnsCount := len(row.Children().Nodes)
				done := columnsCount == 5

				columns := row.Find("td")

				for k := range columns.Nodes {
					column := columns.Eq(k)

					if done {
						switch k {
						case 0:
							// rank
						case 1:
							player.ID = parseIDFromSteckbrief(column.Find("a"))
							player.LastName, player.FirstName = parsePlayerName(column)
						case 2:
							player.CountryUnion = trimmedText(column)
						case 3:
							// won points
						case 4:
							// prize money
						}
					} else {
						switch k {
						case 0: // seed
						case 1:
							player.LastName, player.FirstName = parsePlayerName(column)
						case 2: // license
						case 3:
							player.CountryUnion = trimmedText(column)
						case 4:
							player.TotalPoints = trimmedText(column)
						case 5:
							// team total points
						case 6:
							// signout link
						}
					}
				}
				players = append(players, player)

			}
			break
		}
	}

	return players, nil
}

func parseFullTournament(html io.Reader) (*FullTournament, error) {
	doc, err := goquery.NewDocumentFromReader(html)

	t := &FullTournament{}

	name := doc.Find(".ajax>h2").Text()

	htmlNotes, err := doc.Find(".extrainfo").Html()

	if err != nil {
		return nil, err
	}

	t.Name = name
	t.HTMLNotes = htmlNotes

	table := doc.Find("tbody")

	for i := range table.Nodes {
		r := table.Eq(i)
		rows := r.Find("tr")

		if rows.First().Children().Eq(0).Text() == "Kategorie" {
			for j := range rows.Nodes {
				row := rows.Eq(j).Children()

				switch j {
				case 0:
					t.League = trimmedText(row.Eq(1))
				case 1:
					t.StartDate = trimmedText(row.Eq(1))
				case 2:
					t.Location = trimmedText(row.Eq(1))
				case 3:
					t.Mode = trimmedText(row.Eq(1))
				case 4:
					t.MinTeams = trimmedText(row.Eq(1))
				case 5:
					t.MaxPoints = trimmedText(row.Eq(1))
				case 6:
					t.EndRegistration = trimmedText(row.Eq(1))
				case 7:
					t.Organiser = trimmedText(row.Eq(1))
				case 8:
					t.Phone = trimmedText(row.Eq(1))
				case 9:
					t.Email = trimmedText(row.Eq(1))
				case 10:
					t.Web = trimmedText(row.Eq(1))
				case 11:
					t.CurrentPoints = trimmedText(row.Eq(1))
				}
			}

			break
		}
	}

	t.Players, err = parseFullTournamentPlayers(doc)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) GetTournament(link string) (*FullTournament, error) {
	resp, err := http.Get(link)

	if err != nil {
		return nil, err
	}

	return parseFullTournament(resp.Body)
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
