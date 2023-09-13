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
	Created by 斑斑砖 on 2023/9/6.
	Description：
*/

// =================================================================== 创建课程,并未进行预热到缓存

type Course struct {
	data *request.Course
}

func CreateCourse(userID int, data *request.Course) (interface{}, code.Code) {
	return Course{data: data}.Do(userID)
}
func (course Course) Do(userID int) (interface{}, code.Code) {
	if _, c := course.check(); c != code.OK {
		return nil, c
	}
	if _, c := course.create(userID); c != code.OK {
		return nil, c
	}
	return nil, code.OK

}

// 校验课程分类是否存
func (course Course) check() (interface{}, code.Code) {
	if ok, _ := respository.Exist(&models.CourseCategory{}, "id", course.data.CategoryID); !ok {
		return nil, code.ERROR_COURSE_CATEGORY_NOT_EXIST
	}
	return nil, code.OK
}

// 创建到数据库
func (course Course) create(userID int) (interface{}, code.Code) {
	course.data.UserID = userID
	if _, err := respository.Creat("course", course.data, ""); err != nil {
		return nil, code.ERROR_DB_OPE
	}
	return course.data, code.OK
}

// =================================================================== 返回教师创建的课程列表

type Courses struct {
	data *request.Pages
}

func (list Courses) Do(userID int) (interface{}, code.Code) {

	var courses []*response.CourseResponse
	var resp response.List
	count, err := respository.List(
		models.Course{},
		&courses,
		list.data,
		"start_time",
		"user_id=?",
		userID,
	)
	if err != nil {
		return nil, code.ERROR_DB_OPE
	}
	for _, course := range courses {
		result, err := global.Redis.SIsMember(keys.PreLoadCourseListKey, strconv.Itoa(int(course.ID))).Result()
		if err != nil {
			return nil, code.ERROR_DB_OPE
		}
		if result {
			course.IsPreLoad = true
		}
	}
	resp.Results = courses
	resp.Count = count
	return resp, code.OK
}
func ListCourse(userID int, data *request.Pages) (interface{}, code.Code) {
	return Courses{
		data: data,
	}.Do(userID)
}

// =================================================================== 更新课程

type UpdateCourses struct {
	data *request.UpdateCourse
}

func UpdateCourse(userID int, data *request.UpdateCourse) (interface{}, code.Code) {
	return UpdateCourses{data: data}.Do(userID)
}
func (course UpdateCourses) Do(userID int) (interface{}, code.Code) {
	if _, c := course.check(userID); c != code.OK {
		return nil, c
	}
	if _, c := course.update(); c != code.OK {
		return nil, c
	}
	return nil, code.OK
}
func (course UpdateCourses) check(userID int) (interface{}, code.Code) {
	//是否存在该课程，且是否是自己创建的
	exist, err := respository.Exist(&models.Course{}, "id=? and user_id=?", course.data.ID, userID)
	if err != nil {
		return nil, 0
	}
	if !exist {
		return nil, code.ERROR_COURSE_NOT_EXIST
	}
	return nil, code.OK
}
func (course UpdateCourses) update() (interface{}, code.Code) {
	if err := respository.Updates(
		&models.Course{}, &course.data, "id=?", course.data.ID); err != nil {
		return nil, code.ERROR_UPDATE_USER
	}
	return nil, code.OK
}

// =================================================================== 获取课程分类

type CourseCategory struct {
	data *request.Pages
}

func ListCourseCategory(data *request.Pages) (interface{}, code.Code) {
	return CourseCategory{data: data}.Do()

}
func (category CourseCategory) Do() (interface{}, code.Code) {
	var categories []*response.CategoryResponse
	var resp response.List
	count, err := respository.List(
		models.CourseCategory{},
		&categories,
		category.data,
		"id",
		"",
	)
	if err != nil {
		return nil, code.ERROR_DB_OPE
	}
	resp.Results = categories
	resp.Count = count
	return resp, code.OK
}

// =================================================================== 教师接收到admin的通知，进行课程预热到缓存

type PreLoadCourse struct {
	data *request.PreloadCourse
}

func PreLoadCourse2Redis(data *request.PreloadCourse) (interface{}, code.Code) {
	return PreLoadCourse{data: data}.Do()
}
func (course PreLoadCourse) Do() (interface{}, code.Code) {
	if _, c := course.check(); c != code.OK {
		return nil, c
	}
	//加载到redis
	if _, c := course.load2Redis(); c != code.OK {
		return nil, c
	}
	return nil, code.OK
}

