package controller

import (
	"auto-course-web/global/code"
	"auto-course-web/models/request"
	"auto-course-web/service"
	"auto-course-web/utils"
	"github.com/gin-gonic/gin"
)

/*
Created by 斑斑砖 on 2023/8/14.
Description：
	注册
*/

// =================================================================== 登录注册

// RegisterController 注册
func RegisterController(ctx *gin.Context) {
	//参数校验
	validate, err := utils.BindValidJson[request.Register](ctx)
	//参数校验失败
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	_, c := service.Register(&validate)
	if c != code.OK {

		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), nil)
}

// LoginController 登录
func LoginController(ctx *gin.Context) {
	//1. 参数校验
	validate, err := utils.BindValidJson[request.Login](ctx)
	//参数校验失败
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	//2. 调用服务
	data, c := service.Login(validate.Username, validate.Password)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), data)
}

// SendEmailController 发送邮件
func SendEmailController(ctx *gin.Context) {
	validate, err := utils.BindValidJson[request.SendEmail](ctx)
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	data, c := service.SendEmail(&validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), data)

}

// =================================================================== 用户相关

// GetUserController 获取用户
func GetUserController(ctx *gin.Context) {
	//1. 从ctx提取用户id
	user, _ := utils.GetUser(ctx)

	//2. 无需进行校验， 调用服务
	data, c := service.GetUserInfo(int(user.ID), user.Role)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), data)
}

// UpdateInfoController 修改用户信息
func UpdateInfoController(ctx *gin.Context) {
	user, _ := utils.GetUser(ctx)
	//参数校验
	validate, err := utils.BindValidJson[request.UserInfo](ctx)
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	//调用服务
	data, c := service.UpdateInfo(user.ID, &validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), data)
}
