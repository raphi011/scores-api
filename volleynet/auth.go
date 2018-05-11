package volleynet

import (
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

	cookie := resp.Header.Get("Set-Cookie")

	c.Cookie = cookie[:strings.Index(cookie, ";")]

	return nil
}
