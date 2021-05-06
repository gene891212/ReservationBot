package router

import (
	"linebot-server/controller"
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
	userRepo := controller.NewUser()
	linebotRepo := controller.NewLinebot()

	router.GET("/", handler.LiffPage)
	router.POST("/callback", linebotRepo.EchoBot)

	api := router.Group("/api")
	// api.POST("/submit", handler.SentReservation)
	api.GET("/users", userRepo.GetUsersHandler)
	api.GET("/user/:name", userRepo.GetUserHandler)
	return router
}
