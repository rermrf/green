package middwares

import (
	"github.com/gin-gonic/gin"
	"green/internal/handler"
	"log"
	"net/http"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 捕获所有 panic 的错误
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic occurred: %v", r)
				// 返回通用的错误响应
				c.JSON(http.StatusOK, handler.Result[string]{
					Code: 5,
					Msg:  "系统错误",
				})
				c.Abort()
			}
		}()

		// 处理请求
		c.Next()

		// 检查 gin.Context 中是否有错误
		if len(c.Errors) > 0 {
			// 获取最后一个错误信息
			err := c.Errors.Last().Err
			log.Printf("error occurred: %v", err)

			// 返回自定义的错误响应
			c.JSON(http.StatusOK, handler.Result[string]{
				Code: 4,
				Msg:  err.Error(),
			})
			c.Abort()
		}
	}
}
