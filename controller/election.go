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
	validate, err := utils.BindValidQuery[request.SelectCourse](ctx)
	if err != nil {
		utils.Fail(ctx, code.ERROR_REQUEST_PARAM, err.Error(), nil)
		return
	}
	// 2. 调用服务

	data, c := service.ListSelectCourse(&validate)
	if c != code.OK {
		utils.Fail(ctx, c, code.GetMsg(c), nil)
		return
	}
	utils.Success(
		ctx, code.GetMsg(c), data,
	)

}
