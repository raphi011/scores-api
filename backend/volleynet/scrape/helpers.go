package scrape

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

func parseFloat(text string) float32 {
	val, err := strconv.ParseFloat(text, 32)
	if err != nil {
		return 0
	}

	return float32(val)
}

var numberRegex = regexp.MustCompile("\\d+")

func findInt(text string) (nr int, ok bool) {
	nrstr := numberRegex.FindString(text)

	nr, err := strconv.Atoi(nrstr)
	ok = err == nil

	return
}

func trimSelectionHTML(s *goquery.Selection) string {
	html, _ := s.Html()
	return strings.TrimSpace(strings.Replace(html, "<br/>", " ", -1))
}

func trimmSelectionText(s *goquery.Selection) string {
	return strings.TrimSpace(s.Text())
}

func parseLicenseNr(s string) (id int, year int) {
	parts := strings.Split(s, "/")

	if len(parts) != 2 {
		return -1, -1
	}

	id, _ = strconv.Atoi(parts[0])
	year, _ = strconv.Atoi(parts[1])

	return id, year
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

var lastNameRegex = regexp.MustCompile("\\p{Lu}+\\b")
var firstNameRegex = regexp.MustCompile("\\p{Lu}\\p{Ll}+\\b")

func parsePlayerName(c *goquery.Selection) (string, string) {
	playerName := c.Text()

	lastName := strings.Title(strings.ToLower(strings.Join(lastNameRegex.FindAllString(playerName, -1), " ")))
	firstName := strings.Title(strings.Join(firstNameRegex.FindAllString(playerName, -1), " "))

	return strings.Title(firstName), strings.Title(lastName)
}

func parsePlayerIDFromSteckbrief(s *goquery.Selection) (int, error) {
	href := parseHref(s)

	if href == "" {
		return 0, errors.New("could not parse playerid from empty steckbrief href")
	}

	dashIndex := strings.Index(href, "-")

	id, err := strconv.Atoi(href[dashIndex+1:])

	return id, errors.Wrap(err, "could not parse playerid from steckbrief link")
}

func parseStartEndDates(s *goquery.Selection) (time.Time, time.Time, error) {
	a := trimmSelectionText(s)

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

func parseDate(dateString string) (time.Time, error) {
	date, err := time.Parse("02.01.2006", strings.TrimSpace(dateString))
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to parse date '%s'", dateString)
	}

	return date, nil
}

func parseHref(anchor *goquery.Selection) string {
	href, _ := anchor.Attr("href")

	return strings.TrimSpace(href)
}

func parseHTML(html io.Reader) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(html)

	return doc, errors.Wrap(err, "invalid html")
}
