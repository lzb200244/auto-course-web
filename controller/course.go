package controller

import (
	"auto-course-web/global/code"
	"auto-course-web/models/request"
	"auto-course-web/models/response"
	"auto-course-web/service"
	"auto-course-web/utils"
	"github.com/gin-gonic/gin"
)

/*
Created by 斑斑砖 on 2023/9/6.
	Description：
		课程相关的/创建课程/课程列表/我加入的课程等
*/

func CreateCourseController(ctx *gin.Context) {
	// TODO 用户权限校验 ,只允许老师/管理员创建 这一步应该在中间层进行,

	// 1. 参数校验
	validate, err := utils.BindValidJson[request.Course](ctx)
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	user, _ := utils.GetUser(ctx)
	// 2. 调用服务
	_, c := service.CreateCourse(int(user.ID), &validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return

	}
	utils.Success(ctx, code.GetMsg(c), nil)

}

// ListCourseController 获取课程我创建的列表
func ListCourseController(ctx *gin.Context) {
	validate, err := utils.BindValidQuery[request.Pages](ctx)
	if err != nil {

		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	user, _ := utils.GetUser(ctx)
	// 2. 调用服务
	data, c := service.ListCourses(int(user.ID), &validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	d := data.(*response.List)
	utils.Results(
		ctx, int(d.Count), code.GetMsg(c), d.Data,
	)
}

// PublishCourseController 发布课程到缓存预热
func PublishCourseController(ctx *gin.Context) {
	validate, err := utils.BindValidJson[request.PreloadCourse](ctx)
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	_, c := service.PreLoadCourse2Redis(&validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(code.OK), nil)
}
func PublishListCourseController(ctx *gin.Context) {
	validate, err := utils.BindValidQuery[request.Pages](ctx)
	if err != nil {

		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	user, _ := utils.GetUser(ctx)
	// 2. 调用服务
	data, c := service.ListPublishCourses(int(user.ID), &validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	d := data.(*response.List)
	utils.Results(
		ctx, int(d.Count), code.GetMsg(c), d.Data,
	)
}

func CancelPublishCourseController(ctx *gin.Context) {
	validate, err := utils.BindValidJson[request.CancelPublishCourse](ctx)
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	user, _ := utils.GetUser(ctx)
	_, c := service.CancelCourse2Redis(int(user.ID), &validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(code.OK), nil)
}
