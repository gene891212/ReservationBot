package router

import (
	"linebot-server/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/callback", handler.Callback)

	return router
}
