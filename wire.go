//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"green/internal/handler"
	"green/internal/ioc"
	"green/internal/repository"
	"green/internal/repository/dao"
	"green/internal/service"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		ioc.InitGin,
		ioc.InitMiddlewares,
		ioc.InitDB,

		dao.NewUserDao,

		repository.NewCachedUserRepository,

		service.NewUserService,

		handler.NewUserHandler,
	)
	return new(gin.Engine)
}
