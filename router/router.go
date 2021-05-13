package router

import (
	"linebot-server/controller"

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

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{})
	})

	userRepo := controller.NewUser()
	linebotRepo := controller.NewLinebot()
	messageRepo := controller.NewMessage()

	router.POST("/callback", linebotRepo.EchoBot)

	api := router.Group("/api")
	// api.POST("/submit", messageRepo)
	api.GET("/users", userRepo.GetUsersHandler)
	api.GET("/user/:name", userRepo.GetUserHandler)

	api.POST("/message", messageRepo.CreateMessageHandler)
	return router
}
