package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// CORMBuilder 处理跨域
type CORMBuilder struct {
}

func NewCormBuilder() *CORMBuilder {
	return &CORMBuilder{}
}

func (t *CORMBuilder) Build() gin.HandlerFunc {
	return cors.New(cors.Config{
		//允不允许带用户认证信息（cookie）
		AllowCredentials: true,
		//允许的http请求方法
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		//允许带的http请求头
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		//允许的来源
		//AllowOrigins: []string{},
		//允许的来源（用自定义方法解析请求头的origin字段）
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return false
		},
		MaxAge: 12 * time.Hour,
	})
}
