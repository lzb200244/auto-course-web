package respository

import (
	"auto-course-web/global"
	"auto-course-web/initialize"
	"auto-course-web/models"
	"auto-course-web/models/response"
	"encoding/json"
	"fmt"
	"testing"
)

/*
Created by 斑斑砖 on 2023/9/16.
Description：
*/

func TestQueryCourseWithCategoryList(t *testing.T) {
	initialize.InitConfig("../config/dev.conf.yml")
	initialize.InitLogger()
	initialize.InitMysql()
	var course response.CourseResponse
	global.MysqlDB.Model(models.Course{}).Preload("Category").Preload("College").Find(&course)
	marshal, _ := json.Marshal(&course)
	fmt.Println(string(marshal))
}
