package volleynet

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

type TournamentTeam struct {
	TournamentID int     `json:"tournamentId"`
	TotalPoints  int     `json:"totalPoints"`
	Player1      *Player `json:"player1"`
	Player2      *Player `json:"player2"`
	Seed         int     `json:"seed"`
	Rank         int     `json:"rank"`
	WonPoints    int     `json:"wonPoints"`
	PrizeMoney   int     `json:"prizeMoney"`
	Deregistered bool    `json:"deregistered"`
}

type Tournament struct {
	Start            time.Time `json:"start"`
	End              time.Time `json:"end"`
	Name             string    `json:"name"`
	Season           int       `json:"season"`
	League           string    `json:"league"`
	Link             string    `json:"link"`
	EntryLink        string    `json:"entryLink"`
	ID               int       `json:"id"`
	Status           string    `json:"status"` // done, upcoming, canceled
	RegistrationOpen bool      `json:"registrationOpen"`
	Gender           string    `json:"gender"`
}

type FullTournament struct {
	Tournament
	CreatedAt       time.Time        `json:"createdAt"`
	UpdatedAt       time.Time        `json:"updatedAt"`
	Teams           []TournamentTeam `json:"teams"`
	Location        string           `json:"location"`
	HTMLNotes       string           `json:"htmlNotes"`
	Mode            string           `json:"mode"`
	MaxTeams        int              `json:"maxTeams"`
	MinTeams        int              `json:"minTeams"`
	MaxPoints       int              `json:"maxPoints"`
	EndRegistration string           `json:"endRegistration"`
	Organiser       string           `json:"organiser"`
	Phone           string           `json:"phone"`
	Email           string           `json:"email"`
	Web             string           `json:"web"`
	CurrentPoints   string           `json:"currentPoints"`
	LivescoringLink string           `json:"livescoringLink"`
	Latitude        float32          `json:"latitude"`
	Longitude       float32          `json:"longitude"`
}

var registerUrl string = "https://beach.volleynet.at/Admin/index.php?screen=Beach/Profile/TurnierAnmeldung&screen=Beach%2FProfile%2FTurnierAnmeldung&parent=0&prev=0&next=0&cur="

func (c *Client) GetUniqueWriteCode(tournamentID int) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%d", registerUrl, tournamentID), nil)
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

