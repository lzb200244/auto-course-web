package v1api

import (
	"auto-course-web/controller"
	"auto-course-web/global/auth"
	"auto-course-web/router/middleware"
	"github.com/gin-gonic/gin"
)

/*
Created by 斑斑砖 on 2023/9/5.
Description：

*/

func SetupAdmin(group *gin.RouterGroup) {
	group.Use(middleware.HasRole(auth.Admin))
	//创建新的路由页面
	group.POST("routes", controller.CreatePageController)
	//修改页面
	group.PUT("routes", controller.ModifyPageController)
	//创建课程分类
	group.POST("categories", controller.CreateCategoryController)
	//赋予权限
	group.PUT("permissions", controller.AddAuthController)
	//删除权限
	group.DELETE("permissions", controller.DelAuthController)
	//创建新的权限
	group.POST("permissions", controller.CreateAuthController)
	//通知教师进行选课
	group.POST("/teachers/notify", controller.Notice2TeacherController)
	//通知学生进行选课
	group.POST("/students/notify", controller.Notice2StudentController)

	//	获取所有的预选课程
	group.GET("courses", controller.GetAllCourseController)

}
