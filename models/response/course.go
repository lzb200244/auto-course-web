package response

import "auto-course-web/models"

/*
Created by 斑斑砖 on 2023/9/6.
Description：

*/

type CourseResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc" `
	Code    string `json:"code" `
	Teacher string `json:"teacher" `
	Cover   string `json:"cover" `
	//	学分
	Credit uint32 `json:"credit" `
	//课程分类
	CategoryID uint                   `json:"categoryID" `
	Category   *models.CourseCategory `json:"category"`
	CollegeID  uint                   `json:"collegeID" `
	College    *models.College        `json:"college"`
	//上课地点
	Address string `json:"address"`
	//	上课时间段
	Schedule  string `json:"duration"`
	IsPreLoad bool   `json:"isPreLoad" gorm:"-"`

	//	开课时间
	StartTime int64 `json:"startTime" `
	EndTime   int64 `json:"endTime" `
}
type PublishCourseResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Teacher  string `json:"teacher" `
	Cover    string `json:"cover"`
	Capacity string `json:"capacity" gorm:"-"`
	//	学分
	Credit uint32 `json:"credit" `
	//课程分类
	CategoryID uint `json:"categoryID" `
	//	上课时间段
	Schedule string `json:"schedule"`
	//	开课时间
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name" `
	Desc string `json:"desc" `
}
type ScheduleResponse struct {
	ID       uint   `json:"id"`
	Duration string `json:"duration" `
}
