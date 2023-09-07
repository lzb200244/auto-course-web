package v1api

import (
	"auto-course-web/controller"
	"auto-course-web/router/middleware"
	"github.com/gin-gonic/gin"
)

/*
Created by 斑斑砖 on 2023/9/5.
Description：

*/

func SetupAdmin(route *gin.RouterGroup) {
	route.Use(middleware.IsAdmin())

	//创建新的页面
	route.POST("page", controller.CreatePageController)
	route.PUT("page", controller.ModifyPageController)
	//赋予权限
	route.PUT("auth", controller.AddAuthController)
	//删除权限
	route.DELETE("auth", controller.DelAuthController)
	//创建新的权限
	route.POST("auth", controller.CreateAuthController)
}
