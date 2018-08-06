package volleynet

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Client is the interface to the volleynet api, use DefaultClient()
// to get a new Client.
type Client struct {
	PostURL string
	GetURL  string
	Cookie  string
}

// DefaultClient returns a Client with the correct PostURL and GetURL fields set.
func DefaultClient() *Client {
	return &Client{
		PostURL: "https://beach.volleynet.at",
		GetURL:  "http://www.volleynet.at",
	}
}

func (c *Client) buildGetAPIURL(relativePath string, routeArgs ...interface{}) *url.URL {
	return buildGetAPIURL(c.GetURL, "/api", relativePath, routeArgs...)
}

func (c *Client) buildGetURL(relativePath string, routeArgs ...interface{}) *url.URL {
	return buildGetAPIURL(c.GetURL, "", relativePath, routeArgs...)
}

func escapeArgs(path string, routeArgs ...interface{}) string {
	escapedArgs := make([]interface{}, len(routeArgs))

	for i, val := range routeArgs {
		str, ok := val.(string)

		if ok {
			escapedArgs[i] = url.PathEscape(str)
		} else if stringer, ok := val.(fmt.Stringer); ok {
			escapedArgs[i] = stringer.String()
		} else {
			escapedArgs[i] = val
		}
	}

	return fmt.Sprintf(path, escapedArgs...)
}

func buildGetAPIURL(host, prefixedPath, relativePath string, routeArgs ...interface{}) *url.URL {
	path := escapeArgs(relativePath, routeArgs...)

	link, err := url.Parse(host + prefixedPath + path)

	if err != nil {
		panic("cannot parse client GetUrl")
	}

	return link
}

func (c *Client) buildPostURL(relativePath string, routeArgs ...interface{}) *url.URL {
	path := escapeArgs(relativePath, routeArgs...)

	link, err := url.Parse(c.PostURL + path)

	if err != nil {
		panic("cannot parse client GetUrl")
	}

	return link
}

// GetTournamentLink returns the link for a tournament.
func (c *Client) GetTournamentLink(t *Tournament) string {
	url := c.buildGetURL("/beach/bewerbe/%s/phase/%s/sex/%s/saison/%d/cup/%d",
		t.League,
		t.League,
		t.Gender,
		t.Season,
		t.ID,
	)

	return url.String()
}

// GetAPITournamentLink returns the API link for a tournament.
func (c *Client) GetAPITournamentLink(t *Tournament) string {
	url := c.buildGetAPIURL("/beach/bewerbe/%s/phase/%s/sex/%s/saison/%d/cup/%d",
		t.League,
		t.League,
		t.Gender,
		t.Season,
		t.ID,
	)

	return url.String()
}

// Login authenticates the user against the volleynet page, if
// successfull the Client cookie is set, else an error is returned.
func (c *Client) Login(username, password string) error {
	form := url.Values{}
	form.Add("login_name", username)
	form.Add("login_pass", password)
	form.Add("action", "Beach/Profile/ProfileLogin")
	form.Add("submit", "OK")
	form.Add("mode", "X")

	url := c.buildPostURL("/Admin/formular")
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

// AllTournaments reads all tournaments of a certain gender, league and year.
// To get more detailed tournamnent information call `ComplementTournament`.
func (c *Client) AllTournaments(gender, league, year string) ([]Tournament, error) {
	url := c.buildGetAPIURL(
		"/beach/bewerbe/%s/phase/%s/sex/%s/saison/%s/information/all",
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

	return parseTournamentList(resp.Body, c.GetURL)
}

// Ladder reads all players of a certain gender.
func (c *Client) Ladder(gender string) ([]Player, error) {
	url := c.buildGetAPIURL(
		"/beach/bewerbe/Rangliste/phase/%s",
		genderLong(gender),
	).String()

	resp, err := http.Get(url)

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

// ComplementTournament adds the missing information from `AllTournaments`.
func (c *Client) ComplementTournament(tournament Tournament) (
	*FullTournament, error) {
	url := c.GetAPITournamentLink(&tournament)

	fmt.Print(url)
	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.Wrapf(err, "loading tournament %d failed", tournament.ID)
	}

	t, err := parseFullTournament(resp.Body, tournament)

	if err != nil {
		return nil, errors.Wrapf(err, "parsing tournament %d failed", tournament.ID)
	}

	t.Link = c.GetTournamentLink(&tournament)

	return t, nil
}

func (c *Client) loadUniqueWriteCode(tournamentID int) (string, error) {
	url := c.buildPostURL(
		"/Anmelden/0-%d-00",
		tournamentID,
	)

	req, err := http.NewRequest(
		"GET",
		url.String(),
		nil)

	if err != nil {
		return "", errors.Wrap(err, "creating request failed")
	}

	req.Header.Add("Cookie", c.Cookie)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return "", errors.Wrap(err, "loading unique writecode failed")
	}

	code, err := parseUniqueWriteCode(resp.Body)

	return code, errors.Wrap(err, "parsing unique writecode failed")
}

// TournamentEntry signs a player up for a tournament. A valid session Cookie must be set.
func (c *Client) TournamentEntry(playerName string, playerID, tournamentID int) error {
	if c.Cookie == "" {
		return errors.New("cookie must be set")
	}

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

	url := c.buildPostURL("/Admin/formular")

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

// SearchPlayers searches for players via firstName, lastName and their birthdate in dd.mm.yyyy format.
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

	url := c.buildPostURL("/Admin/formular")

	response, err := http.PostForm(url.String(), form)

	if err != nil {
		return nil, err
	}

	return parsePlayers(response.Body)
}
