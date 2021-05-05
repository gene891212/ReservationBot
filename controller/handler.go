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

type Repository struct {
	Db *sql.DB
}

func New() Repository {
	db, err := database.InitDb()
	if err != nil {
		panic(err)
	}
	handler := Repository{Db: db}
	return handler
}

func (handler *Repository) GetUserHandler(c *gin.Context) {
	name, _ := c.Params.Get("name")
	user, err := models.GetUser(handler.Db, name)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(400, gin.H{
				"status":  400,
				"message": "name not found",
			})
			return
		} else {
			c.JSON(500, gin.H{
				"status": 500,
				"error":  err.Error(),
			})
			return
		}
	}
	c.JSON(200, user)
}

func (handler *Repository) GetUsersHandler(c *gin.Context) {
	user, err := models.GetUsers(handler.Db)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(200, user)
}

// func (handler *Repository) CreateUserHandler(c *gin.Context) {

// }

func (handler *Repository) EchoBot(c *gin.Context) {
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
				c.JSON(500, gin.H{
					"status":  500,
					"message": err.Error(),
				})
				return
			}
			models.CreateUser(handler.Db, profile)
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
