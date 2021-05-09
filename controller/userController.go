package controller

import (
	"database/sql"
	"linebot-server/database"
	"linebot-server/models"

	"github.com/gin-gonic/gin"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUser() UserRepo {
	db, err := database.InitDb()
	if err != nil {
		panic(err)
	}
	repo := UserRepo{Db: db}
	return repo
}

func (repo *UserRepo) GetUserHandler(c *gin.Context) {
	name, _ := c.Params.Get("name")
	user, err := models.GetUser(repo.Db, name)
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

func (repo *UserRepo) GetUsersHandler(c *gin.Context) {
	user, err := models.GetUsers(repo.Db)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"detail": "GetUsersHandler get users error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(200, user)
}

// func (handler *Repository) CreateUserHandler(c *gin.Context) {

// }
