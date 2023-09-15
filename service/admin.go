package service

import (
	"auto-course-web/global"
	"auto-course-web/global/auth"
	"auto-course-web/global/code"
	"auto-course-web/global/keys"
	"auto-course-web/initialize/consumer"
	"auto-course-web/models"
	"auto-course-web/models/mq"
	"auto-course-web/models/request"
	"auto-course-web/models/response"
	"auto-course-web/respository"
	"auto-course-web/utils/tencent"
	"strconv"
)

/*
Created by 斑斑砖 on 2023/9/2.
Description：

	权限的服务层,赋予权限和删除权限等操作
*/

// ================================================================== 给角色新增权限

type Auths struct {
	Data *request.Auths
}

func AddAuth(data *request.Auths) (interface{}, code.Code) {
	return Auths{Data: data}.Do()
}
func (auths Auths) Do() (interface{}, code.Code) {
	// 1.角色是否存在
	if _, c := auths.check(); c != code.OK {
		return nil, c
	}
	// 2. 给角色赋予权限
	if err := respository.AddAuth(auths.Data.RoleID, auths.Data.Permission); err != nil {
		//TODO	log
		return nil, code.ERROR_ADD_AUTH
	}
	return nil, code.OK
}

func (auths Auths) check() (interface{}, code.Code) {
	exist, err := respository.Exist(&models.Role{}, "id", auths.Data.RoleID)
	if err != nil {
		//TODO log
		return nil, code.ERROR_DB_OPE
	}
	if !exist {
		return nil, code.ERROR_ROLE_NOT_EXIST
	}
	return nil, code.OK
}

// ================================================================== 删除角色的某个权限

type Auth struct {
	Data *request.Auth
}

func DelAuth(data *request.Auth) (interface{}, code.Code) {
	return Auth{
		Data: data,
	}.Do()
}

func (auth Auth) Do() (interface{}, code.Code) {
	// 1.角色是否存在

	// 2. 给角色赋予权限
	if err := respository.DeleteAuth(auth.Data.RoleID, auth.Data.PermissionID); err != nil {
		//TODO	log
		return nil, code.ERROR_DEL_AUTH
	}
	return nil, code.OK
}

// ================================================================== 创建新的权限

type Permission struct {
	Data *request.Permission
}

func CreatePermission(data *request.Permission) (interface{}, code.Code) {
	return Permission{Data: data}.Do()
}
func (p Permission) Do() (interface{}, code.Code) {
	if err := respository.Create(&models.Permission{Name: p.Data.Name}); err != nil {
		return nil, code.ERROR_DB_OPE
	}
	return nil, code.OK
}

// ================================================================== 创建新的页面

type Component struct {
	data *request.Component
}

func CreatePage(data *request.Component) (interface{}, code.Code) {
	return Component{data: data}.Do()
}
func (c Component) Do() (interface{}, code.Code) {
	var roles []*models.Role

	if _, err := respository.List(
		models.Role{}, &roles, nil, "", "id in ?", c.data.Role,
	); err != nil {
		//TODO log
		return nil, code.ERROR_DB_OPE
	}
	comp := models.Router{
		Name:      c.data.Name,
		Component: c.data.Component,
		Path:      c.data.Path,
		Redirect:  c.data.Redirect,
		Parent:    c.data.Parent,
		Role:      roles,
		Meta: models.Meta{
			Title:       c.data.Meta.Title,
			KeepAlive:   c.data.Meta.KeepAlive,
			RequireAuth: c.data.Meta.RequireAuth,
		},
	}
	if _, err := respository.Creat("router", &comp, ""); err != nil {
		//TODO log
		return nil, code.ERROR_DB_OPE
	}
	return nil, code.OK

}

type UpdateComponent struct {
	data *request.Component
}

func UpdatePage(data *request.Component) (interface{}, code.Code) {
	return UpdateComponent{data: data}.Do()
}
func (c UpdateComponent) Do() (interface{}, code.Code) {
	var roles []*models.Role
	global.MysqlDB.Model(models.Router{}).Find(roles, "id in ", c.data.Role)
	comp := models.Router{
		Name:      c.data.Name,
		Component: c.data.Component,
		Path:      c.data.Path,
		Redirect:  c.data.Redirect,
		Parent:    c.data.Parent,
		Disable:   c.data.Disable,
		Role:      roles,
		Meta: models.Meta{
			Title:       c.data.Meta.Title,
			KeepAlive:   c.data.Meta.KeepAlive,
			RequireAuth: c.data.Meta.RequireAuth,
		},
	}
	if err := global.MysqlDB.Updates(&comp).Error; err != nil {
		//TODO log
		return nil, code.ERROR_DB_OPE
	}

	return nil, code.OK

}

// ================================================================= 通知教师预发布课程

