package v1api

import (
	"auto-course-web/controller"
	"github.com/gin-gonic/gin"
)

/*
Created by 斑斑砖 on 2023/9/6.
	Description：

*/

func SetupCourse(group *gin.RouterGroup) {
	// 创建课程
	teacher := group.Group("teacher")
	{
		teacher.POST("/", controller.CreateCourseController)
		//teacher.GET("/", func(context *gin.Context) {
		//	utils.Success(context, code.GetMsg(code.OK), nil)
		//})
		teacher.GET("/", controller.ListCourseController)
	}
	//student := group.Group("/student")
	//{
	//	student.POST("/", controller.CreateCourseController)
	//	student.GET("/", controller.CreateCourseController)
	//}
}
