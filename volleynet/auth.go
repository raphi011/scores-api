package volleynet

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

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
