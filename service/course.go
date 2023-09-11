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
	resp := &response.List{}
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
	resp.Data = courses
	resp.Count = count
	return resp, code.OK
}
func ListCourses(userID int, data *request.Pages) (interface{}, code.Code) {
	return NewListCourse(data).Do(userID)
}

// =================================================================== 获取课程分类

type CourseCategory struct {
	data *request.Pages
}

func NewGetCourseCategory(data *request.Pages) *CourseCategory {
	return &CourseCategory{data: data}
}
func GetCourseCategory(data *request.Pages) (interface{}, code.Code) {
	return NewGetCourseCategory(data).Do()

}
func (category *CourseCategory) Do() (interface{}, code.Code) {
	var categories []*response.CategoryResponse
	resp := &response.List{}
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
	resp.Data = categories
	resp.Count = count
	return resp, code.OK
}

// =================================================================== 教师接收到admin的通知，进行课程预热到缓存

type PreLoadCourse struct {
	data *request.PreloadCourse
}

func NewPublishCourse(data *request.PreloadCourse) *PreLoadCourse {
	return &PreLoadCourse{data: data}
}

func PreLoadCourse2Redis(data *request.PreloadCourse) (interface{}, code.Code) {
	return NewPublishCourse(data).Do()
}
func (course *PreLoadCourse) Do() (interface{}, code.Code) {
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
func (course *PreLoadCourse) check() (interface{}, code.Code) {
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
func (course *PreLoadCourse) load2Redis() (interface{}, code.Code) {
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

func NewCancelPublishCourse(data *request.CancelPublishCourse) *CancelPublishCourse {
	return &CancelPublishCourse{data: data}
}
func CancelCourse2Redis(userID int, data *request.CancelPublishCourse) (interface{}, code.Code) {
	return NewCancelPublishCourse(data).Do(userID)
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

func NewListPublishCourse(data *request.Pages) *ListPublishCourse {
	return &ListPublishCourse{
		data: data,
	}
}
func ListPublishCourses(userID int, data *request.Pages) (interface{}, code.Code) {
	return NewListPublishCourse(data).Do(userID)
}
func (list *ListPublishCourse) Do(userID int) (interface{}, code.Code) {
	result, _ := global.Redis.SMembers(keys.PreLoadCourseListKey).Result()
	var courses []*response.PublishCourseResponse
	resp := &response.List{}
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
	resp.Data = courses
	resp.Count = count
	return resp, code.OK
}
