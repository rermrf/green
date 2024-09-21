package ioc

import (
	"green/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitGin(mdls []gin.HandlerFunc, userHdl *handler.UserHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	userHdl.RegisterRoutes(server)
	return server
}

func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}
