package handler

import (
	"fmt"
	"linebot-server/lib"
	"time"

	"github.com/gin-gonic/gin"
)

func SentReservation(c *gin.Context) {
	send := fmt.Sprintf("%v %v", c.PostForm("date"), c.PostForm("time"))
	sendTime, err := time.Parse("2006-01-02 15:04", send)
	if err != nil {
		panic(err)
	}
	fmt.Println(sendTime)
	fmt.Println(c.PostForm("reciver"))
	fmt.Println(c.PostForm("message"))

	accessToken := c.PostForm("accessToken")
	userProfile := lib.GetUserProfile(accessToken)
	fmt.Println(userProfile)
	c.JSON(200, userProfile)
}
