package v1api

import (
	"auto-course-web/controller"
	"auto-course-web/global/auth"
	"auto-course-web/router/middleware"
	"github.com/gin-gonic/gin"
)

/*
Created by 斑斑砖 on 2023/9/6.
	Description：课程的curd，课程分类，教师进行预发布的路由控制

*/

func SetupCourse(group *gin.RouterGroup) {
	t := group.Use(middleware.HasRole(auth.Teacher))
	// 获取课程
	t.GET("/", controller.ListCourseController)
	// 创建课程
	t.POST("/", controller.CreateCourseController)
	// 编辑课程
	t.PUT("/", controller.UpdateCourseController)
	//获取课程详细
	t.GET("detail/:courseID", controller.DetailCourseController)

	//获取课程分类
	t.GET("/category", controller.ListCourseCategoryController)
	//获取时间表
	t.GET("/schedule", controller.ListCourseScheduleController)
	//教师发布课程到缓存
	t.POST("/publish", controller.PublishCourseController)
	t.DELETE("/publish", controller.CancelPublishCourseController)
	t.GET("/publish", controller.PublishListCourseController)

	//选课列表
	//t.GET("/selects/", controller.ListSelectCourseController)

}
