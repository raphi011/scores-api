package client

import (
	"fmt"
	"net/url"

	"github.com/raphi011/scores/volleynet"
)

func (c *Default) buildGetAPIURL(relativePath string, routeArgs ...interface{}) *url.URL {
	return buildGetAPIURL(c.GetURL, "/api", relativePath, routeArgs...)
}

func (c *Default) buildGetURL(relativePath string, routeArgs ...interface{}) *url.URL {
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

func (c *Default) buildPostURL(relativePath string, routeArgs ...interface{}) *url.URL {
	path := escapeArgs(relativePath, routeArgs...)

	link, err := url.Parse(c.PostURL + path)

	if err != nil {
		panic("cannot parse client GetUrl")
	}

	return link
}

// GetTournamentLink returns the link for a tournament.
func (c *Default) GetTournamentLink(t *volleynet.TournamentInfo) string {
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
func (c *Default) GetAPITournamentLink(t *volleynet.TournamentInfo) string {
	url := c.buildGetAPIURL("/beach/bewerbe/%s/phase/%s/sex/%s/saison/%d/cup/%d",
		t.League,
		t.League,
		t.Gender,
		t.Season,
		t.ID,
	)

	return url.String()
}
