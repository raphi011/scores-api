package scrape

import (
	"fmt"
	"io"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/raphi011/scores-api"
	"github.com/raphi011/scores-api/volleynet"
)

// Tournament adds remaining details to the tournament (parsed by TournamentList()).
func Tournament(
	html io.Reader,
	now time.Time,
	tournament *volleynet.TournamentInfo) (*volleynet.Tournament, error) {

	doc, err := parseHTML(html)

	if err != nil {
		return nil, fmt.Errorf("parse tournament: %w", err)
	}

	t := &volleynet.Tournament{TournamentInfo: *tournament}

	parseTournamentNotes(doc, t)

	if err = parseTournamentDetails(doc, t); err != nil {
		return nil, fmt.Errorf("parse tournament details: %w", err)
	}

	if err = parseFullTournamentTeams(doc, t); err != nil {
		return nil, fmt.Errorf("parse tournament teams: %w", err)
	}

	if isDone(t, now) {
		t.Status = volleynet.StatusDone
	}

	return t, nil
}

// isDone returns true if the results for the tournament are in
// or the tournament has not taken place.
func isDone(t *volleynet.Tournament, now time.Time) bool {
	if t.Status != volleynet.StatusUpcoming {
		// nothing to do here
		return false
	}

	if len(t.Teams) > 0 && t.Teams[0].Result > 0 {
		return true
	}

	return hasNotTakenPlace(t, now)
}

// hasNotTakenPlace returns true if no results were added a week after the
// tournament has ended - thus assuming that the tournament has not taken place.
func hasNotTakenPlace(t *volleynet.Tournament, now time.Time) bool {
	updateDeadline := t.End.AddDate(0, 0, 7)

	return isDateAfter(now, updateDeadline)
}

func getDate(d time.Time) time.Time {
	date := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
	return date
}

func isDateAfter(d1, d2 time.Time) bool {
	return getDate(d1).After(getDate(d2))
}

func parseTournamentNotes(doc *goquery.Document, t *volleynet.Tournament) {
	htmlNotes := doc.Find(".extrainfo")

	if htmlNotes.Find("iframe").Length() > 0 {
		t.HTMLNotes = "Cannot display these notes yet."
	} else {
		t.HTMLNotes, _ = htmlNotes.Html()
	}
}

type detailsParser func(*goquery.Selection, *volleynet.Tournament) error

var parseTournamentDetailsMap = map[string]detailsParser{
	"Kategorie": func(value *goquery.Selection, t *volleynet.Tournament) error {
		t.SubLeague = trimmSelectionText(value)
		t.SubLeagueKey = scores.Sluggify(trimmSelectionText(value))

		return nil
	},
	"Modus": func(value *goquery.Selection, t *volleynet.Tournament) error {
		t.Mode = trimmSelectionText(value)
		t.MaxTeams, _ = findInt(t.Mode)

		return nil
	},
	"Teiln. Qual.": func(value *goquery.Selection, t *volleynet.Tournament) error {
		// TODO: not min teams but min teams for qualification, this is misleading
		t.MinTeams, _ = findInt(value.Text())

		return nil
	},
	"Datum": func(value *goquery.Selection, t *volleynet.Tournament) error {
		var err error
		t.Start, t.End, err = parseStartEndDates(value)

		return fmt.Errorf("parsing start/end dates from tournamentId: %d, value: %s %w", t.ID, value.Text(), err)
	},
	"Ort": func(value *goquery.Selection, t *volleynet.Tournament) error {
		t.Location = trimSelectionHTML(value)

		return nil
	},
	"Max. Punkte": func(value *goquery.Selection, t *volleynet.Tournament) error {
		t.MaxPoints, _ = findInt(value.Text())

		return nil
	},
	"Veranstalter": func(value *goquery.Selection, t *volleynet.Tournament) error {
		t.Organiser = trimmSelectionText(value)

		return nil
	},
	"Telefon": func(value *goquery.Selection, t *volleynet.Tournament) error {
		t.Phone = trimmSelectionText(value)

		return nil
	},
	"EMail": func(value *goquery.Selection, t *volleynet.Tournament) error {
		t.Email = trimmSelectionText(value)

		return nil
	},
	"Web": func(value *goquery.Selection, t *volleynet.Tournament) error {
		t.Website = trimmSelectionText(value)

		return nil
	},
	"Vorläufige Punkte": func(value *goquery.Selection, t *volleynet.Tournament) error {
		t.CurrentPoints = trimmSelectionText(value)

		return nil
	},
	"Nennschluss": func(value *goquery.Selection, t *volleynet.Tournament) error {
		endRegistration, _ := parseDate(value.Text())
		if endRegistration.IsZero() {
			t.EndRegistration = nil
		} else {
			t.EndRegistration = &endRegistration
		}

		return nil
	},
}

func parseTournamentDetails(doc *goquery.Document, t *volleynet.Tournament) error {
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
					if err := parser(value, t); err != nil {
						return fmt.Errorf("error parsing column %s with value %+v %w", value.Text(), t, err)
					}
				}
			}
		}
	}

	return nil
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
					j++ // if it's not possible to parse a player, skip the entire team
					continue
				}

				player.Gender = t.Gender

				if team.Player1 == nil {
					team.Player1 = player
				} else {
					team.Player2 = player

					if !team.Deregistered {
						t.SignedupTeams++
					}

					t.Teams = append(t.Teams, team)

					team = &volleynet.TournamentTeam{}
					team.TournamentID = t.ID

				}
			}
		}
	}

	return nil
}

// EntryResult contains the data that is returned by the
// TournamentEntry endpoint.
type EntryResult struct {
	Successfull bool `json:"successfull"`
}

// Entry parses the result of a tournament entry.
func Entry(body io.Reader) (EntryResult, error) {
	doc, err := parseHTML(body)
	result := EntryResult{}

	if err != nil {
		return result, fmt.Errorf("could not parse html: %w", err)
	}

	selection := doc.Find("[name='XX_unique_write_XXBeach/Profile/TurnierAnmeldungErfolgreich']")

	if selection.Length() > 0 {
		result.Successfull = true
	}

	return result, nil
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
				team.Result, _ = findInt(column.Text())
			case 1:
				player.ID, err = parsePlayerIDFromSteckbrief(column.Find("a"))
				player.FirstName, player.LastName = parsePlayerName(column)
			case 2:
				player.CountryUnion = trimmSelectionText(column)
			case 3:
				team.WonPoints, _ = findInt(column.Text())
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
				player.TotalPoints, _ = findInt(column.Text())
			}
		} else if columnsCount == 7 {
			switch k {
			case 0:
				team.Seed, _ = findInt(column.Text())
			case 1:
				player.ID, err = parsePlayerIDFromSteckbrief(column.Find("a"))
				player.FirstName, player.LastName = parsePlayerName(column)
			case 2:
				player.License = trimmSelectionText(column)
			case 3:
				player.CountryUnion = trimmSelectionText(column)
			case 4:
				player.TotalPoints, _ = findInt(column.Text())
			case 5:
				team.TotalPoints, _ = findInt(column.Text())
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
