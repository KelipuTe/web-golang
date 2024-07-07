package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-golang/backend/internal/domain"
	"web-golang/backend/internal/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

type signUpDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Confirm  string `json:"confirm"`
}

func (t *UserHandler) SignUp(ctx *gin.Context) {
	dto := &signUpDTO{}

	if err := ctx.Bind(&dto); err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
	}
	if dto.Password != dto.Confirm {
		ctx.String(http.StatusUnprocessableEntity, "两次输入的密码不同")
	}

	t.svc.SignUp(ctx, domain.User{
		Email:    dto.Email,
		Password: dto.Password,
	})

	ctx.String(http.StatusOK, "注册成功")
}

func (t *UserHandler) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"hello": "hello"})
}

func (t *UserHandler) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"hello": "hello"})
}
