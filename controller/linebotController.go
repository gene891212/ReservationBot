package controller

import (
	"database/sql"
	"linebot-server/database"
	"linebot-server/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type LinebotRepo struct {
	Db *sql.DB
}

func NewLinebot() LinebotRepo {
	db, err := database.InitDb()
	if err != nil {
		panic(err)
	}
	repo := LinebotRepo{Db: db}
	return repo
}

func (repo *LinebotRepo) EchoBot(c *gin.Context) {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SCRECT"),
		os.Getenv("ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeFollow {
			userId := event.Source.UserID
			profile, err := bot.GetProfile(userId).Do()
			if err != nil {
				c.JSON(500, gin.H{
					"status":  500,
					"message": err.Error(),
				})
				return
			}
			models.CreateUser(repo.Db, profile)
		} else if event.Type == linebot.EventTypeMessage {
			switch msg := event.Message.(type) {
			case *linebot.TextMessage:
				_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg.Text)).Do()
				if err != nil {
					c.JSON(500, gin.H{
						"status":  500,
						"message": err,
					})
					return
				}
			}
		}
	}
}
