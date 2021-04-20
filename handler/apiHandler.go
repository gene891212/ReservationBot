package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"linebot-server/stru"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SentReservation(c *gin.Context) {
	accessToken := c.PostForm("accessToken")
	fmt.Println(accessToken)
	fmt.Println(c.PostForm("reciver"))
	send := fmt.Sprintf("%v %v", c.PostForm("date"), c.PostForm("time"))
	sendTime, err := time.Parse("2006-01-02 15:04", send)
	if err != nil {
		panic(err)
	}
	fmt.Println(sendTime)

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
}
