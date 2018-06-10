package volleynet

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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
	} else if gender == "D" {
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

func (c *Client) GetTournamentLink(t *Tournament) string {
	return c.DefaultUrl + t.Link
}

func (c *Client) GetApiTournamentLink(t *Tournament) string {
	return c.ApiUrl + t.Link
}

func GetDocument(html io.Reader) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(html)
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

	return c.parseTournaments(resp.Body)
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
	resp, err := http.Get(tournament.ApiLink)

	if err != nil {
		return nil, err
	}

	t, err := parseFullTournament(resp.Body, tournament)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) TournamentEntry(playerName string, playerID, tournamentID int) error {
	form := url.Values{}

	code, err := c.GetUniqueWriteCode(tournamentID)

	if err != nil {
		return err
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
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", c.Cookie)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return errors.New("entry did not work")
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
