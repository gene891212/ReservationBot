package controller

import (
	"database/sql"
	"linebot-server/database"
	"linebot-server/models"

	"github.com/gin-gonic/gin"
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
