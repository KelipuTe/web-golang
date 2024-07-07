package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PO
type User struct {
	ID       int64  `gorm:""`
	Email    string `gorm:""`
	Password string
}

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (t *UserDao) InsertUser(ctx *gin.Context, po User) error {
	return t.db.WithContext(ctx).Create(&po).Error
}
