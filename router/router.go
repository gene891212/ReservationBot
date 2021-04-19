package router

import (
	"linebot-server/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	if gin.Mode() == gin.TestMode {
		router.LoadHTMLGlob("../view/*")
		router.Static("/static", "../static/*/*")
	} else {
		router.LoadHTMLGlob("view/*")
		router.Static("/static", "static")
	}

	router.GET("/", handler.LiffPage)
	router.POST("/callback", handler.LineEchoBot)

	api := router.Group("/api")
	api.POST("/submit", handler.SentReservation)
	return router
}
