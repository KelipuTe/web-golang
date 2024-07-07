package middleware

import (
	"github.com/gin-gonic/gin"
)

type SessionBuilder struct {
}

func NewSessionBuilder() *SessionBuilder {
	return &SessionBuilder{}
}

func (t *SessionBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
