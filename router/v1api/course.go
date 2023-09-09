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

		teacher.GET("/", controller.ListCourseController)
		//教师发布课程到缓存
		teacher.POST("/publish", controller.PublishCourseController)
		teacher.DELETE("/publish", controller.CancelPublishCourseController)
		teacher.GET("/publish", controller.PublishListCourseController)
	}

}
