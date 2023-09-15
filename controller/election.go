package controller

import (
	"auto-course-web/global/code"
	"auto-course-web/models/request"
	"auto-course-web/service"
	"auto-course-web/utils"
	"github.com/gin-gonic/gin"
)

/*
Created by 斑斑砖 on 2023/9/14.
Description：

*/
// ==================================================================== 获取抢课区的数据

func ListSelectCourseController(ctx *gin.Context) {
	validate, err := utils.BindValidQuery[request.SelectCourseReq](ctx)
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	// 2. 调用服务

	data, c := service.ListSelectCourses(&validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(
		ctx, code.GetMsg(c), data,
	)

}

// CreateSelectCourseController 选课
func CreateSelectCourseController(ctx *gin.Context) {

	validate, err := utils.BindValidJson[request.CreateCourseReq](ctx)
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	user, _ := utils.GetUser(ctx)
	// 2. 调用服务
	data, c := service.CreateSelectCourse(int(user.ID), &validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(code.OK), data)
}

// ==================================================================== 我选的课程

func ListMySelectCourseController(ctx *gin.Context) {
	user, _ := utils.GetUser(ctx)
	// 2. 调用服务
	data, c := service.ListMySelectCourses(int(user.ID))
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(code.OK), data)
}
