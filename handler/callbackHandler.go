package handler

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	config = mysql.Config{
		User:   os.Getenv("user"),
		Passwd: os.Getenv("passwd"),
		Net:    os.Getenv("tcp"),
		Addr:   os.Getenv("addr"),
		DBName: os.Getenv("dbname"),
	}
)

func Callback(c *gin.Context) {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SCRECT"),
		os.Getenv("ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		panic(err)
	}
	for _, event := range events {
		switch msg := event.Message.(type) {
		case *linebot.TextMessage:
			_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg.Text)).Do()
			if err != nil {
				panic(err)
			}
		}
	}
}
