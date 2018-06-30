package volleynet

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"

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
	} else if gender == "W" {
		return "Damen"
	}

	return ""
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

var registerUrl string = "https://beach.volleynet.at/Admin/index.php?screen=Beach/Profile/TurnierAnmeldung&screen=Beach%2FProfile%2FTurnierAnmeldung&parent=0&prev=0&next=0&cur="

func (c *Client) LoadUniqueWriteCode(tournamentID int) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%d", registerUrl, tournamentID), nil)
	req.Header.Add("Cookie", c.Cookie)
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return "", errors.Wrap(err, "loading unique writecode failed")
	}

	code, err := parseUniqueWriteCode(resp.Body)

	return code, errors.Wrap(err, "parsing unique writecode failed")
}

func (c *Client) GetTournamentLink(t *Tournament) string {
	return c.DefaultUrl + t.Link
}

func (c *Client) GetApiTournamentLink(link string) string {
	return c.ApiUrl + link
}

func ParseHtml(html io.Reader) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(html)

	return doc, errors.Wrap(err, "invalid html")
}

func (c *Client) Login(username, password string) error {
	form := url.Values{}
	form.Add("login_name", username)
	form.Add("login_pass", password)

	form.Add("action", "Beach/Profile/ProfileLogin")
	form.Add("submit", "OK")
	form.Add("mode", "X")

	resp, err := http.PostForm(c.PostUrl, form)

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("unauthorized")
	}

	c.Cookie = resp.Header.Get("Set-Cookie")

	semicolonIndex := strings.Index(c.Cookie, ";")

	if semicolonIndex > 0 {
		c.Cookie = c.Cookie[:semicolonIndex]
	}

	return nil
}

func (c *Client) AllTournaments(gender, league, year string) ([]Tournament, error) {
	url, err := url.Parse(c.ApiUrl)

	if err != nil {
		return nil, err
	}

	url.Path += fmt.Sprintf(c.AmateurPath, league, league, gender, year)

	encodedURL := url.String()
	resp, err := http.Get(encodedURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return parseTournaments(resp.Body)
}

func (c *Client) Ladder(gender string) ([]Player, error) {
	url, err := url.Parse(c.ApiUrl)

	// gender: Herren, Damen
	url.Path += fmt.Sprintf(c.LadderPath, genderLong(gender))

	encodedURL := url.String()
	resp, err := http.Get(encodedURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return parseLadder(resp.Body)
}

func (c *Client) ComplementTournament(tournament Tournament) (
	*FullTournament, error) {
	resp, err := http.Get(c.GetApiTournamentLink(tournament.Link))

	if err != nil {
		return nil, errors.Wrap(err, "loading tournament failed")
	}

	t, err := parseFullTournament(resp.Body, tournament)

	if err != nil {
		return nil, errors.Wrap(err, "parsing tournament failed")
	}

	return t, nil
}

func (c *Client) TournamentEntry(playerName string, playerID, tournamentID int) error {
	form := url.Values{}

	code, err := c.LoadUniqueWriteCode(tournamentID)

	if err != nil {
		return errors.Wrapf(err, "loading unique writecode failed for tournamentID: %d", tournamentID)
	}

	form.Add("action", "Beach/Profile/TurnierAnmeldung")
	form.Add("XX_unique_write_XXBeach/Profile/TurnierAnmeldung", code)
	form.Add("parent", "0")
	form.Add("prev", "0")
	form.Add("next", "0")
	form.Add("cur", strconv.Itoa(tournamentID))
	form.Add("name_b", playerName)
	form.Add("bte_per_id_b", strconv.Itoa(playerID))
	form.Add("submit", "Anmelden")

	req, err := http.NewRequest("POST", c.PostUrl, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return errors.Wrap(err, "creating tournamententry request failed")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", c.Cookie)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return errors.Wrapf(err, "tournamententry request for tournamentID: %d failed", tournamentID)
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("tournamententry request for tournamentID: %d failed with code %d",
			tournamentID,
			resp.StatusCode)
	}

	return nil
}

func (c *Client) SearchPlayers(firstName, lastName, birthday string) ([]PlayerInfo, error) {
	form := url.Values{}

	form.Add("XX_unique_write_XXAdmin/Search", "0.50981600 1525795371")
	form.Add("popup", "1")
	form.Add("add", "")
	form.Add("target", "bte_per_id_b")
	form.Add("txm_language", "de")
	form.Add("sai_id", "")
	form.Add("action", "Admin/Search")
	form.Add("submit", "Suchen")
	form.Add("search", "Person")
	form.Add("per_name", lastName)
	form.Add("per_vorname", firstName)
	form.Add("per_geburtsdatum", birthday)
	form.Add("doit", "1")
	form.Add("text", "0")

	response, err := http.PostForm(c.PostUrl, form)

	if err != nil {
		return nil, err
	}

	return parsePlayers(response.Body)
}
