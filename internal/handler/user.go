package handler

import (
	service "green/internal/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.POST("/signup", h.Signup)
}

func (h *UserHandler) Signup(ctx *gin.Context) {
	ctx.String(http.StatusOK, "signup")
}
