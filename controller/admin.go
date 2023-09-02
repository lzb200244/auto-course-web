package controller

import (
	"github.com/gin-gonic/gin"
	"go-template/global/code"
	"go-template/models/request"
	"go-template/service"
	"go-template/utils"
)

/*
Created by 斑斑砖 on 2023/9/2.
Description：
	权限的curd
*/

// AddAuthController 角色添加权限
func AddAuthController(ctx *gin.Context) {
	//参数校验
	validate, err := utils.BindValidJson[request.Auths](ctx)
	//参数校验失败
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	_, c := service.SetAuth(validate.RoleID, validate.Permission)
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
	_, c := service.DelAuth(validate.RoleID, validate.PermissionID)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), nil)
}

func CreateAuthController(ctx *gin.Context) {
	//参数校验
	validate, err := utils.BindValidJson[request.Permission](ctx)
	//参数校验失败
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	_, c := service.CreatePermission(validate.Name)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(ctx, code.GetMsg(c), nil)

}
