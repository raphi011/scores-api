package client

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/raphi011/scores/volleynet"
	"github.com/raphi011/scores/volleynet/parse"
)

// Client is the interface to the volleynet api, use DefaultClient()
// to get a new Client.
type Client interface {
	GetTournamentLink(t *volleynet.Tournament) string
	GetAPITournamentLink(t *volleynet.Tournament) string
	Login(username, password string) (*volleynet.LoginData, error)
	AllTournaments(gender, league string, year int) ([]volleynet.Tournament, error)
	Ladder(gender string) ([]volleynet.Player, error)
	ComplementTournament(tournament volleynet.Tournament) (*volleynet.FullTournament, error)
	TournamentWithdrawal(tournamentID int) error
	TournamentEntry(playerName string, playerID, tournamentID int) error
	SearchPlayers(firstName, lastName, birthday string) ([]volleynet.PlayerInfo, error)
}

// ClientImpl implements the Client interface
type ClientImpl struct {
	PostURL string
	GetURL  string
	Cookie  string
}

// DefaultClient returns a Client with the correct PostURL and GetURL fields set.
func DefaultClient() Client {
	return &ClientImpl{
		PostURL: "https://beach.volleynet.at",
		GetURL:  "http://www.volleynet.at",
	}
}

func (c *ClientImpl) buildGetAPIURL(relativePath string, routeArgs ...interface{}) *url.URL {
	return buildGetAPIURL(c.GetURL, "/api", relativePath, routeArgs...)
}

func (c *ClientImpl) buildGetURL(relativePath string, routeArgs ...interface{}) *url.URL {
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

func (c *ClientImpl) buildPostURL(relativePath string, routeArgs ...interface{}) *url.URL {
	path := escapeArgs(relativePath, routeArgs...)

	link, err := url.Parse(c.PostURL + path)

	if err != nil {
		panic("cannot parse client GetUrl")
	}

	return link
}

// GetTournamentLink returns the link for a tournament.
func (c *ClientImpl) GetTournamentLink(t *volleynet.Tournament) string {
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
func (c *ClientImpl) GetAPITournamentLink(t *volleynet.Tournament) string {
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
func (c *ClientImpl) Login(username, password string) (*volleynet.LoginData, error) {
	form := url.Values{}
	form.Add("login_name", username)
	form.Add("login_pass", password)
	form.Add("action", "Beach/Profile/ProfileLogin")
	form.Add("submit", "OK")
	form.Add("mode", "X")

	url := c.buildPostURL("/Admin/formular").String()
	resp, err := http.PostForm(url, form)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unauthorized")
	}
	defer resp.Body.Close()

	loginData, err := parse.Login(resp.Body)

	c.Cookie = resp.Header.Get("Set-Cookie")

	semicolonIndex := strings.Index(c.Cookie, ";")

	if semicolonIndex > 0 {
		c.Cookie = c.Cookie[:semicolonIndex]
	}

	return loginData, nil
}

// AllTournaments reads all tournaments of a certain gender, league and year.
// To get more detailed tournamnent information call `ComplementTournament`.
func (c *ClientImpl) AllTournaments(gender, league string, year int) ([]volleynet.Tournament, error) {
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

	return parse.TournamentList(resp.Body, c.GetURL)
}

// Ladder reads all players of a certain gender.
func (c *ClientImpl) Ladder(gender string) ([]volleynet.Player, error) {
	url := c.buildGetAPIURL(
		"/beach/bewerbe/Rangliste/phase/%s",
		genderLong(gender),
	).String()

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return parse.Ladder(resp.Body)
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
func (c *ClientImpl) ComplementTournament(tournament volleynet.Tournament) (
	*volleynet.FullTournament, error) {
	url := c.GetAPITournamentLink(&tournament)

	fmt.Print(url)
	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.Wrapf(err, "loading tournament %d failed", tournament.ID)
	}

	t, err := parse.FullTournament(resp.Body, tournament)

	if err != nil {
		return nil, errors.Wrapf(err, "parsing tournament %d failed", tournament.ID)
	}

	t.Link = c.GetTournamentLink(&tournament)

	return t, nil
}

func (c *ClientImpl) loadUniqueWriteCode(tournamentID int) (string, error) {
	url := c.buildPostURL(
		"/Admin/index.php?screen=Beach/Profile/TurnierAnmeldung&parent=0&prev=0&next=0&cur=%d",
		tournamentID,
	).String()

	req, err := http.NewRequest(
		"GET",
		url,
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

	code, err := parse.UniqueWriteCode(resp.Body)

	return code, errors.Wrap(err, "parsing unique writecode failed")
}

// TournamentWithdrawal signs a player out of a tournament. A valid session Cookie must be set.
func (c *ClientImpl) TournamentWithdrawal(tournamentID int) error {
	url := c.buildPostURL("/Abmelden/0-%d-00-0", tournamentID).String()

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", c.Cookie)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return errors.Wrapf(err, "tournamentwithdrawal request for tournamentID: %d failed", tournamentID)
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("tournamentwithdrawal request for tournamentID: %d failed with code %d",
			tournamentID,
			resp.StatusCode)
	}

	return nil
}

// TournamentEntry signs a player up for a tournament. A valid session Cookie must be set.
func (c *ClientImpl) TournamentEntry(playerName string, playerID, tournamentID int) error {
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

	url := c.buildPostURL("/Admin/formular").String()

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return errors.Wrap(err, "creating tournamententry request failed")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", c.Cookie)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	defer resp.Body.Close()

	// TODO: volleynet always returns 200, parse content and see what really happened ..

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
func (c *ClientImpl) SearchPlayers(firstName, lastName, birthday string) ([]volleynet.PlayerInfo, error) {
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

	return parse.Players(response.Body)
}
