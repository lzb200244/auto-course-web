package controller

import (
	"auto-course-web/global/code"
	"auto-course-web/models/request"
	"auto-course-web/service"
	"auto-course-web/utils"
	"github.com/gin-gonic/gin"
)

/*
Created by 斑斑砖 on 2023/9/2.
Description：
	权限的curd
*/
// ============================================================= 权限相关

// AddAuthController 角色添加权限
func AddAuthController(ctx *gin.Context) {
	//参数校验
	validate, err := utils.BindValidJson[request.Auths](ctx)
	//参数校验失败
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	_, c := service.AddAuth(&validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), nil)
}

// DelAuthController 删除角色权限
func DelAuthController(ctx *gin.Context) {
	//参数校验
	validate, err := utils.BindValidJson[request.Auth](ctx)
	//参数校验失败
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	_, c := service.DelAuth(&validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), nil)
}

// CreateAuthController 创建新的权限
func CreateAuthController(ctx *gin.Context) {
	//参数校验
	validate, err := utils.BindValidJson[request.Permission](ctx)
	//参数校验失败
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	_, c := service.CreatePermission(&validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), nil)
}

// ============================================================== 路由相关

// CreatePageController 创建新的路由页面
func CreatePageController(ctx *gin.Context) {
	validate, err := utils.BindValidJson[request.Component](ctx)
	//参数校验失败
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	_, c := service.CreatePage(&validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return

	}
	utils.Success(ctx, code.GetMsg(c), nil)

}

// UpdatePageController 修改路由页面信息
func UpdatePageController(ctx *gin.Context) {
	validate, err := utils.BindValidJson[request.Component](ctx)
	//参数校验失败
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	_, c := service.UpdatePage(&validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return

	}
	utils.Success(ctx, code.GetMsg(c), nil)
}

// ============================================================== 通知相关

// Notice2TeacherController 通知教师进行预发布课程
func Notice2TeacherController(ctx *gin.Context) {
	_, c := service.Notice2Teacher()
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), nil)
}

// Notice2StudentController 通知学生进行选课
func Notice2StudentController(ctx *gin.Context) {
	_, c := service.Notice2Student()
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), nil)
}

// ============================================================== 课程相关

// ListPublishCourseController 获取所有的预选课程
func ListPublishCourseController(ctx *gin.Context) {
	validate, err := utils.BindValidQuery[request.Pages](ctx)
	//参数校验失败
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}

	data, c := service.ListPreloadCourse(&validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), data)
}

// ============================================================== 分类相关

// CreateCategoryController 创建新的课程分类
func CreateCategoryController(ctx *gin.Context) {
	validate, err := utils.BindValidJson[request.Category](ctx)
	//参数校验失败
	if err != nil {

		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	_, c := service.CreateCategory(&validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), nil)
}
