package volleynet

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var numberRegex = regexp.MustCompile("\\d+")

func parseFloat(text string) float32 {
	val, err := strconv.ParseFloat(text, 32)
	if err != nil {
		return 0
	}

	return float32(val)
}

func parseNumber(text string) int {
	nrstr := numberRegex.FindString(text)
	nr, err := strconv.Atoi(nrstr)

	if err != nil {
		return -1
	}

	return nr
}

func trimmedText(s *goquery.Selection) string {
	return strings.TrimSpace(s.Text())
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

	return href
}
