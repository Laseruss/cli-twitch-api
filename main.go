package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Credentials struct {
	clientId     string
	clientSecret string
	userToken    string
}

func setCredentials(env map[string]string) Credentials {
	return Credentials{
		clientId:     env["CLIENT_ID"],
		clientSecret: env["CLIENT_SECRET"],
		userToken:    env["USER_TOKEN"],
	}
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

func main() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		fmt.Printf("could not load env variables: %s\n", err)
		os.Exit(1)
	}

	credentials := setCredentials(myEnv)

	fmt.Printf("%+v\n", credentials)

}

// Other stuffs

type oauthToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

/*
func (c *credentials) getOauthToken() error {
	client := http.Client{}

	body := strings.NewReader(fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=client_credentials", c.id, c.secret))
	req, err := http.NewRequest("POST", "https://id.twitch.tv/oauth2/token", body)
	if err != nil {
		return err
	}

	req.Header = http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	t := &oauthToken{}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(t)
	if err != nil {
		return err
	}

	c.token = *t

	return nil
}
*/

// Request follower list code
/*
	client := http.Client{}

	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/streams/followed?user_id=104328365", nil)
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", userToken)},
		"Client-Id":     {clientID},
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	followerList := Following{}

	err = json.NewDecoder(res.Body).Decode(&followerList)
	if err != nil {
		panic(err)
	}

	for _, item := range followerList.Data {
		fmt.Printf("%s is playing %s to %d viewers\n", item.UserName, item.GameName, item.ViewerCount)
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Link to stream: https://twitch.tv/%s\n", item.UserLogin)
		fmt.Println()
		fmt.Println("----------------------------------------------------------------")
		fmt.Println()
	}
*/
