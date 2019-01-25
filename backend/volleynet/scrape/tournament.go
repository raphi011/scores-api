package scrape

import (
	"fmt"
	"io"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/volleynet"
)

// Tournament adds remaining details to the tournament (parsed by tournament_list)
func Tournament(
	html io.Reader,
	now time.Time,
	tournament *volleynet.TournamentInfo) (*volleynet.Tournament, error) {

	doc, err := parseHTML(html)

	if err != nil {
		return nil, errors.Wrap(err, "parseFullTournament failed")
	}

	t := &volleynet.Tournament{TournamentInfo: *tournament}

	parseTournamentNotes(doc, t)
	parseTournamentDetails(doc, t)

	err = parseFullTournamentTeams(doc, t)

	if err != nil {
		return nil, errors.Wrap(err, "error parsing tournament teams")
	}

	if len(t.Teams) == 0 && isDateAfter(now, t.End) ||
		len(t.Teams) > 0 && t.Teams[0].Result > 0 {
		t.Status = volleynet.StatusDone
	}

	return t, nil
}

func getDate(d time.Time) time.Time {
	date := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
	log.Print(date)
	return date
}

func isDateAfter(tournament, current time.Time) bool {
	return getDate(tournament).After(getDate(current))
}

func parseTournamentNotes(doc *goquery.Document, t *volleynet.Tournament) {
	htmlNotes := doc.Find(".extrainfo")

	if htmlNotes.Find("iframe").Length() > 0 {
		t.HTMLNotes = "Cannot display these notes yet."
	} else {
		t.HTMLNotes, _ = htmlNotes.Html()
	}
}

type detailsParser func(*goquery.Selection, *volleynet.Tournament)

var parseTournamentDetailsMap = map[string]detailsParser{
	"Kategorie": func(value *goquery.Selection, t *volleynet.Tournament) {
		t.SubLeague = trimmSelectionText(value)
		t.SubLeagueKey = scores.Sluggify(trimmSelectionText(value))
	},
	"Modus": func(value *goquery.Selection, t *volleynet.Tournament) {
		t.Mode = trimmSelectionText(value)
		t.MaxTeams = findInt(t.Mode)
	},
	"Teiln. Qual.": func(value *goquery.Selection, t *volleynet.Tournament) {
		// TODO: not min teams but min teams for qualification, this is misleading
		t.MinTeams = findInt(value.Text())
	},
	"Datum": func(value *goquery.Selection, t *volleynet.Tournament) {
		var err error
		t.Start, t.End, err = parseStartEndDates(value)

		if err != nil {
			log.Warnf("error parsing start/end dates from tournamentId: %d, value: %s", t.ID, value.Text())
		}
	},
	"Ort": func(value *goquery.Selection, t *volleynet.Tournament) {
		t.Location = trimSelectionHTML(value)
	},
	"Max. Punkte": func(value *goquery.Selection, t *volleynet.Tournament) {
		t.MaxPoints = findInt(value.Text())
	},
	"Veranstalter": func(value *goquery.Selection, t *volleynet.Tournament) {
		t.Organiser = trimmSelectionText(value)
	},
	"Telefon": func(value *goquery.Selection, t *volleynet.Tournament) {
		t.Phone = trimmSelectionText(value)
	},
	"EMail": func(value *goquery.Selection, t *volleynet.Tournament) {
		t.Email = trimmSelectionText(value)
	},
	"Web": func(value *goquery.Selection, t *volleynet.Tournament) {
		t.Website = trimmSelectionText(value)
	},
	"Vorläufige Punkte": func(value *goquery.Selection, t *volleynet.Tournament) {
		t.CurrentPoints = trimmSelectionText(value)
	},
	"Nennschluss": func(value *goquery.Selection, t *volleynet.Tournament) {
		endRegistration, _ := parseDate(value.Text())
		if endRegistration.IsZero() {
			t.EndRegistration = nil
		} else {
			t.EndRegistration = &endRegistration
		}
	},
}

func parseTournamentDetails(doc *goquery.Document, t *volleynet.Tournament) {
	table := doc.Find("tbody")

	for i := range table.Nodes {
		r := table.Eq(i)
		rows := r.Find("tr")

		firstColumnName := rows.First().Children().Eq(0).Text()

		if _, ok := parseTournamentDetailsMap[firstColumnName]; ok {
			for j := range rows.Nodes {
				row := rows.Eq(j).Children()
				columnName := row.Eq(0).Text()
				value := row.Eq(1)

				if parser, ok := parseTournamentDetailsMap[columnName]; ok {
					parser(value, t)
				}
			}
		}
	}
}

func parseFullTournamentTeams(doc *goquery.Document, t *volleynet.Tournament) error {
	tables := doc.Find("tbody")
	t.Teams = []*volleynet.TournamentTeam{}

	for i := range tables.Nodes {
		table := tables.Eq(i)
		rows := table.Find("tr")

		if rows.First().Children().Eq(0).Text() == "Nr." {
			team := &volleynet.TournamentTeam{}
			team.TournamentID = t.ID

			for j := range rows.Nodes {
				if j == 0 {
					continue
				}

				player, err := parsePlayerRow(rows.Eq(j), team)

				if err != nil {
					log.Warnf("error parsing player: %s", err)
					j++ // if it's not possible to parse a player, skip the entire team
					continue
				}

				player.Gender = t.Gender

				if team.Player1 == nil {
					team.Player1 = player
				} else {
					team.Player2 = player
					t.Teams = append(t.Teams, team)
					team = &volleynet.TournamentTeam{}
					team.TournamentID = t.ID

					if !team.Deregistered {
						t.SignedupTeams++
					}
				}
			}
		}
	}

	return nil
}

func parsePlayerRow(row *goquery.Selection, team *volleynet.TournamentTeam) (player *volleynet.Player, err error) {
	player = &volleynet.Player{}

	columnsCount := len(row.Children().Nodes)

	columns := row.Find("td")

	for k := range columns.Nodes {
		var err error
		column := columns.Eq(k)

		if columnsCount == 5 {
			switch k {
			case 0:
				team.Result = findInt(column.Text())
			case 1:
				player.ID, err = parsePlayerIDFromSteckbrief(column.Find("a"))
				player.FirstName, player.LastName = parsePlayerName(column)
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
				player.FirstName, player.LastName = parsePlayerName(column)
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
				player.FirstName, player.LastName = parsePlayerName(column)
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
				player.FirstName, player.LastName = parsePlayerName(column)
			case 1:
				player.CountryUnion = trimmSelectionText(column)
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
