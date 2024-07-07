package main

import (
	"github.com/gin-gonic/gin"
	"web-golang/backend/internal/api"
	"web-golang/backend/internal/api/middleware"
	"web-golang/backend/internal/ioc"
	"web-golang/backend/internal/repo"
	"web-golang/backend/internal/repo/dao"
	"web-golang/backend/internal/service"
)

func main() {
	backServer := gin.Default()
	backServer.Use(middleware.NewCormBuilder().Build())

	db := ioc.InitDB()
	udao := dao.NewUserDao(db)
	urepo := repo.NewUserRepo(udao)
	usvc := service.NewUserService(urepo)
	uhdl := api.NewUserHandler(usvc)

	ug := backServer.Group("api/user")
	ug.POST("/signup", uhdl.SignUp)
	ug.POST("/login", uhdl.Login)
	ug.POST("/logout", uhdl.Logout)

	backServer.Run(":8502")
}
