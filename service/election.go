package service

import (
	"auto-course-web/global"
	"auto-course-web/global/code"
	"auto-course-web/global/election"
	"auto-course-web/global/keys"
	"auto-course-web/initialize/consumer"
	"auto-course-web/models"
	"auto-course-web/models/mq"
	"auto-course-web/models/request"
	"auto-course-web/models/response"
	"auto-course-web/respository"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

/*
Created by 斑斑砖 on 2023/9/14.
Description：

*/

type SelectCourseList struct {
	data *request.SelectCourseReq
}

func ListSelectCourses(data *request.SelectCourseReq) (interface{}, code.Code) {
	return SelectCourseList{
		data: data,
	}.Do()
}
func (list SelectCourseList) Do() (interface{}, code.Code) {
	data, c := list.list()
	if c != code.OK {
		return nil, c
	}
	return data, code.OK
}
func (list SelectCourseList) list() (interface{}, code.Code) {
	result, _ := global.Redis.SMembers(keys.SelectCourseListKey).Result()
	var courses []*response.SelectCourseResponse
	var resp response.List
	count, err := respository.QuerySelectCourse(
		models.Course{},
		&courses,
		list.data.Pager,
		list.data.Title,
		"id in ?  ",
		"",
		list.data.CategoryID,
		result,
	)
	if err != nil {
		//TODO log
		return nil, code.ERROR_DB_OPE
	}
	left, err := global.Redis.HGetAll(keys.SelectCourseKey).Result()
	capacity, err := global.Redis.HGetAll(keys.PreLoadCourseKey).Result()
	for _, course := range courses {
		courseID := strconv.Itoa(int(course.ID))
		l, _ := strconv.Atoi(left[courseID])
		c, _ := strconv.Atoi(capacity[courseID])
		course.Capacity = c
		course.Left = l
	}

	resp.Results = courses
	resp.Count = count
	return resp, code.OK
}

type SelectCourse struct {
	data *request.CreateCourseReq
}

func CreateSelectCourse(userID int, data *request.CreateCourseReq) (interface{}, code.Code) {
	return SelectCourse{
		data: data,
	}.Do(userID)
}
func (create SelectCourse) Do(userID int) (interface{}, code.Code) {
	_, c := create.check()
	if c != code.OK {
		return nil, c
	}
	//	2. 判断用户操作，再次点击就是退课操作
	/*
		key := keys.UserSelectedCourseListKey + strconv.Itoa(userID)
		//是否超过选课上限
		count, err := global.Redis.SCard(key).Result()
		if err != nil {
			//TODO log
			return nil, code.ERROR_DB_OPE
		}
		if count >= keys.SelectCourseMax {
			return nil, code.ERROR_SELECT_COURSE_BEYOND
		}
	*/
	//原子操作
	data, c := create.created(userID)
	if c != code.OK {
		return nil, c
	}
	return data, code.OK
	//非原子操作
	/*
		result, err := global.Redis.SIsMember(key, create.data.ID).Result()
		if err != nil {
			return nil, code.ERROR_DB_OPE
		}
		if result {
			//	取消操作
			_, c := create.cancel(userID)
			if c != code.OK {
				return nil, c
			}
		} else {
			//	选课操作
			_, c := create.create(userID)
			if c != code.OK {
				return nil, c
			}
		}
		return nil, code.OK

	*/
}
func (create SelectCourse) check() (interface{}, code.Code) {
	//	1. 校验课程是否存在
	result, err := global.Redis.SIsMember(keys.SelectCourseListKey, create.data.ID).Result()
	if err != nil {
		return nil, code.ERROR_DB_OPE
	}
	if !result {
		return nil, code.ERROR_COURSE_NOT_EXIST
	}
	return nil, code.OK
}
func (create SelectCourse) create(userID int) (interface{}, code.Code) {
	//1. 通过lua脚本进行原子操作，其中包括，判断是否还剩余课程，剩余就进行创建进行对应课程-1操作，加入到用户的已选集合
	script := redis.NewScript(keys.Lua2CreateCourse)
	courseID := strconv.Itoa(create.data.ID)
	//2. lua脚本
	keyList := []string{keys.SelectCourseKey, courseID, keys.UserSelectedCourseListKey}
	args := []string{strconv.Itoa(userID)}
	val, err := script.Run(global.Redis, keyList, args).Result()
	if err != nil {
		fmt.Println(err)
		return nil, code.ERROR_DB_OPE
	}
	status, _ := val.(int64)
	//3. 是否选课成功
	if status == 0 {
		//课程抢完了
		return nil, code.ERROR_COURSE_NOT_ENOUGH
	}
	//创建丢入消息队列进行创建选课记录等操作。
	return nil, code.OK

}
func (create SelectCourse) cancel(userID int) (interface{}, code.Code) {
	//取消选课操作
	//1. 执行lua脚本
	script := redis.NewScript(keys.Lua2CancelCourse)
	courseID := strconv.Itoa(create.data.ID)
	//2. lua脚本
	keyList := []string{keys.SelectCourseKey, courseID, keys.UserSelectedCourseListKey}
	args := []string{strconv.Itoa(userID)}
	err := script.Run(global.Redis, keyList, args).Err()
	if err != nil {
		return nil, code.ERROR_DB_OPE
	}
	//3. 丢入消息队列进行处理退课等操作
	return nil, code.OK
}
func (create SelectCourse) created(userID int) (interface{}, code.Code) {
	//1. 通过lua脚本进行原子操作，其中包括，判断是否还剩余课程，剩余就进行创建进行对应课程-1操作，加入到用户的已选集合
	script := redis.NewScript(keys.LuaScript2SelectCourse)
	courseID := strconv.Itoa(create.data.ID)
	userKey := keys.UserSelectedCourseListKey + strconv.Itoa(userID)
	keyList := []string{userKey, keys.SelectCourseKey, courseID}
	val, err := script.Run(global.Redis, keyList).Result()
	if err != nil {
		//TODO log
		return nil, code.ERROR_DB_OPE
	}
	res, ok := val.([]interface{})
	if !ok {
		//TODO log
		return nil, code.ERROR_DB_OPE
	}
	status := res[0].(int64)
	switch status {
	case election.CourseFull:
		//课程抢完了
		return nil, code.ERROR_COURSE_NOT_ENOUGH
	case election.CourseSuccess:

		//选课成功
		//1.放入创建记录的消息队列
		//	1.在mysql创建一条记录
		//  2.记录日志
		go consumer.RobConsumer.Product(
			&mq.CourseReq{
				CourseID: create.data.ID,
				UserID:   userID,
			},
		)
		//2. 放入推送的消息队列
		//
	case election.CourseWithdraw:

		//退课成功
		//1.放入退课记录的消息队列
		//  1.在mysql删除对于记录
		//  2.记录日志
		go consumer.DropConsumer.Product(
			&mq.CourseReq{
				CourseID: create.data.ID,
				UserID:   userID,
			},
		)
		////2. 放入推送的消息队列
		//go consumer.PushConsumer.Product(
		//	1,
		//)

	}
	return response.ElectionsResponse{
		CourseID: create.data.ID,
		Capacity: int(res[1].(int64)),
	}, code.OK
}

// ============================================================================ 我选择的课程

type MySelectCourseList struct {
}

func ListMySelectCourses(userID int) (interface{}, code.Code) {
	return MySelectCourseList{}.Do(userID)
}
func (list MySelectCourseList) Do(userID int) (interface{}, code.Code) {
	key := keys.UserSelectedCourseListKey + strconv.Itoa(userID)
	result, err := global.Redis.SMembers(key).Result()
	if err != nil {
		return nil, code.ERROR_DB_OPE
	}
	var courses []*response.SelectCourseResponse
	var resp response.List
	count, err := respository.QuerySelectCourse(
		models.Course{},
		&courses,
		nil,
		"",
		"id in ?  ",
		"",
		0,
		result,
	)
	resp.Results = courses
	resp.Count = count
	return resp, code.OK

}
