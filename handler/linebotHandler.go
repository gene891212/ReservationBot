package handler

import (
	"database/sql"
	"linebot-server/lib"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func LineEchoBot(c *gin.Context) {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}

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
		if event.Type == linebot.EventTypeFollow {
			userId := event.Source.UserID
			profile, err := bot.GetProfile(userId).Do()
			if err != nil {
				panic(err)
			}
			lib.InsertUserToDB(db, profile)
		} else if event.Type == linebot.EventTypeMessage {
			switch msg := event.Message.(type) {
			case *linebot.TextMessage:
				_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg.Text)).Do()
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
