package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/line/line-bot-sdk-go/linebot"
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
		c.Status(500)
		return
	}
	for _, event := range events {
		switch msg := event.Message.(type) {
		case *linebot.TextMessage:
			_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg.Text)).Do()
			if err != nil {
				c.Status(500)
				return
			}
		}
	}
}

func main() {
	router := gin.Default()
	router.POST("/callback", Callback)
	err := router.Run(":7000")
	if err != nil {
		log.Fatal(err)
	}
}
