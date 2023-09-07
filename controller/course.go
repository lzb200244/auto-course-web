package controller

import (
	"auto-course-web/global/code"
	"auto-course-web/models/request"
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
	utils.Results(
		ctx, validate.Page, validate.Size, 10, code.GetMsg(c), data,
	)
}
