package lib

import (
	"encoding/json"
	"io/ioutil"
	"linebot-server/stru"
	"net/http"
)

func GetUserProfile(accessToken string) stru.UserInfo {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.line.me/v2/profile", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var u stru.UserInfo
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}

	return u
}
