package volleynet

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/raphi011/scores"
)

type GroupedTournaments struct {
	Upcoming []Tournament `json:"upcoming"`
	Past     []Tournament `json:"past"`
	Played   []Tournament `json:"played"`
}

type Tournament struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Name      string    `json:"name"`
	League    string    `json:"league"`
	Link      string    `json:"link"`
	EntryLink string    `json:"entryLink"`
	ID        string    `json:"id"`
}

type Team struct {
	TotalPoints string  `json:"totalPoints"`
	Player1     *Player `json:"player1"`
	Player2     *Player `json:"player2"`
	SeedOrRank  string  `json:"seedOrRank"`
	WonPoints   string  `json:"wonPoints"`
	PrizeMoney  string  `json:"prizeMoney"`
}

type Player struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	ID           string `json:"id"`
	Birthday     string `json:"birthday"`
	TotalPoints  string `json:"totalPoints"`
	CountryUnion string `json:"countryUnion"`
	License      string `json:"license"`
	Login        string `json:"login"`
}

type FullTournament struct {
	Tournament
	Teams           []Team `json:"teams"`
	Status          string `json:"status"`
	Location        string `json:"location"`
	HTMLNotes       string `json:"htmlNotes"`
	Mode            string `json:"mode"`
	MaxTeams        int    `json:"maxTeams"`
	MinTeams        string `json:"minTeams"`
	MaxPoints       string `json:"maxPoints"`
	EndRegistration string `json:"endRegistration"`
	Organiser       string `json:"organiser"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	Web             string `json:"web"`
	CurrentPoints   string `json:"currentPoints"`
	LivescoringLink string `json:"livescoringLink"`
}

var registerUrl string = "https://beach.volleynet.at/Admin/index.php?screen=Beach/Profile/TurnierAnmeldung&screen=Beach%2FProfile%2FTurnierAnmeldung&parent=0&prev=0&next=0&cur="

func (c *Client) GetUniqueWriteCode(tournamentID string) (string, error) {
	req, err := http.NewRequest("GET", registerUrl+tournamentID, nil)
	req.Header.Add("Cookie", c.Cookie)
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return "", nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return "", nil
	}

	input := doc.Find("input[name='XX_unique_write_XXBeach/Profile/TurnierAnmeldung']")

	val, exists := input.Attr("value")

	if !exists {
		return "", errors.New("writeCode not found")
	}

	return val, nil
}

func (c *Client) TournamentEntry(playerName, playerID, tournamentID string) error {
	form := url.Values{}

	code, err := c.GetUniqueWriteCode(tournamentID)

	if err != nil {
		return err
	}

	form.Add("action", "Beach/Profile/TurnierAnmeldung")
	form.Add("XX_unique_write_XXBeach/Profile/TurnierAnmeldung", code)
	form.Add("parent", "0")
	form.Add("prev", "0")
	form.Add("next", "0")
	form.Add("cur", tournamentID)
	form.Add("name_b", playerName)
	form.Add("bte_per_id_b", playerID)
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

func parseTournamentIDFromLink(link string) string {
	dashIndex := strings.LastIndex(link, "/")

	return link[dashIndex+1:]
}

func parsePlayerIDFromSteckbrief(s *goquery.Selection) string {
	href := parseHref(s)

	dashIndex := strings.Index(href, "-")

	return href[dashIndex+1:]
}

func parseFullTournamentTeams(body *goquery.Document) ([]Team, error) {
	teams := []Team{}

	tables := body.Find("tbody")

	for i := range tables.Nodes {
		table := tables.Eq(i)
		rows := table.Find("tr")

		if rows.First().Children().Eq(0).Text() == "Nr." {
			team := Team{}

			for j := range rows.Nodes {
				if j == 0 {
					continue
				}

				player := &Player{}

				row := rows.Eq(j)
				columnsCount := len(row.Children().Nodes)

				columns := row.Find("td")

				for k := range columns.Nodes {
					column := columns.Eq(k)

					if columnsCount == 5 {
						switch k {
						case 0:
							team.SeedOrRank = trimmedText(column)
						case 1:
							player.ID = parsePlayerIDFromSteckbrief(column.Find("a"))
							player.LastName, player.FirstName, player.Login = parsePlayerName(column)
						case 2:
							player.CountryUnion = trimmedText(column)
						case 3:
							team.WonPoints = trimmedText(column)
						case 4:
							team.PrizeMoney = trimmedText(column)
						}
					} else if columnsCount == 4 {
						switch k {
						case 0:
							player.ID = parsePlayerIDFromSteckbrief(column.Find("a"))
							player.LastName, player.FirstName, player.Login = parsePlayerName(column)
						case 1:
							player.License = trimmedText(column)
						case 2:
							player.CountryUnion = trimmedText(column)
						case 3:
							player.TotalPoints = trimmedText(column)
						}
					} else if columnsCount == 7 {
						switch k {
						case 0:
							team.SeedOrRank = trimmedText(column)
						case 1:
							player.LastName, player.FirstName, player.Login = parsePlayerName(column)
						case 2:
							player.License = trimmedText(column)
						case 3:
							player.CountryUnion = trimmedText(column)
						case 4:
							player.TotalPoints = trimmedText(column)
						case 5:
							team.TotalPoints = trimmedText(column)
						case 6:
							// signout link
						}
					} else if columnsCount == 2 {
						switch k {
						case 0:
							player.ID = parsePlayerIDFromSteckbrief(column.Find("a"))
							player.LastName, player.FirstName, player.Login = parsePlayerName(column)
						case 1:
							player.License = trimmedText(column)
						}
					} else {
						return nil, errors.New("unknown tournament player table structure")
					}
				}
				if team.Player1 == nil {
					team.Player1 = player
				} else {
					team.Player2 = player
					teams = append(teams, team)
					team = Team{}
				}

			}
			break
		}
	}

	return teams, nil
}

func parseFullTournament(html io.Reader) (*FullTournament, error) {
	doc, err := GetDocument(html)
	htmlString, _ := doc.Html()
	log.Print(htmlString)

	t := &FullTournament{}

	name := doc.Find("h2").Text()

	htmlNotes := doc.Find(".extrainfo")

	if htmlNotes.Find("iframe").Length() > 0 {
		t.HTMLNotes = "Cannot display these notes yet."
	} else {
		t.HTMLNotes, _ = htmlNotes.Html()
	}

	t.Name = name

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
					dateString := row.Eq(1).Text()
					t.StartDate, err = parseDate(dateString)
				case 2:
					t.Location = trimmedText(row.Eq(1))
				case 3:
					t.Mode = trimmedText(row.Eq(1))
					t.MaxTeams = parseNumber(t.Mode)
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

	t.Teams, err = parseFullTournamentTeams(doc)

	if err != nil {
		return nil, err
	}

	return t, nil
}

var numberRegex = regexp.MustCompile("\\d+")

func parseNumber(text string) int {
	nrstr := numberRegex.FindString(text)
	nr, err := strconv.Atoi(nrstr)

	if err != nil {
		return -1
	}

	return nr
}

func (c *Client) getTournament(id string) (*Tournament, error) {
	games, err := c.AllTournaments("M", "AMATEUR TOUR", "2018")

	if err != nil {
		return nil, err
	}

	for _, t := range games {
		if t.ID == id {
			return &t, nil
		}
	}

	return nil, scores.ErrorNotFound
}

func (c *Client) getTournamentLink(t *Tournament) string {
	return c.DefaultUrl + t.Link
}

func (c *Client) getApiTournamentLink(t *Tournament) string {
	return c.ApiUrl + t.Link
}

func (c *Client) GetTournament(id string) (*FullTournament, error) {
	tournament, err := c.getTournament(id)

	if err != nil {
		return nil, err
	}
	link := c.getApiTournamentLink(tournament)

	resp, err := http.Get(link)

	if err != nil {
		return nil, err
	}

	t, err := parseFullTournament(resp.Body)

	if err != nil {
		return nil, err
	}

	t.ID = parseTournamentIDFromLink(link)
	t.Link = c.getTournamentLink(tournament)

	return t, nil
}

func (c *Client) AllTournaments(gender, league, year string) ([]Tournament, error) {
	url, err := url.Parse(c.ApiUrl)

	if err != nil {
		return nil, err
	}

	url.Path += fmt.Sprintf(c.AmateurPath, league, league, gender, year)

	encodedURL := url.String()
	resp, err := http.Get(encodedURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return parseTournaments(resp.Body)
}

func trimmedText(s *goquery.Selection) string {
	return strings.TrimSpace(s.Text())
}

func trimmedTournamentName(s *goquery.Selection) string {
	name := trimmedText(s)
	index := strings.Index(name, "- ")

	if index > 0 {
		return name[index+2:]
	}

	return name
}

func parseHref(anchor *goquery.Selection) string {
	href, _ := anchor.Attr("href")

	return href
}

func parseDate(dateString string) (time.Time, error) {
	date, err := time.Parse("02.01.2006", strings.TrimSpace(dateString))
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to parse date '%s'", dateString)
	}

	return date, nil
}

func parseStartEndDates(s *goquery.Selection) (time.Time, time.Time, error) {
	a := trimmedText(s)

	dates := strings.Split(a, "-")

	if len(dates) != 2 {
		return time.Time{}, time.Time{}, errors.New("unknown start/enddate format")
	}

	startDate, err := parseDate(dates[0])
	endDate, err1 := parseDate(dates[1])

	if err1 != nil {
		err = err1
	}

	return startDate, endDate, err
}

func parseTournaments(html io.Reader) ([]Tournament, error) {
	doc, err := goquery.NewDocumentFromReader(html)

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
				tournament.StartDate, tournament.EndDate, err = parseStartEndDates(c)
			case 2:
				tournament.Link = parseHref(c.Find("a"))
				tournament.Name = trimmedTournamentName(c)
				tournament.ID = parseTournamentIDFromLink(tournament.Link)
			case 3:
				tournament.League = trimmedText(c)
			case 4:
				tournament.EntryLink = parseHref(c.Find("a"))
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
