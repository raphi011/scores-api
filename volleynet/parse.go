package volleynet

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/PuerkitoBio/goquery"
)

type PlayerInfo struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Login     string `json:"login"`
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

type TournamentTeam struct {
	TournamentID int     `json:"tournamentId"`
	TotalPoints  int     `json:"totalPoints"`
	Player1      *Player `json:"player1"`
	Player2      *Player `json:"player2"`
	Seed         int     `json:"seed"`
	Rank         int     `json:"rank"`
	WonPoints    int     `json:"wonPoints"`
	PrizeMoney   float32 `json:"prizeMoney"`
	Deregistered bool    `json:"deregistered"`
}

const (
	StatusUpcoming = "upcoming"
	StatusDone     = "done"
	StatusCanceled = "canceled"
)

type Tournament struct {
	Start            time.Time `json:"start"`
	End              time.Time `json:"end"`
	Name             string    `json:"name"`
	Season           int       `json:"season"`
	League           string    `json:"league"`
	SubLeague        string    `json:"subLeague"`
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

func parsePlayerIDFromSteckbrief(s *goquery.Selection) (int, error) {
	href := parseHref(s)

	dashIndex := strings.Index(href, "-")

	return strconv.Atoi(href[dashIndex+1:])
}

func parseFullTournamentTeams(body *goquery.Document, tournamentID int, gender string) ([]TournamentTeam, error) {
	teams := []TournamentTeam{}

	tables := body.Find("tbody")

	for i := range tables.Nodes {
		table := tables.Eq(i)
		rows := table.Find("tr")

		if rows.First().Children().Eq(0).Text() == "Nr." {
			team := TournamentTeam{}
			team.TournamentID = tournamentID

			for j := range rows.Nodes {
				if j == 0 {
					continue
				}

				player, err := parsePlayerRow(rows.Eq(j), &team)

				if err != nil {
					return nil, err
				}

				player.Gender = gender

				if team.Player1 == nil {
					team.Player1 = player
				} else {
					team.Player2 = player
					teams = append(teams, team)
					team = TournamentTeam{}
					team.TournamentID = tournamentID
				}

			}
		}
	}

	return teams, nil
}

func parsePlayerRow(row *goquery.Selection, team *TournamentTeam) (player *Player, err error) {
	player = &Player{}

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
				team.PrizeMoney = parseFloat(column.Text())
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
				team.Deregistered = trimmedText(column) == ""
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

	return
}

func parseFullTournament(
	html io.Reader,
	tournament Tournament) (*FullTournament, error) {

	doc, err := GetDocument(html)

	t := &FullTournament{Tournament: tournament}

	htmlNotes := doc.Find(".extrainfo")

	if htmlNotes.Find("iframe").Length() > 0 {
		t.HTMLNotes = "Cannot display these notes yet."
	} else {
		t.HTMLNotes, _ = htmlNotes.Html()
	}

	table := doc.Find("tbody")

	for i := range table.Nodes {
		r := table.Eq(i)
		rows := r.Find("tr")

		if rows.First().Children().Eq(0).Text() == "Kategorie" {
			for j := range rows.Nodes {
				row := rows.Eq(j).Children()
				value := row.Eq(1)

				switch j {
				case 2:
					t.Location = trimmedText(value)
				case 3:
					t.Mode = trimmedText(value)
					t.MaxTeams = parseNumber(t.Mode)
				case 4:
					t.MinTeams = parseNumber(value.Text())
				case 5:
					t.MaxPoints = parseNumber(value.Text())
				case 6:
					t.EndRegistration = trimmedText(value)
				case 7:
					t.Organiser = trimmedText(value)
				case 8:
					t.Phone = trimmedText(value)
				case 9:
					t.Email = trimmedText(value)
				case 10:
					t.Web = trimmedText(value)
				case 11:
					t.CurrentPoints = trimmedText(value)
				}
			}

			break
		}
	}

	t.Teams, err = parseFullTournamentTeams(
		doc,
		tournament.ID,
		tournament.Gender)

	if err != nil {
		return nil, err
	}

	return t, nil
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
		gender = "W"
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

func readUrlPart(url, start string) string {
	startIndex := strings.Index(url, start)

	if startIndex == -1 {
		return ""
	}

	url = url[startIndex+len(start):]

	endIndex := strings.Index(url, "/")

	if endIndex == -1 {
		return url
	}

	return url[:endIndex]
}

func extractTournamentLinkData(link string) Tournament {
	id, _ := strconv.Atoi(readUrlPart(link, "cup/"))
	season, _ := strconv.Atoi(readUrlPart(link, "saison/"))

	return Tournament{
		Gender:    readUrlPart(link, "sex/"),
		League:    readUrlPart(link, "bewerbe/"),
		SubLeague: readUrlPart(link, "phase/"),
		ID:        id,
		Season:    season,
		Link:      link,
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

		columns := r.Find("td")

		if len(columns.Nodes) != 5 {
			continue
		}

		column := columns.Eq(2)

		tournament := extractTournamentLinkData(parseHref(column.Find("a")))
		tournament.Name = trimmedTournamentName(column)

		column = columns.Eq(1)
		tournament.Start, tournament.End, err = parseStartEndDates(column)

		column = columns.Eq(4)
		content := trimmedText(column)
		if content == "Abgesagt" {
			tournament.Status = StatusCanceled
			tournament.RegistrationOpen = false
		} else if entryLink := column.Find("a"); entryLink.Length() == 1 {
			tournament.Status = StatusUpcoming
			tournament.EntryLink = parseHref(entryLink)
			tournament.RegistrationOpen = true
		} else {
			tournament.Status = StatusDone
			tournament.RegistrationOpen = false
		}

		tournaments = append(tournaments, tournament)
	}

	return tournaments, nil
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

var lastNameRegex = regexp.MustCompile("\\p{Lu}+\\b")
var firstNameRegex = regexp.MustCompile("\\p{Lu}\\p{Ll}+\\b")

func parsePlayerName(c *goquery.Selection) (string, string, string) {
	playerName := c.Text()

	lastName := strings.Join(lastNameRegex.FindAllString(playerName, -1), " ")
	firstName := strings.Join(firstNameRegex.FindAllString(playerName, -1), " ")

	return strings.Title(firstName), strings.Title(lastName), strings.Join([]string{firstName, lastName}, ".")
}
