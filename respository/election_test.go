package respository

import (
	"auto-course-web/initialize"
	"auto-course-web/models"
	"fmt"
	"testing"
)

/*
Created by 斑斑砖 on 2023/9/14.
Description：
*/
func TestQuerySelectCourse(t *testing.T) {
	initialize.InitConfig("../config/dev.conf.yml")
	initialize.InitLogger()
	initialize.InitMysql()
	var course []models.Course
	var m models.Course
	m.Title = "数学课程"
	fmt.Println(course)
	//selectCourse, err := QuerySelectCourse(&models.Course{}, &course, nil, nil, m)
	//fmt.Println(course)
	//fmt.Println(err)
	//fmt.Println(selectCourse)

}