type NoticeTeacher struct {
}

func Notice2Teacher() (interface{}, code.Code) {
	return NoticeTeacher{}.Do()
}

func (n NoticeTeacher) Do() (interface{}, code.Code) {
	//发布通知给教师，进行预选发布

	//1. 获取所有教师email
	var emails []string
	global.MysqlDB.
		Model(&models.User{}).
		Select("email").
		Where("role_id = ?", auth.Teacher).Find(&emails)

	//2. 发送邮件（异步）
	//go tencent.SendEmail("课程通预发布通知", "亲爱的老师：您好,您的课程正在进行预发布。", emails)
	consumer.EmailConsumer.Product(
		mq.EmailReq{
			Title:   "课程通预发布通知",
			Message: "亲爱的老师：您好,您的课程正在进行预发布。",
			Users:   emails,
		},
	)

	//3. 开启预发布通道 ,不存在时才进行创建,存在了只进行预先通知,不进行再次开启通道
	global.Redis.SetNX(keys.IsPreLoadedKey, 1, keys.PreSelectedDurationKey)

	return nil, code.OK
}

// ================================================================= 通知学生加入选课阶段

type NoticeStudent struct {
}

func Notice2Student() (interface{}, code.Code) {
	return NoticeStudent{}.Do()
}
func (notice NoticeStudent) Do() (interface{}, code.Code) {
	//发布通知给学生，进行选课，开启选课

	//1. 获取所有教师email
	var emails []string
	global.MysqlDB.
		Model(&models.User{}).
		Select("email").
		Where("role_id = ?", auth.Student).Find(&emails)
	//2. 发送邮件（异步）TODO 消息队列进行处理
	go tencent.SendEmail("课程通选课通知", "xxx学生：您好,您的课程正在进行选课阶段。", emails)
	//3. 开启预发布通道 ,不存在时才进行创建,存在了只进行预先通知,不进行再次开启通道
	global.Redis.SetNX(keys.IsSelectCourseKey, 1, keys.SelectCourseDurationKey)
	//4. 迁移：待选区课程-》选课区
	if _, c := notice.load2Redis(); c != code.OK {
		return nil, c
	}
	return nil, code.OK
}
func (notice NoticeStudent) load2Redis() (interface{}, code.Code) {
	pipeline := global.Redis.Pipeline()
	courseSet := global.Redis.SMembers(keys.PreLoadCourseListKey)
	for _, course := range courseSet.Val() {
		pipeline.SAdd(keys.SelectCourseListKey, course)
	}
	data := global.Redis.HGetAll(keys.PreLoadCourseKey)
	for k, v := range data.Val() {
		pipeline.HSet(keys.SelectCourseKey, k, v)
	}
	//pipeline.Expire(keys.SelectCourseKey, keys.SelectCourseDurationKey)
	//pipeline.Expire(keys.SelectCourseListKey, keys.SelectCourseDurationKey)
	if _, err := pipeline.Exec(); err != nil {
		return nil, code.ERROR_DB_OPE
	}
	return nil, code.OK
}

// ================================================================= 获取所有的预选课程

type PreloadCourse struct {
	data *request.Pages
}

func ListPreloadCourse(data *request.Pages) (interface{}, code.Code) {
	return PreloadCourse{data: data}.Do()
}
func (list PreloadCourse) Do() (interface{}, code.Code) {
	result, _ := global.Redis.SMembers(keys.PreLoadCourseListKey).Result()
	var courses []*response.PublishCourseResponse
	var resp response.List
	count, err := respository.List(
		models.Course{},
		&courses,
		list.data,
		"start_time",
		"id in ?", result,
	)
	if err != nil {
		//TODO log
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

// ================================================================== 创建课程分类

type Category struct {
	data *request.Category
}

func CreateCategory(data *request.Category) (interface{}, code.Code) {
	return Category{data: data}.Do()
}
func (category Category) Do() (interface{}, code.Code) {
	if _, c := category.check(); c != code.OK {
		return nil, c
	}
	if _, c := category.create(); c != code.OK {
		return nil, c
	}
	return nil, code.OK
}

// 校验分类名称是否存在
func (category Category) check() (interface{}, code.Code) {
	exist, err := respository.Exist(models.CourseCategory{}, "name", category.data.Name)
	if err != nil {
		//TODO log
		return nil, code.ERROR_DB_OPE
	}
	if exist {
		return nil, code.ERROR_COURSE_CATEGORY_EXIST
	}
	return nil, code.OK
}

// 创建分类
func (category Category) create() (interface{}, code.Code) {
	if err := respository.Create(
		&models.CourseCategory{
			Name: category.data.Name, Desc: category.data.Desc},
	); err != nil {
		//TODO log
		return nil, code.ERROR_DB_OPE
	}
	return nil, code.OK
}
