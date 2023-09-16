package models

import (
	"auto-course-web/global"
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
		User{}, Role{}, Permission{}, &Router{}, &College{},
		&CourseCategory{}, &Course{}, &CourseSchedule{},
	)
	if err != nil {
		global.Logger.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}
}
