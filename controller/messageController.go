package controller

import (
	"database/sql"
	"fmt"
	"linebot-server/database"
	"linebot-server/models"
	"time"

	"github.com/gin-gonic/gin"
)

type MessageRepo struct {
	Db *sql.DB
}

func NewMessage() MessageRepo {
	db, err := database.InitDb()
	if err != nil {
		panic(err)
	}
	repo := MessageRepo{Db: db}
	return repo
}

func (repo *MessageRepo) CreateMessageHandler(c *gin.Context) {
	senderProfile, err := models.GetUserByAccessToken(c.PostForm("accessToken"))
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"detail": "CreateMessageHandler get user by access token error",
			"error":  err.Error(),
		})
		return
	}

	sender, err := models.GetUser(repo.Db, senderProfile.DisplayName)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"detail": "CreateMessageHandler get sender detail error",
			"error":  err.Error(),
		})
		return
	}

	reciver, err := models.GetUser(repo.Db, c.PostForm("reciver"))
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"detail": "CreateMessageHandler get reciver detail error",
			"error":  err.Error(),
		})
		return
	}

	content := c.PostForm("content")
	datetimeStr := fmt.Sprintf("%v %v", c.PostForm("date"), c.PostForm("time"))
	datetime, err := time.Parse("2006-01-02 15:04", datetimeStr)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"detail": "CreateMessageHandler parse datetime error",
			"error":  err.Error(),
		})
		return
	}

	message := models.Message{
		Content: content,
		Sender:  sender,
		Reciver: reciver,
		Time:    datetime,
	}
	err = models.CreateMessage(repo.Db, message)
	if err != nil {
		c.JSONP(500, gin.H{
			"status": 500,
			"detail": "CreateMessageHandler create message error",
			"error":  err.Error(),
		})
		return
	}

}