func (c *Client) TournamentEntry(playerName string, playerID, tournamentID int) error {
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
	form.Add("cur", strconv.Itoa(tournamentID))
	form.Add("name_b", playerName)
	form.Add("bte_per_id_b", strconv.Itoa(playerID))
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

func parseTournamentIDFromLink(link string) (int, error) {
	dashIndex := strings.LastIndex(link, "/")

	id := link[dashIndex+1:]

	return strconv.Atoi(id)
}

func parsePlayerIDFromSteckbrief(s *goquery.Selection) (int, error) {
	href := parseHref(s)

	dashIndex := strings.Index(href, "-")

	return strconv.Atoi(href[dashIndex+1:])
}

func parseFullTournamentTeams(body *goquery.Document) ([]TournamentTeam, error) {
	teams := []TournamentTeam{}

	tables := body.Find("tbody")

	for i := range tables.Nodes {
		table := tables.Eq(i)
		rows := table.Find("tr")

		if rows.First().Children().Eq(0).Text() == "Nr." {
			team := TournamentTeam{}

			for j := range rows.Nodes {
				if j == 0 {
					continue
				}

				player := &Player{}

				row := rows.Eq(j)
				columnsCount := len(row.Children().Nodes)

				columns := row.Find("td")

				for k := range columns.Nodes {
					var err error
					column := columns.Eq(k)

					if columnsCount == 5 {
						switch k {
						case 0:
							team.Rank = parseNumber(column.Text())
						case 1:
							player.ID, err = parsePlayerIDFromSteckbrief(column.Find("a"))
							player.LastName, player.FirstName, player.Login = parsePlayerName(column)
						case 2:
							player.CountryUnion = trimmedText(column)
						case 3:
							team.WonPoints = parseNumber(column.Text())
						case 4:
							team.PrizeMoney = parseNumber(column.Text())
						}
					} else if columnsCount == 4 {
						switch k {
						case 0:
							player.ID, err = parsePlayerIDFromSteckbrief(column.Find("a"))
							player.LastName, player.FirstName, player.Login = parsePlayerName(column)
						case 1:
							player.License = trimmedText(column)
						case 2:
							player.CountryUnion = trimmedText(column)
						case 3:
							player.TotalPoints = parseNumber(column.Text())
						}
					} else if columnsCount == 7 {
						switch k {
						case 0:
							team.Seed = parseNumber(column.Text())
						case 1:
							player.ID, err = parsePlayerIDFromSteckbrief(column.Find("a"))
							player.LastName, player.FirstName, player.Login = parsePlayerName(column)
						case 2:
							player.License = trimmedText(column)
						case 3:
							player.CountryUnion = trimmedText(column)
						case 4:
							player.TotalPoints = parseNumber(column.Text())
						case 5:
							team.TotalPoints = parseNumber(column.Text())
						case 6:
							// signout link
						}
					} else if columnsCount == 2 {
						switch k {
						case 0:
							player.ID, err = parsePlayerIDFromSteckbrief(column.Find("a"))
							player.LastName, player.FirstName, player.Login = parsePlayerName(column)
						case 1:
							player.License = trimmedText(column)
						}
					} else {
						return nil, errors.New("unknown tournament player table structure")
					}

					if err != nil {
						return nil, err
					}
				}
				if team.Player1 == nil {
					team.Player1 = player
				} else {
					team.Player2 = player
					teams = append(teams, team)
					team = TournamentTeam{}
				}

			}
			break
		}
	}

	return teams, nil
}

func parseFullTournament(html io.Reader) (*FullTournament, error) {
	doc, err := GetDocument(html)

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
				value := row.Eq(1)

				switch j {
				case 0:
					t.League = trimmedText(value)
				case 1:
					t.Start, t.End, err = parseStartEndDates(value)

					if err != nil {
						return nil, err
					}
				case 2:
					t.Location = trimmedText(row.Eq(1))
				case 3:
					t.Mode = trimmedText(row.Eq(1))
					t.MaxTeams = parseNumber(t.Mode)
				case 4:
					t.MinTeams = parseNumber(value.Text())
				case 5:
					t.MaxPoints = parseNumber(value.Text())
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

	if len(t.Teams) > 0 {
		t.Status = getTournamentStatusFromTeam(t.Teams[0])
	} else {
		t.Status = getTournamentStatusFromStartDate(t.Start)
	}

	return t, nil
}

func getTournamentStatusFromTeam(t TournamentTeam) string {
	if t.Rank > 0 {
		return "done"
	} else {
		return "upcoming"
	}
}

func getTournamentStatusFromStartDate(start time.Time) string {
	if time.Now().After(start) {
		return "done"
	} else {
		return "upcoming"
	}
}

func (c *Client) getTournament(id int) (*Tournament, error) {
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

func (c *Client) GetTournamentLink(t *Tournament) string {
	return c.DefaultUrl + t.Link
}

func (c *Client) GetApiTournamentLink(t *Tournament) string {
	return c.ApiUrl + t.Link
}

func (c *Client) GetTournament(id int, link string) (*FullTournament, error) {
	resp, err := http.Get(link)

	if err != nil {
		return nil, err
	}

	t, err := parseFullTournament(resp.Body)

	if err != nil {
		return nil, err
	}

	t.ID = id
	t.Link = c.GetTournamentLink(&t.Tournament)

	return t, nil
}

func (c *Client) Ladder(gender string) ([]Player, error) {
	url, err := url.Parse(c.ApiUrl)

	// gender: Herren, Damen
	url.Path += fmt.Sprintf(c.LadderPath, genderLong(gender))

	encodedURL := url.String()
	resp, err := http.Get(encodedURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return parseLadder(resp.Body)
}

func parseLadder(html io.Reader) ([]Player, error) {
	doc, err := goquery.NewDocumentFromReader(html)

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
		gender = "F"
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
				p.Rank = parseNumber(c.Text())
			case 2:
				p.FirstName, p.LastName, p.Login = parsePlayerName(c)
				p.ID, err = parsePlayerIDFromSteckbrief(c.Find("a"))
			case 3:
				break
			case 4:
				p.CountryUnion = trimmedText(c)
			case 5:
				p.Club = trimmedText(c)
			case 6:
				p.TotalPoints = parseNumber(c.Text())
			}
		}

		players = append(players, p)

	}

	return players, nil
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

func trimmedTournamentName(s *goquery.Selection) string {
	name := trimmedText(s)
	index := strings.Index(name, "- ")

	if index > 0 {
		return name[index+2:]
	}

	return name
}

func parseStartEndDates(s *goquery.Selection) (time.Time, time.Time, error) {
	a := trimmedText(s)

	dates := strings.Split(a, "-")

	dateCount := len(dates)

	if dateCount == 1 {
		startDate, err := parseDate(dates[0])

		if err != nil {
			return time.Time{}, time.Time{}, err
		}

		return startDate, startDate, nil
	} else if dateCount == 2 {
		startDate, err := parseDate(dates[0])
		if err != nil {
			return time.Time{}, time.Time{}, err
		}
		endDate, err := parseDate(dates[1])
		if err != nil {
			return time.Time{}, time.Time{}, err
		}

		return startDate, endDate, nil
	} else {
		return time.Time{}, time.Time{}, errors.New("unknown start/enddate format")
	}
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
				tournament.Start, tournament.End, err = parseStartEndDates(c)
			case 2:
				tournament.Link = parseHref(c.Find("a"))
				tournament.Name = trimmedTournamentName(c)
				tournament.ID, err = parseTournamentIDFromLink(tournament.Link)

				if err != nil {
					return nil, err
				}
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
