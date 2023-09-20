package service

import (
	"auto-course-web/global"
	"auto-course-web/initialize"
	"auto-course-web/models"
	"testing"
)

/*
Created by 斑斑砖 on 2023/9/16.
Description：
*/
type Resp struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc" `
	Code    string `json:"code"`
	Teacher string `json:"teacher" `
	Cover   string `json:"cover"`
	//	学分
	Credit uint32 `json:"credit" `
	//课程分类
	CategoryID uint `json:"categoryID" `
	Category   *models.CourseCategory

	//	上课时间段
	Schedule string `json:"schedule"`

	IsPreLoad bool `json:"isPreLoad" gorm:"-"`
	//	开课时间
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime" `
}

func TestSelectCourse(t *testing.T) {
	initialize.InitConfig("../config/dev.conf.yml")
	initialize.InitLogger()
	initialize.InitMysql()
	resp := Resp{}
	global.MysqlDB.Model(models.Course{}).Preload("Category").Preload("Category.College").Find(&resp)
	//marshal, err := json.Marshal(resp)
	//fmt.Println(string(marshal))
}
