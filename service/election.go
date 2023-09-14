package service

import (
	"auto-course-web/global"
	"auto-course-web/global/code"
	"auto-course-web/global/keys"
	"auto-course-web/models"
	"auto-course-web/models/request"
	"auto-course-web/models/response"
	"auto-course-web/respository"
	"strconv"
)

/*
Created by 斑斑砖 on 2023/9/14.
Description：

*/

type SelectCourse struct {
	data *request.SelectCourse
}

func ListSelectCourse(data *request.SelectCourse) (interface{}, code.Code) {
	return SelectCourse{
		data: data,
	}.Do()
}
func (slt SelectCourse) Do() (interface{}, code.Code) {
	data, c := slt.list()
	if c != code.OK {
		return nil, c
	}
	return data, code.OK
}
func (slt SelectCourse) list() (interface{}, code.Code) {
	result, _ := global.Redis.SMembers(keys.SelectCourseListKey).Result()
	var courses []*response.SelectCourseResponse
	var resp response.List

	count, err := respository.QuerySelectCourse(
		models.Course{},
		&courses,
		slt.data.Pager,
		slt.data.Title,
		"id in ?  ",
		"",
		slt.data.CategoryID,
		result,
	)
	if err != nil {
		//TODO log
		return nil, code.ERROR_DB_OPE
	}
	m, err := global.Redis.HGetAll(keys.SelectCourseKey).Result()
	for _, course := range courses {
		course.Capacity = m[strconv.Itoa(int(course.ID))]
	}
	resp.Results = courses
	resp.Count = count
	return resp, code.OK
}