// 校验课程是否存在/是否处于预选课状态
func (course PreLoadCourse) check() (interface{}, code.Code) {
	//1. 是否处于预选课状态
	if result, _ := global.Redis.Exists(keys.IsPreLoadedKey).Result(); result == 0 {
		//不处于预选课状态
		return nil, code.ERROR_PRELOAD_COURSE_NOT_OPEN
	}

	//2. 是否已经发布过
	result, err := global.Redis.SIsMember(keys.PreLoadCourseListKey, course.data.CourseID).Result()
	if err != nil {
		return nil, 0
	}
	if result {
		return nil, code.ERROR_COURSE_ALREADY
	}
	if ok, _ := respository.Exist(&models.Course{}, "id", course.data.CourseID); !ok {
		return nil, code.ERROR_COURSE_NOT_EXIST
	}
	return nil, code.OK
}

// 预热到redis
func (course PreLoadCourse) load2Redis() (interface{}, code.Code) {
	err := global.Redis.HSet(
		keys.PreLoadCourseKey,
		strconv.Itoa(course.data.CourseID),
		course.data.Capacity,
	).Err()
	//放入已经预热de课程的集合
	err = global.Redis.SAdd(keys.PreLoadCourseListKey, course.data.CourseID).Err()
	if err != nil {
		return nil, code.ERROR_DB_OPE
	}
	return nil, code.OK
}

// =================================================================== 取消发布课程

type CancelPublishCourse struct {
	data *request.CancelPublishCourse
}

func CancelCourse2Redis(userID int, data *request.CancelPublishCourse) (interface{}, code.Code) {
	return CancelPublishCourse{data: data}.Do(userID)
}
func (course CancelPublishCourse) Do(userID int) (interface{}, code.Code) {
	//是否是我的创建的课程

	if _, c := course.check(userID); c != code.OK {
		return nil, c
	}
	if _, c := course.delete2Redis(); c != code.OK {
		return nil, c
	}

	return nil, code.OK
}
func (course CancelPublishCourse) check(userID int) (interface{}, code.Code) {
	//1. 是否处于预选课状态
	if result, _ := global.Redis.Exists(keys.IsPreLoadedKey).Result(); result == 0 {
		//不处于预选课状态
		return nil, code.ERROR_PRELOAD_COURSE_CLOSE
	}
	exist, err := respository.Exist(models.Course{}, "id=? and user_id=?", course.data.CourseID, userID)
	if err != nil {
		return nil, code.ERROR_DB_OPE
	}
	if !exist {
		return nil, code.ERROR_COURSE_NOT_EXIST
	}
	return nil, code.OK

}
func (course CancelPublishCourse) delete2Redis() (interface{}, code.Code) {
	if err := global.Redis.HDel(keys.PreLoadCourseKey, strconv.Itoa(course.data.CourseID)).Err(); err != nil {
		return nil, code.ERROR_COURSE_DELETE
	}
	if err := global.Redis.SRem(keys.PreLoadCourseListKey, course.data.CourseID).Err(); err != nil {
		return nil, code.ERROR_COURSE_DELETE
	}
	return nil, code.OK
}

// =================================================================== 处于发布课程

type ListPublishCourse struct {
	data *request.Pages
}

func ListPublishCourses(userID int, data *request.Pages) (interface{}, code.Code) {
	return ListPublishCourse{
		data: data,
	}.Do(userID)
}
func (list ListPublishCourse) Do(userID int) (interface{}, code.Code) {
	result, _ := global.Redis.SMembers(keys.PreLoadCourseListKey).Result()
	var courses []*response.PublishCourseResponse
	var resp response.List
	count, err := respository.List(
		models.Course{},
		&courses,
		list.data,
		"start_time",
		"user_id=? and id in ?", userID, result,
	)
	if err != nil {
		return nil, code.ERROR_DB_OPE
	}
	m, err := global.Redis.HGetAll(keys.PreLoadCourseKey).Result()
	for _, course := range courses {
		course.Capacity = m[strconv.Itoa(int(course.ID))]
	}
	resp.Results = courses
	resp.Count = count
	return resp, code.OK
}
