package volleynet

import (
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

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
	Phase            string    `json:"phase"`
	Link             string    `json:"link"`
	EntryLink        string    `json:"entryLink"`
	ID               int       `json:"id"`
	Status           string    `json:"status"` // done, upcoming, canceled
	RegistrationOpen bool      `json:"registrationOpen"`
	Gender           string    `json:"gender"`
}

func parseTournamentList(html io.Reader) ([]Tournament, error) {

	doc, err := parseHtml(html)

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
		tournament.Name = trimmTournamentName(column)

		column = columns.Eq(1)
		tournament.Start, tournament.End, err = parseStartEndDates(column)

		column = columns.Eq(4)
		content := trimmSelectionText(column)
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

func extractTournamentLinkData(link string) Tournament {
	id, _ := strconv.Atoi(readUrlPart(link, "cup/"))
	season, _ := strconv.Atoi(readUrlPart(link, "saison/"))

	return Tournament{
		Gender: readUrlPart(link, "sex/"),
		League: readUrlPart(link, "bewerbe/"),
		Phase:  readUrlPart(link, "phase/"),
		ID:     id,
		Season: season,
		Link:   link,
	}
}

func readUrlPart(link, start string) string {
	startIndex := strings.Index(link, start)

	if startIndex == -1 {
		return ""
	}

	link = link[startIndex+len(start):]

	endIndex := strings.Index(link, "/")

	if endIndex != -1 {
		link = link[:endIndex]
	}

	unescaped, err := url.PathUnescape(link)

	if err != nil {
		return link
	}

	return unescaped
}

func trimmTournamentName(s *goquery.Selection) string {
	name := trimmSelectionText(s)
	index := strings.Index(name, "- ")

	if index > 0 {
		return name[index+2:]
	}

	return name
}
