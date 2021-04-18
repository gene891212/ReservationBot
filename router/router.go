package router

import (
	"linebot-server/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	if gin.Mode() == gin.TestMode {
		router.LoadHTMLGlob("../view/*")
	} else {
		router.LoadHTMLGlob("view/*")
	}

	router.GET("/", handler.LiffPage)
	router.POST("/callback", handler.LinceEchoBot)

	api := router.Group("/api")
	api.POST("/submit", handler.SentReservation)
	return router
	// curl -d reciver=hihi http://localhost:7000/api/submit
}
