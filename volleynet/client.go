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
	Cookie      string
}

func DefaultClient() *Client {
	return &Client{
		PostUrl:     "https://beach.volleynet.at/Admin/formular",
		ApiUrl:      "http://www.volleynet.at/api/",
		DefaultUrl:  "www.volleynet.at",
		AmateurPath: "beach/bewerbe/%s/phase/%s/sex/%s/saison/%s/information/all",
	}
}

func GetDocument(html io.Reader) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(html)
}
