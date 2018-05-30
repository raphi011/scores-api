package volleynet

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

type Client struct {
	PostUrl     string
	DefaultUrl  string
	ApiUrl      string
	AmateurPath string
	LadderPath  string
	Cookie      string
}

func genderLong(gender string) string {
	if gender == "M" {
		return "Herren"
	} else {
		return "Damen"
	}
}

func DefaultClient() *Client {
	return &Client{
		PostUrl:     "https://beach.volleynet.at/Admin/formular",
		ApiUrl:      "http://www.volleynet.at/api/",
		DefaultUrl:  "www.volleynet.at",
		AmateurPath: "beach/bewerbe/%s/phase/%s/sex/%s/saison/%s/information/all",
		LadderPath:  "beach/bewerbe/Rangliste/phase/%s",
	}
}

func GetDocument(html io.Reader) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(html)
}
