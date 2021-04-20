package handler

import (
	"fmt"
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
	res, err := client.Get("https://api.line.me/v2/profile")
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", "https://api.line.me/v2/profile", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Bearer", accessToken)
	res, err = client.Do(req)

	fmt.Println(res)
}
