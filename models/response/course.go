package response

/*
Created by 斑斑砖 on 2023/9/6.
Description：

*/

type CourseResponse struct {
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
	//	上课时间段
	Schedule  string `json:"schedule"`
	IsPreLoad bool   `json:"isPreLoad" gorm:"-"`
	//	开课时间
	StartTime int64 `json:"startTime"`
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
