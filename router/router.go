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

	router.POST("/callback", handler.Callback)
	router.GET("/test", handler.Index)

	return router
}
