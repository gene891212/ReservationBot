package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SentReservation(c *gin.Context) {

	fmt.Println(c.PostForm("reciver"))
}
