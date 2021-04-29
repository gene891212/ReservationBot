package handler

import (
	"fmt"
	"linebot-server/lib"
	"time"

	"github.com/gin-gonic/gin"
)

func SentReservation(c *gin.Context) {
	// db, err := sql.Open("mysql", config.FormatDSN())
	// if err != nil {
	// 	panic(err)
	// }
	send := fmt.Sprintf("%v %v", c.PostForm("date"), c.PostForm("time"))
	sendTime, err := time.Parse("2006-01-02 15:04", send)
	if err != nil {
		panic(err)
	}

	reciver := c.PostForm("reciver")
	content := c.PostForm("content")

	go lib.ReserveMessage(sendTime, reciver, content)

	accessToken := c.PostForm("accessToken")
	userProfile := lib.GetUserProfile(accessToken)
	// fmt.Println(userProfile)

	// fmt.Printf("%+v", lib.GetUser(db, c.PostForm("reciver"))[0])
	c.JSON(200, userProfile)
}
