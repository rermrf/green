//go:build wireinject

package startup

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
		InitDB,
		dao.NewUserDao,
		repository.NewCachedUserRepository,
		service.NewUserService,
		handler.NewUserHandler,
		ioc.InitGin,
		ioc.InitMiddlewares,
	)
	return new(gin.Engine)
}
