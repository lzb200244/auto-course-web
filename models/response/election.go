package response

/*
Created by 斑斑砖 on 2023/9/15.
Description：
*/

type ElectionsResponse struct {
	CourseID int `json:"courseID"`
	Capacity int `json:"capacity"`
}
type SelectCourseResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc" `

	Teacher  string `json:"teacher" `
	Cover    string `json:"cover"`
	Capacity int    `json:"capacity" gorm:"-"`
	Left     int    `json:"left" gorm:"-"`
	Code     string `json:"code"`
	//	学分
	Credit uint32 `json:"credit" `
	//课程分类
	CategoryID uint `json:"categoryID" `
	//	上课时间段
	Schedule string `json:"schedule"`
	//	开课时间
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime" `
}
