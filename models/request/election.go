package request

/*
Created by 斑斑砖 on 2023/9/14.
Description：
	选课的请求参数
*/

// SelectCourseReq 获取选课列表
type SelectCourseReq struct {
	Title      string `json:"title" form:"title" url:"title"   label:"课程名称"`
	CategoryID uint   `json:"category" form:"category" url:"category"  label:"分类ID"`
	Pager      *Pages
}
type CreateCourseReq struct {
	ID int `json:"courseID" form:"courseID" url:"courseID" validate:"required"  label:"课程ID"`
}
