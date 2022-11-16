package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	baseUrl = "https://api.twitch.tv/helix"
)

func main() {
	client, err := newClient()
	if err != nil {
		panic(err)
	}

	liveStreams, err := client.getLiveFollowing()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", liveStreams.Data)
}

func (c *Client) getUserID(username string) error {
	url := fmt.Sprintf("%s/users?login=%s", baseUrl, username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", c.credentials.accessToken)},
		"Client-Id":     {c.credentials.clientId},
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	id := struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&id)
	if err != nil {
		return err
	}

	c.credentials.userId = id.Data[0].ID

	return nil
}

type Following struct {
	Data []struct {
		UserName    string `json:"user_name"`
		UserLogin   string `json:"user_login"`
		GameName    string `json:"game_name"`
		Title       string `json:"title"`
		ViewerCount int    `json:"viewer_count"`
	} `json:"data"`
}

func (c *Client) getLiveFollowing() (*Following, error) {
	url := fmt.Sprintf("%s/streams/followed?user_id=%s", baseUrl, c.credentials.userId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", c.credentials.accessToken)},
		"Client-Id":     {c.credentials.clientId},
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	followerList := Following{}

	err = json.NewDecoder(resp.Body).Decode(&followerList)
	if err != nil {
		return nil, err
	}

	return &followerList, nil
}
