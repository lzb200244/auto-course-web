package router

import (
	"auto-course-web/controller"
	"auto-course-web/global"
	"auto-course-web/global/code"
	"auto-course-web/router/middleware"
	"auto-course-web/router/v1api"
	"auto-course-web/utils"
	"auto-course-web/utils/qiniu"
	"github.com/gin-gonic/gin"
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
		v1.POST("users/register", controller.RegisterController)
		v1.POST("users/login", controller.LoginController)
		// ==================================================================== 需要进行认证的
		authored := v1.Group("")
		authored.Use(middleware.JWT())
		{
			credit := authored.Group("credits")
			{
				credit.GET("kodo", func(context *gin.Context) {
					utils.Success(context, code.GetMsg(code.OK), qiniu.GetCredits())
				})
			}
			user := authored.Group("users")
			v1api.SetupUser(user)
			admin := authored.Group("admin")
			v1api.SetupAdmin(admin)
			course := authored.Group("courses")
			v1api.SetupCourse(course)

		}

	}

	return router
}
