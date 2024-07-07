package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	frontServer := gin.Default()
	frontServer.LoadHTMLGlob("frontend/*")
	frontServer.GET("/user/signup", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signup.html", gin.H{})
	})

	frontServer.Run(":8501")
}
