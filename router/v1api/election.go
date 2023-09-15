package v1api

import (
	"auto-course-web/controller"
	"github.com/gin-gonic/gin"
)

/*
Created by 斑斑砖 on 2023/9/14.
Description：

*/

func SetupElection(group *gin.RouterGroup) {
	group.GET("selects/", controller.ListSelectCourseController)
	group.POST("selects/", controller.CreateSelectCourseController)
	group.GET("selects/my", controller.ListMySelectCourseController)
}
