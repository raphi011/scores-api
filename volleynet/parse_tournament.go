package volleynet

import (
	"fmt"
	"io"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Player adds additional information to `PlayerInfo`
type Player struct {
	PlayerInfo
	Gender       string `json:"gender"`
	TotalPoints  int    `json:"totalPoints"`
	Rank         int    `json:"rank"`
	Club         string `json:"club"`
	CountryUnion string `json:"countryUnion"`
	License      string `json:"license"`
}

// TournamentTeam is the current status of the team entry in a
// tournament, if the tournament is finished it may also contain
// the seed
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

// FullTournament adds additional information to `Tournament`
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

func parseFullTournament(
	html io.Reader,
	tournament Tournament) (*FullTournament, error) {

	doc, err := parseHTML(html)

	if err != nil {
		return nil, errors.Wrap(err, "parseFullTournament failed")
	}

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
					t.Location = trimmSelectionText(value)
				case 3:
					t.Mode = trimmSelectionText(value)
					t.MaxTeams = findInt(t.Mode)
				case 4:
					t.MinTeams = findInt(value.Text())
				case 5:
					t.MaxPoints = findInt(value.Text())
				case 6:
					t.EndRegistration = trimmSelectionText(value)
				case 7:
					t.Organiser = trimmSelectionText(value)
				case 8:
					t.Phone = trimmSelectionText(value)
				case 9:
					t.Email = trimmSelectionText(value)
				case 10:
					t.Web = trimmSelectionText(value)
				case 11:
					t.CurrentPoints = trimmSelectionText(value)
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
		return nil, errors.Wrap(err, "error parsing tournament teams")
	}

	return t, nil
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
					log.Warnf("error parsing player: %s", err)
					j++ // if it's not possible to parse a player, skip the entire team
					continue
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
				team.Rank = findInt(column.Text())
			case 1:
				player.ID, err = parsePlayerIDFromSteckbrief(column.Find("a"))
				player.LastName, player.FirstName, player.Login = parsePlayerName(column)
			case 2:
				player.CountryUnion = trimmSelectionText(column)
			case 3:
				team.WonPoints = findInt(column.Text())
			case 4:
				team.PrizeMoney = parseFloat(column.Text())
			}
		} else if columnsCount == 4 {
			switch k {
			case 0:
				player.ID, err = parsePlayerIDFromSteckbrief(column.Find("a"))
				player.LastName, player.FirstName, player.Login = parsePlayerName(column)
			case 1:
				player.License = trimmSelectionText(column)
			case 2:
				player.CountryUnion = trimmSelectionText(column)
			case 3:
				player.TotalPoints = findInt(column.Text())
			}
		} else if columnsCount == 7 {
			switch k {
			case 0:
				team.Seed = findInt(column.Text())
			case 1:
				player.ID, err = parsePlayerIDFromSteckbrief(column.Find("a"))
				player.LastName, player.FirstName, player.Login = parsePlayerName(column)
			case 2:
				player.License = trimmSelectionText(column)
			case 3:
				player.CountryUnion = trimmSelectionText(column)
			case 4:
				player.TotalPoints = findInt(column.Text())
			case 5:
				team.TotalPoints = findInt(column.Text())
			case 6:
				// signout link
				team.Deregistered = trimmSelectionText(column) == ""
			}
		} else if columnsCount == 2 {
			switch k {
			case 0:
				player.ID, err = parsePlayerIDFromSteckbrief(column.Find("a"))
				player.LastName, player.FirstName, player.Login = parsePlayerName(column)
			case 1:
				player.License = trimmSelectionText(column)
			}
		} else {
			return nil, fmt.Errorf("unknown tournament player table row count: %d", columnsCount)
		}

		if err != nil {
			return nil, err
		}
	}

	return
}
