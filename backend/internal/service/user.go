package service

import (
	"github.com/gin-gonic/gin"
	"web-golang/backend/internal/domain"
	"web-golang/backend/internal/repo"
)

type UserService struct {
	repo *repo.UserRepo
}

func NewUserService(repo *repo.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (t *UserService) SignUp(ctx *gin.Context, do domain.User) {
	t.repo.CreateUser(ctx, do)
}
