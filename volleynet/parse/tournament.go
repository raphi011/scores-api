package parse

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/raphi011/scores/volleynet"
	log "github.com/sirupsen/logrus"
)

func FullTournament(
	html io.Reader,
	tournament volleynet.Tournament) (*volleynet.FullTournament, error) {

	doc, err := parseHTML(html)

	if err != nil {
		return nil, errors.Wrap(err, "parseFullTournament failed")
	}

	t := &volleynet.FullTournament{Tournament: tournament}

	parseTournamentNotes(doc, t)
	parseTournamentDetails(doc, t)

	err = parseFullTournamentTeams(doc, t)

	if err != nil {
		return nil, errors.Wrap(err, "error parsing tournament teams")
	}

	if len(t.Teams) > 0 && t.Teams[0].Rank > 0 {
		t.Status = volleynet.StatusDone
	}

	return t, nil
}

func parseTournamentNotes(doc *goquery.Document, t *volleynet.FullTournament) {
	htmlNotes := doc.Find(".extrainfo")

	if htmlNotes.Find("iframe").Length() > 0 {
		t.HTMLNotes = "Cannot display these notes yet."
	} else {
		t.HTMLNotes, _ = htmlNotes.Html()
	}
}

type detailsParser func(*goquery.Selection, *volleynet.FullTournament)

var parseTournamentDetailsMap = map[string]detailsParser{
	"Kategorie": func(value *goquery.Selection, t *volleynet.FullTournament) {
		t.Phase = trimmSelectionText(value)
	},
	"Modus": func(value *goquery.Selection, t *volleynet.FullTournament) {
		t.Mode = trimmSelectionText(value)
		t.MaxTeams = findInt(t.Mode)
	},
	"Teiln. Qual.": func(value *goquery.Selection, t *volleynet.FullTournament) {
		t.MinTeams = findInt(value.Text())
	},
	"Datum": func(value *goquery.Selection, t *volleynet.FullTournament) {
		var err error
		t.Start, t.End, err = parseStartEndDates(value)

		if err != nil {
			log.Warn("error parsing start/end dates from tournamentId: %d, value: %s", t.ID, value.Text())
		}
	},
	"Ort": func(value *goquery.Selection, t *volleynet.FullTournament) {
		t.Location = trimmSelectionHtml(value)
	},
	"Max. Punkte": func(value *goquery.Selection, t *volleynet.FullTournament) {
		t.MaxPoints = findInt(value.Text())
	},
	"Veranstalter": func(value *goquery.Selection, t *volleynet.FullTournament) {
		t.Organiser = trimmSelectionText(value)
	},
	"Telefon": func(value *goquery.Selection, t *volleynet.FullTournament) {
		t.Phone = trimmSelectionText(value)
	},
	"EMail": func(value *goquery.Selection, t *volleynet.FullTournament) {
		t.Email = trimmSelectionText(value)
	},
	"Web": func(value *goquery.Selection, t *volleynet.FullTournament) {
		t.Web = trimmSelectionText(value)
	},
	"Vorläufige Punkte": func(value *goquery.Selection, t *volleynet.FullTournament) {
		t.CurrentPoints = trimmSelectionText(value)
	},
	"Nennschluss": func(value *goquery.Selection, t *volleynet.FullTournament) {
		t.EndRegistration, _ = parseDate(value.Text())
	},
}

func parseTournamentDetails(doc *goquery.Document, t *volleynet.FullTournament) {
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

func parseFullTournamentTeams(doc *goquery.Document, t *volleynet.FullTournament) error {
	tables := doc.Find("tbody")

	for i := range tables.Nodes {
		table := tables.Eq(i)
		rows := table.Find("tr")

		if rows.First().Children().Eq(0).Text() == "Nr." {
			team := volleynet.TournamentTeam{}
			team.TournamentID = t.ID

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

				player.Gender = t.Gender

				if team.Player1 == nil {
					team.Player1 = player
				} else {
					team.Player2 = player
					t.Teams = append(t.Teams, team)
					team = volleynet.TournamentTeam{}
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
				team.Rank = findInt(column.Text())
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
