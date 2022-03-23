package client

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/raphi011/scores-api/volleynet"
	"github.com/raphi011/scores-api/volleynet/scrape"
)

// Client is the interface to the volleynet api, use DefaultClient()
// to get a new Client.
type Client interface {
	Login(username, password string) (*scrape.LoginData, error)

	Tournaments(gender, league string, year int) ([]*volleynet.TournamentInfo, error)
	Ladder(gender string) ([]*volleynet.Player, error)
	ComplementTournament(tournament *volleynet.TournamentInfo) (*volleynet.Tournament, error)

	WithdrawFromTournament(tournamentID int) error
	EnterTournament( /*playerName string, */ playerID, tournamentID int) error

	SearchPlayers(firstName, lastName, birthday string) ([]*scrape.PlayerInfo, error)
}

// defaultClient implements the Client interface
type defaultClient struct {
	PostURL string
	GetURL  string
	Cookie  string
}

// Default returns a Client with the correct PostURL and GetURL fields set.
func Default() Client {
	return &defaultClient{
		PostURL: "https://beach.volleynet.at",
		GetURL:  "http://www.volleynet.at",
	}
}

// Login authenticates the user against the volleynet page, if
// successfull the Client cookie is set, else an error is returned.
func (c *defaultClient) Login(username, password string) (*scrape.LoginData, error) {
	form := url.Values{}
	form.Add("login_name", username)
	form.Add("login_pass", password)
	form.Add("action", "Beach/Profile/ProfileLogin")
	form.Add("submit", "OK")
	form.Add("mode", "X")

	url := c.buildPostURL("/Admin/formular").String()
	resp, err := http.PostForm(url, form)

	if err != nil {
		return nil, fmt.Errorf("client login: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("login status: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	loginData, err := scrape.Login(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("parse login: %w", err)
	}

	c.Cookie = resp.Header.Get("Set-Cookie")

	semicolonIndex := strings.Index(c.Cookie, ";")

	if semicolonIndex > 0 {
		c.Cookie = c.Cookie[:semicolonIndex]
	}

	return loginData, nil
}

// Tournaments reads all tournaments of a certain gender, league and year.
// To get all details of a tournamnent use `Client.ComplementTournament`.
func (c *defaultClient) Tournaments(gender, league string, year int) ([]*volleynet.TournamentInfo, error) {
	url := c.buildGetAPIURL(
		"/beach/bewerbe/%s/phase/%s/sex/%s/saison/%d/information/all",
		league,
		league,
		gender,
		year,
	)

	resp, err := http.Get(url.String())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return scrape.TournamentList(resp.Body, c.GetURL)
}

// Ladder loads all ranked players of a certain gender.
func (c *defaultClient) Ladder(gender string) ([]*volleynet.Player, error) {
	url := c.buildGetAPIURL(
		"/beach/bewerbe/Rangliste/phase/%s",
		genderLong(gender),
	).String()

	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("loading ladder %q failed %w", gender, err)
	}

	defer resp.Body.Close()

	return scrape.Ladder(resp.Body)
}

func genderLong(gender string) string {
	if gender == "M" {
		return "Herren"
	} else if gender == "W" {
		return "Damen"
	}

	return ""
}

// ComplementTournament adds the missing information from `Tournaments`.
func (c *defaultClient) ComplementTournament(tournament *volleynet.TournamentInfo) (
	*volleynet.Tournament, error) {
	url := c.getAPITournamentLink(tournament)

	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("loading tournament %d failed %w", tournament.ID, err)
	}

	t, err := scrape.Tournament(resp.Body, time.Now(), tournament)

	if err != nil {
		return nil, fmt.Errorf("parsing tournament %d failed %w", tournament.ID, err)
	}

	t.Link = c.getTournamentLink(tournament)

	return t, nil
}

func (c *defaultClient) loadUniqueWriteCode(tournamentID int) (string, error) {
	url := c.buildPostURL(
		"/Admin/index.php?screen=Beach/Profile/TurnierAnmeldung&parent=0&prev=0&next=0&cur=%d",
		tournamentID,
	).String()

	req, err := http.NewRequest(
		"GET",
		url,
		nil)

	if err != nil {
		return "", fmt.Errorf("creating request failed: %w", err)
	}

	req.Header.Add("Cookie", c.Cookie)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return "", fmt.Errorf("loading unique writecode failed: %w", err)
	}

	code, err := scrape.UniqueWriteCode(resp.Body)

	return code, fmt.Errorf("parsing unique writecode failed: %w", err)
}

// WithdrawFromTournament withdraws a player from a tournament.
// A valid session Cookie must be set.
func (c *defaultClient) WithdrawFromTournament(tournamentID int) error {
	url := c.buildPostURL("/Abmelden/0-%d-00-0", tournamentID).String()

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", c.Cookie)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("tournamentwithdrawal request for tournamentID: %d failed %w", tournamentID, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("tournamentwithdrawal request for tournamentID: %d failed with code %d",
			tournamentID,
			resp.StatusCode)
	}

	return nil
}

// EnterTournament enters a player at a tournament.
// A valid session Cookie must be set.
func (c *defaultClient) EnterTournament( /*playerName string, */ playerID, tournamentID int) error {
	if c.Cookie == "" {
		return errors.New("cookie must be set")
	}

	form := url.Values{}

	code, err := c.loadUniqueWriteCode(tournamentID)

	if err != nil {
		return fmt.Errorf("could not load writecode for tournamentID: %d %w", tournamentID, err)
	}

	form.Add("action", "Beach/Profile/TurnierAnmeldung")
	form.Add("XX_unique_write_XXBeach/Profile/TurnierAnmeldung", code)
	form.Add("parent", "0")
	form.Add("prev", "0")
	form.Add("next", "0")
	form.Add("cur", strconv.Itoa(tournamentID))
	// form.Add("name_b", playerName) TODO: is this really needed?
	form.Add("bte_per_id_b", strconv.Itoa(playerID))
	form.Add("submit", "Anmelden")

	url := c.buildPostURL("/Admin/formular").String()

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(form.Encode()))

	if err != nil {
		return fmt.Errorf("creating tournamententry request failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", c.Cookie)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("tournamententry request for tournamentID: %d failed %w", tournamentID, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("tournamententry request for tournamentID: %d failed with code %d",
			tournamentID,
			resp.StatusCode)
	}

	entryData, err := scrape.Entry(resp.Body)

	if err != nil || !entryData.Successfull {
		return fmt.Errorf("tournamententry request for tournamentID: %d failed %w", tournamentID, err)
	}

	return nil
}

func traceResponse(resp *http.Response) {
	outFile, err := os.Create("/home/raphi/login-response.html")
	if err != nil {
		return
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return
	}
}

// SearchPlayers searches for players via firstName, lastName and their birthdate in dd.mm.yyyy format.
func (c *defaultClient) SearchPlayers(firstName, lastName, birthday string) ([]*scrape.PlayerInfo, error) {
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

	url := c.buildPostURL("/Admin/formular")

	response, err := http.PostForm(url.String(), form)

	if err != nil {
		return nil, err
	}

	return scrape.Players(response.Body)
}
