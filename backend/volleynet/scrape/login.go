package scrape

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

type loginDataParser func(*goquery.Selection, *LoginData)

var parseLoginDataMap = map[string]loginDataParser{
	"Name": func(value *goquery.Selection, d *LoginData) {
		//  parsePlayerName(value) TODO
	},
	"Geburtsdatum": func(value *goquery.Selection, d *LoginData) {
		// d.Birthday = value.Text() TODO
	},
	"Lizenz": func(value *goquery.Selection, d *LoginData) {
		d.License.Type = value.Text()
	},
	"Lizenznummer": func(value *goquery.Selection, d *LoginData) {
		d.License.Nr = value.Text()

		d.ID, _ = parseLicenseNr(d.License.Nr)
	},
	"Beantragt": func(value *goquery.Selection, d *LoginData) {
		d.License.Requested = value.Text()
	},
}

// LoginData contains the data of the form that is shown after a successful login.
type LoginData struct {
	PlayerInfo
	License License
}

// License contains the license data of a player.
type License struct {
	Nr        string
	Type      string
	Requested string
}

// Login parses data from the successfull login page
func Login(html io.Reader) (*LoginData, error) {
	doc, err := parseHTML(html)

	if err != nil {
		return nil, err
	}

	loginData := &LoginData{}

	form := doc.Find("form[name='volleynet']")

	rows := form.Find("tr")

	for j := range rows.Nodes {
		row := rows.Eq(j).Children()
		columnName := row.Eq(1).Text()
		value := row.Eq(2)

		if parser, ok := parseLoginDataMap[columnName]; ok {
			parser(value, loginData)
		}
	}

	return loginData, nil
}
