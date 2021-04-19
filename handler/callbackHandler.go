package handler

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func LineEchoBot(c *gin.Context) {
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
