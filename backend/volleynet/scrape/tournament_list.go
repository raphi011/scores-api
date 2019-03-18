package scrape

import (
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/raphi011/scores"
	"github.com/raphi011/scores/volleynet"
)

// TournamentList parses the list of tournaments
func TournamentList(html io.Reader, host string) ([]*volleynet.TournamentInfo, error) {
	doc, err := parseHTML(html)

	if err != nil {
		return nil, err
	}

	tournaments := []*volleynet.TournamentInfo{}

	rows := doc.Find("tbody>tr")

	for i := range rows.Nodes {
		r := rows.Eq(i)

		columns := r.Find("td")

		if len(columns.Nodes) != 5 {
			continue
		}

		column := columns.Eq(2)

		tournament := extractTournamentLinkData(parseHref(column.Find("a")), host)
		tournament.Name = trimmTournamentName(column)

		column = columns.Eq(1)
		tournament.Start, tournament.End, err = parseStartEndDates(column)

		column = columns.Eq(4)
		content := trimmSelectionText(column)
		if content == "Abgesagt" {
			tournament.Status = volleynet.StatusCanceled
			tournament.RegistrationOpen = false
		} else {
			// `StatusClosed` is set exclusively in parseFullTournament
			tournament.Status = volleynet.StatusUpcoming

			if entryLink := column.Find("a"); entryLink.Length() == 1 {
				tournament.RegistrationOpen = true
				tournament.EntryLink = parseHref(entryLink)
			} else {
				tournament.RegistrationOpen = false
			}
		}

		tournaments = append(tournaments, tournament)
	}

	return tournaments, nil
}

func extractTournamentLinkData(relativeLink, host string) *volleynet.TournamentInfo {
	if len(relativeLink) == 0 {
		return nil
	}

	if relativeLink[0] == '/' {
		relativeLink = relativeLink[1:]
	}

	id, _ := strconv.Atoi(readURLPart(relativeLink, "cup/"))

	season := readURLPart(relativeLink, "saison/")

	return &volleynet.TournamentInfo{
		Gender:       readURLPart(relativeLink, "sex/"),
		League:       readURLPart(relativeLink, "bewerbe/"),
		LeagueKey:    scores.Sluggify(readURLPart(relativeLink, "bewerbe/")),
		SubLeague:    readURLPart(relativeLink, "phase/"),
		SubLeagueKey: scores.Sluggify(readURLPart(relativeLink, "phase/")),
		ID:           id,
		Season:       season,
		Link:         host + "/" + relativeLink,
	}
}

func readURLPart(link, start string) string {
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
