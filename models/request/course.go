package request

/*
Created by 斑斑砖 on 2023/9/6.
Description：
	参数校验
*/

type Course struct {
	ID      int    `json:"id"  label:"课程ID"`
	Title   string `json:"title" validate:"required"  label:"课程名称"`
	Desc    string `json:"desc"  validate:"required"  label:"课程描述"`
	Teacher string `json:"teacher" validate:"required"  label:"讲师"`
	Cover   string `json:"cover"  label:"封面" `
	UserID  int    `json:"userID"  label:"用户ID"`
	Code    string `json:"code"  label:"课程代码"`
	//	学分
	Credit uint32 `json:"credit" validate:"required,min=0"  label:"学分"`
	//课程分类
	CategoryID uint `json:"categoryID" validate:"required"  label:"分类ID"`
	//	上课时间段
	Schedule string `json:"schedule"`
	Duration int    `json:"duration" validate:"required" label:"时间段" gorm:"-"`
	//	开课时间
	StartTime int64 `json:"startTime" validate:"required"  label:"开课时间"`
	EndTime   int64 `json:"endTime" validate:"required"  label:"结束时间"`
}

type UpdateCourse struct {
	ID         int    `json:"id"  label:"课程ID"`
	Title      string `json:"title"   label:"课程名称"`
	Desc       string `json:"desc"    label:"课程描述"`
	Cover      string `json:"cover"  label:"封面" `
	Schedule   string `json:"schedule"   label:"上课时间段"`
	Duration   int    `json:"duration" validate:"required" label:"时间段" gorm:"-"`
	CategoryID uint   `json:"categoryID" validate:"required"  label:"分类ID"`
}

// PreloadCourse 教师接收到管理员的通知，加载课程到缓存
type PreloadCourse struct {
	CourseID int `json:"courseID" form:"courseID" url:"courseID" validate:"required,min=0"  label:"课程ID"`
	//	容量
	Capacity uint32 `json:"capacity" form:"capacity" url:"capacity" validate:"required,min=0"  label:"容量"`
}

type CancelPublishCourse struct {
	CourseID int `json:"courseID" form:"courseID" url:"courseID" validate:"required,min=0"  label:"课程ID"`
}
