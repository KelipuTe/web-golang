package repo

import (
	"github.com/gin-gonic/gin"
	"web-golang/backend/internal/domain"
	"web-golang/backend/internal/repo/dao"
)

type UserRepo struct {
	dao *dao.UserDao
}

func NewUserRepo(dao *dao.UserDao) *UserRepo {
	return &UserRepo{
		dao: dao,
	}
}

func (t *UserRepo) CreateUser(ctx *gin.Context, do domain.User) error {
	return t.dao.InsertUser(ctx, dao.User{
		Email:    do.Email,
		Password: do.Password,
	})
}
