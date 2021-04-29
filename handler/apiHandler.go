package handler

import (
	"database/sql"
	"fmt"
	"linebot-server/lib"
	"time"

	"github.com/gin-gonic/gin"
)

func SentReservation(c *gin.Context) {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}
	send := fmt.Sprintf("%v %v", c.PostForm("date"), c.PostForm("time"))
	sendTime, err := time.Parse("2006-01-02 15:04", send)
	if err != nil {
		panic(err)
	}

	reciver := c.PostForm("reciver")
	content := c.PostForm("content")

	go lib.ReserveMessage(db, sendTime, reciver, content)

	// accessToken := c.PostForm("accessToken")
	// reciverProfile := lib.GetUserProfile(accessToken)
	// fmt.Println(userProfile)

	c.HTML(200, "submitSuccess.tmpl", gin.H{
		"reciver":  reciver,
		"datetime": sendTime.Format("2006-01-02 15:04"),
		"message":  content,
	})
}
