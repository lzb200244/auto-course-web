package v1api

import (
	"auto-course-web/controller"
	"auto-course-web/global"
	"auto-course-web/global/code"
	"auto-course-web/models"
	"auto-course-web/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
Created by 斑斑砖 on 2023/9/5.
Description：

*/

func SetupUser(group *gin.RouterGroup) {
	// 获取用户信息
	//group.GET("/:id", controller.GetUserController)
	group.GET("/", controller.GetUserController)
	// 用户更新信息
	group.PUT("/", controller.UpdateInfoController)
	// 获取用户权限
	group.GET("/permission", func(context *gin.Context) {
		var routes []*models.Router
		user, _ := utils.GetUser(context)
		fmt.Println(user.Role)
		global.MysqlDB.
			Where("role<=?", user.Role).Find(&routes)

		mpRoute := make(map[int]*models.Router, len(routes))
		for _, route := range routes {
			m := mpRoute[int(route.Parent)]
			if m != nil {
				m.Children = append(m.Children, route)
			}
			mpRoute[int(route.ID)] = route
		}
		routes = []*models.Router{}
		if mpRoute[1] != nil && mpRoute[1].Children != nil {
			routes = mpRoute[1].Children
		}
		//返回根下的所有路由
		utils.Success(context, code.GetMsg(code.OK), routes)
	})
}
