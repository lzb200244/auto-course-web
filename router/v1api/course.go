package v1api

import (
	"auto-course-web/controller"
	"auto-course-web/global/auth"
	"auto-course-web/router/middleware"
	"github.com/gin-gonic/gin"
)

/*
Created by 斑斑砖 on 2023/9/6.
	Description：

*/

func SetupCourse(group *gin.RouterGroup) {
	t := group.Use(middleware.HasRole(auth.Teacher))
	// 创建课程
	t.POST("/", controller.CreateCourseController)
	// 编辑课程
	t.GET("/", controller.ListCourseController)
	//获取课程分类
	t.GET("/category", controller.ListCourseCategoryController)

	//教师发布课程到缓存
	t.POST("/publish", controller.PublishCourseController)
	t.DELETE("/publish", controller.CancelPublishCourseController)
	t.GET("/publish", controller.PublishListCourseController)

}
