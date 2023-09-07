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
	UserID  uint   `json:"userID"`
	Cover   string `json:"cover"`
	//	容量
	Capacity uint32 `json:"capacity"`
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
