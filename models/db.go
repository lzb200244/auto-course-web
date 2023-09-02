package models

import (
	"go-template/global"
	"go.uber.org/zap"
	"os"
)

/*
Created by 斑斑砖 on 2023/8/13.
Description：
	表迁移
*/

func Migrate() {

	err := global.MysqlDB.AutoMigrate(
		User{}, Role{}, Permission{},
	)
	/*
		//创建角色 #########################################################################################
		//清空表
		//global.MysqlDB.Where("1 = 1").Delete(Role{})
		roleList := []string{"admin", "opt", "user"}
		for _, v := range roleList {
			global.MysqlDB.Create(&Role{
				Name: v,
			})
		}

		//创建权限 #########################################################################################
		permissionList := []string{"create", "update", "delete"}
		for _, v := range permissionList {
			global.MysqlDB.Create(&Permission{
				Name: v,
			})
		}

	*/
	if err != nil {
		global.Logger.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}
}
