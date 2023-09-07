package service

import (
	"auto-course-web/global/code"
	"auto-course-web/models"
	"auto-course-web/models/request"
	"auto-course-web/models/response"
	"auto-course-web/respository"
)

/*
Created by 斑斑砖 on 2023/9/6.
Description：
*/

type Course struct {
	data *request.Course
}

func NewCourse(data *request.Course) *Course {
	return &Course{data: data}
}

func CreateCourse(userID int, data *request.Course) (interface{}, code.Code) {
	return NewCourse(data).Do(userID)
}
func (course *Course) Do(userID int) (interface{}, code.Code) {
	if _, c := course.check(); c != code.OK {
		return nil, c
	}
	if _, c := course.create(userID); c != code.OK {
		return nil, c
	}
	return nil, code.OK

}

// 校验课程分类是否存
func (course *Course) check() (interface{}, code.Code) {
	if ok, _ := respository.Exist(&models.CourseCategory{}, "id", course.data.CategoryID); !ok {
		return nil, code.ERROR_COURSE_CATEGORY_NOT_EXIST
	}
	return nil, code.OK
}

// 创建到数据库
func (course *Course) create(userID int) (interface{}, code.Code) {
	course.data.UserID = userID
	if _, err := respository.Creat("course", course.data, ""); err != nil {
		return nil, code.ERROR_DB_OPE
	}
	return course.data, code.OK
}

// 进行预热到redis
func (course *Course) load2Redis() (interface{}, code.Code) {
	return nil, code.OK

}

// =================================================================== 返回教师创建的课程列表

type ListCourse struct {
	data *request.Pages
}

func NewListCourse(data *request.Pages) *ListCourse {
	return &ListCourse{
		data: data,
	}
}
func (list *ListCourse) Do(userID int) (interface{}, code.Code) {
	var courses []*response.CourseResponse
	respository.List(
		models.Course{},
		&courses,
		"start_time",
		"user_id=?",
		userID,
	)
	return courses, code.OK
}
func ListCourses(userID int, data *request.Pages) (interface{}, code.Code) {
	return NewListCourse(data).Do(userID)
}
