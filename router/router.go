package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-template/controller"
	"go-template/global"
	"go-template/router/middleware"
	"go-template/utils"
)

/*
Created by 斑斑砖 on 2023/8/14.
Description：

	路由
*/

func InitApiRouter() *gin.Engine {
	var router *gin.Engine
	if global.Config.Project.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		router = gin.Default()

	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
		router.Use(gin.Logger())
	}
	v1 := router.Group("api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			utils.Success(ctx, "success/", "pong")
		})
		v1.POST("register", controller.RegisterController)
		v1.POST("login", controller.LoginController)

		// ===========================================================================================
		//需要进行认证的
		authored := v1.Group("/")
		{
			authored.Use(middleware.JWT())
			authored.GET("authored", func(ctx *gin.Context) {
				user, _ := utils.GetUser(ctx)
				fmt.Println(user)
				utils.Success(ctx, "authored", nil)
			})
		}

	}

	return router
}
