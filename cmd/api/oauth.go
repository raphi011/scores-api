package main

import (
	"encoding/json"
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type credentials struct {
	ClientID    string `json:"client_id"`
	CientSecret string `json:"client_secret"`
}

func googleOAuthConfig(configPath, host string) (*oauth2.Config, error) {
	var credentials credentials
	file, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	redirectURL := host + "/api/auth"

	json.Unmarshal(file, &credentials)
	config := &oauth2.Config{
		ClientID:     credentials.ClientID,
		ClientSecret: credentials.CientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	return config, nil
}
