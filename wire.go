//go:build wireinject

package main

import (
	"green/internal/handler"
	"green/internal/ioc"
	service "green/internal/service/user"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		ioc.InitGin,
		ioc.InitMiddlewares,

		service.NewUserService,

		handler.NewUserHandler,
	)
	return new(gin.Engine)
}
