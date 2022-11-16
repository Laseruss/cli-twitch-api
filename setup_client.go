package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
)

type Client struct {
	client      http.Client
	credentials Credentials
}

type Credentials struct {
	userId       string
	clientId     string
	clientSecret string
	accessToken  string
}

func newClient() (*Client, error) {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		return nil, err
	}

	return &Client{
		client:      http.Client{},
		credentials: setCredentials(myEnv),
	}, nil
}

func setCredentials(env map[string]string) Credentials {
	return Credentials{
		userId:       env["USER_ID"],
		clientId:     env["CLIENT_ID"],
		clientSecret: env["CLIENT_SECRET"],
		accessToken:  env["USER_TOKEN"],
	}
}

type accessToken struct {
	AccessToken string `json:"access_token"`
}

// Method to get an app access token if not able to use a user access token.
// Needs a client id and a client secret to be able to generate the token.
// Limited parts of the api can be used without a user access token.
func (c *Client) getAppAccessToken() error {
	body := strings.NewReader(fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=client_credentials", c.credentials.clientId, c.credentials.clientSecret))
	req, err := http.NewRequest("POST", "https://id.twitch.tv/oauth2/token", body)
	if err != nil {
		return err
	}

	req.Header = http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	t := &accessToken{}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(t)
	if err != nil {
		return err
	}

	c.credentials.accessToken = *&t.AccessToken

	return nil
}
