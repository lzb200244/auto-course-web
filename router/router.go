package router

import (
	"auto-course-web/controller"
	"auto-course-web/global"
	"auto-course-web/global/code"
	"auto-course-web/models"
	"auto-course-web/router/middleware"
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
		v1.POST("register", controller.RegisterController)
		v1.POST("login", controller.LoginController)

		// ====================================================================
		//需要进行认证的
		authored := v1.Group("/")

		authored.Use(middleware.JWT())
		{
			authored.GET("user", controller.GetUserController)
			authored.PUT("user", controller.UpdateInfoController)
			authored.GET("permission", func(context *gin.Context) {
				var routes []*models.Router
				user, _ := utils.GetUser(context)
				global.MysqlDB.
					Where("`limit`<=?", user.Authority).Find(&routes)

				mpRoute := make(map[int]*models.Router, len(routes))
				for _, route := range routes {
					m := mpRoute[int(route.Parent)]
					if m != nil {
						m.Children = append(m.Children, route)
					}
					mpRoute[int(route.ID)] = route
				}
				routes = []*models.Router{}
				if mpRoute[1] != nil && mpRoute[1].Children != nil {
					routes = mpRoute[1].Children
				}
				//返回根下的所有路由
				utils.Success(context, code.GetMsg(code.OK), routes)
			})

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
			//创建新的页面
			admin.POST("page", controller.CreatePageController)
			admin.PUT("page", controller.ModifyPageController)
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
