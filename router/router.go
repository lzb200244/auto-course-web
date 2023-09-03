package router

import (
	"github.com/gin-gonic/gin"
	"go-template/controller"
	"go-template/global"
	"go-template/global/code"
	"go-template/router/middleware"
	"go-template/utils"
	"go-template/utils/qiniu"
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
			utils.Success(ctx, "pong", nil)
		})
		v1.POST("register", controller.RegisterController)
		v1.POST("login", controller.LoginController)

		// ====================================================================
		//需要进行认证的
		authored := v1.Group("/")

		authored.Use(middleware.JWT())
		{
			authored.GET("user", controller.GetUserController)
			authored.PUT("user", controller.UpdateInfoController)
		}

		// =================================================================== 获取凭证
		credit := v1.Group("/credit")
		{
			credit.GET("kodo", func(context *gin.Context) {
				utils.Success(context, code.GetMsg(code.OK), qiniu.GetCredits())
			})
		}

		// =================================================================== 管理员赋予权限的相关curd
		admin := v1.Group("/admin")
		{
			admin.Use(middleware.JWT(), middleware.IsAdmin())
			//赋予权限
			admin.PUT("auth", controller.AddAuthController)
			//删除权限
			admin.DELETE("auth", controller.DelAuthController)
			//创建新的权限
			admin.POST("auth", controller.CreateAuthController)
		}

	}

	return router
}
