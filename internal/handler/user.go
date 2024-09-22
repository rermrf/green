package handler

import (
	"errors"
	regexp "github.com/dlclark/regexp2"
	"green/internal/domain"
	"green/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc         service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
	phoneExp    *regexp.Regexp
}

func NewUserHandler(svc service.UserService) *UserHandler {
	const (
		emailRegexPattern    = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*?&.])[A-Za-z\d@$!%*?&.]{8,72}$`
		phoneRegexPattern    = `^1[3-9]\d{9}$`
	)

	return &UserHandler{
		svc:         svc,
		emailExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		phoneExp:    regexp.MustCompile(phoneRegexPattern, regexp.None),
	}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.POST("/signup", h.Signup)
}

func (h *UserHandler) Login(ctx *gin.Context) {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(errors.New("参数格式错误"))
		return
	}
}

func (h *UserHandler) Signup(ctx *gin.Context) {
	type SignUpRequest struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req SignUpRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(errors.New("参数格式错误"))
		return
	}
	ok, err := h.emailExp.MatchString(req.Email)
	if err != nil {
		// 邮箱匹配错误
		ctx.Error(errors.New("系统错误"))
		return
	}

	if !ok {
		// 邮箱格式不正确
		ctx.Error(errors.New("邮箱格式不正确"))
		return
	}

	if req.Password != req.ConfirmPassword {
		// 两次密码不一致
		ctx.Error(errors.New("两次密码不一致"))
		return
	}

	ok, err = h.passwordExp.MatchString(req.Password)
	if err != nil {
		// TODO: 记录日志
		// 密码匹配错误
		ctx.Error(errors.New("系统错误"))
		return
	}

	if !ok {
		// 密码格式不正确
		ctx.Error(errors.New("密码格式不正确"))
		return
	}
	err = h.svc.Signup(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if errors.Is(err, service.ErrUserDuplicate) {
		ctx.Error(errors.New("邮箱已经注册"))
		return
	}
	if err != nil {
		ctx.Error(errors.New("系统错误"))
		return
	}
	ctx.JSON(http.StatusOK, Result[string]{Msg: "注册成功"})
}
