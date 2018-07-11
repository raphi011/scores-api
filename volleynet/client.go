package volleynet

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Client struct {
	PostURL string
	GetURL  string
	Cookie  string
}

func DefaultClient() *Client {
	return &Client{
		PostURL: "https://beach.volleynet.at/Admin",
		GetURL:  "http://www.volleynet.at/beach",
	}
}

func (c *Client) buildGetURL(relativePath string, routeArgs ...interface{}) *url.URL {
	escapedArgs := make([]interface{}, len(routeArgs))

	for i, val := range routeArgs {
		str, ok := val.(string)

		if ok {
			escapedArgs[i] = url.PathEscape(str)
		} else if stringer, ok := val.(fmt.Stringer); ok {
			escapedArgs[i] = stringer.String()
		} else {
			escapedArgs[i] = ""
		}
	}

	path := fmt.Sprintf(relativePath, escapedArgs...)
	link, err := url.Parse(c.GetURL + path)
	log.Print(link.String())

	if err != nil {
		panic("cannot parse client GetUrl")
	}

	return link
}

func (c *Client) buildPostURL(relativePath string) *url.URL {
	link, err := url.Parse(c.PostURL + relativePath)

	if err != nil {
		panic("cannot parse client GetUrl")
	}

	return link
}

func (c *Client) GetTournamentLink(t *Tournament) string {
	url := c.buildGetURL("/bewerbe/%s/phase/%s/sex/%s/saison/%s/cup/%s",
		t.League,
		t.Phase,
		t.Gender,
		t.Season,
		t.ID,
	)

	return url.String()
}

func (c *Client) Login(username, password string) error {
	form := url.Values{}
	form.Add("login_name", username)
	form.Add("login_pass", password)
	form.Add("action", "Beach/Profile/ProfileLogin")
	form.Add("submit", "OK")
	form.Add("mode", "X")

	url := c.buildPostURL("/formular")
	resp, err := http.PostForm(url.String(), form)

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
	url := c.buildGetURL(
		"/bewerbe/%s/phase/%s/sex/%s/saison/%s/information/all",
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

	return parseTournamentList(resp.Body)
}

func (c *Client) Ladder(gender string) ([]Player, error) {
	url := c.buildGetURL(
		"/beach/bewerbe/Rangliste/phase/%s",
		genderLong(gender),
	)

	resp, err := http.Get(url.String())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return parseLadder(resp.Body)
}

func genderLong(gender string) string {
	if gender == "M" {
		return "Herren"
	} else if gender == "W" {
		return "Damen"
	}

	return ""
}

func (c *Client) ComplementTournament(tournament Tournament) (
	*FullTournament, error) {
	url := c.GetTournamentLink(&tournament)

	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.Wrapf(err, "loading tournament %d failed", tournament.ID)
	}

	t, err := parseFullTournament(resp.Body, tournament)

	if err != nil {
		return nil, errors.Wrapf(err, "parsing tournament %d failed", tournament.ID)
	}

	return t, nil
}

func (c *Client) loadUniqueWriteCode(tournamentID int) (string, error) {
	url := c.buildGetURL(
		"/index.php?screen=Beach/Profile/TurnierAnmeldung&screen=Beach%2FProfile%2FTurnierAnmeldung&parent=0&prev=0&next=0&cur=%d",
		tournamentID,
	)

	req, err := http.NewRequest(
		"GET",
		url.String(),
		nil)

	req.Header.Add("Cookie", c.Cookie)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return "", errors.Wrap(err, "loading unique writecode failed")
	}

	code, err := parseUniqueWriteCode(resp.Body)

	return code, errors.Wrap(err, "parsing unique writecode failed")
}

func (c *Client) TournamentEntry(playerName string, playerID, tournamentID int) error {
	form := url.Values{}

	code, err := c.loadUniqueWriteCode(tournamentID)

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

	url := c.buildPostURL("/formular")

	req, err := http.NewRequest("POST", url.String(), bytes.NewBufferString(form.Encode()))
	if err != nil {
		return errors.Wrap(err, "creating tournamententry request failed")
	}
	req.Header.Set("Content-Type", "application/x-www-form-=rlencoded")
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

	url := c.buildPostURL("/formular")

	response, err := http.PostForm(url.String(), form)

	if err != nil {
		return nil, err
	}

	return parsePlayers(response.Body)
}
