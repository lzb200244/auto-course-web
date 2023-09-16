package controller

import (
	"auto-course-web/global/code"
	"auto-course-web/models/request"
	"auto-course-web/service"
	"auto-course-web/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
Created by 斑斑砖 on 2023/9/6.
	Description：
		课程相关的/创建课程/课程列表/我发布的课程/撤回发布课程等
*/
// =================================================================== 课程相关

// CreateCourseController 创建课程
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
	data, c := service.ListCourse(int(user.ID), &validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(
		ctx, code.GetMsg(c), data,
	)
}

// UpdateCourseController 修改课程信息
func UpdateCourseController(ctx *gin.Context) {
	validate, err := utils.BindValidJson[request.UpdateCourse](ctx)
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	user, _ := utils.GetUser(ctx)
	_, c := service.UpdateCourse(int(user.ID), &validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), nil)

}

func DetailCourseController(ctx *gin.Context) {
	courseID := ctx.Param("courseID")
	val, err := strconv.Atoi(courseID)
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, code.GetMsg(code.ERROR_REQUEST_PARAM), nil)
		return
	}
	data, c := service.DetailCourse(val)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), data)
}

// =================================================================== 课程预发布，缓存预热

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

// PublishListCourseController 获取我发布的课程
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
	utils.Success(
		ctx, code.GetMsg(c), data,
	)
}

// CancelPublishCourseController 取消发布
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

// ==================================================================== 课程分类

// ListCourseCategoryController 获取课程分类
func ListCourseCategoryController(ctx *gin.Context) {

	// 2. 调用服务
	data, c := service.ListCourseCategory()
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(
		ctx, code.GetMsg(c), data,
	)
}

// =================================================================== 时间表

// ListCourseScheduleController 获取课程分类
func ListCourseScheduleController(ctx *gin.Context) {

	// 2. 调用服务
	data, c := service.ListCourseSchedule()
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(
		ctx, code.GetMsg(c), data,
	)
}
