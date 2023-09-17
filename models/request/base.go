package request

/*
Created by 斑斑砖 on 2023/9/6.
Description：

*/

type Pages struct {
	Page int `json:"page"   label:"页码" validate:"min=0" form:"page" url:"page"`
	Size int `json:"size"    label:"每页数量" validate:"min=0,max=20" form:"size" url:"size"`
}

type Bucket struct {
	Bucket string `json:"bucket" form:"bucket" url:"bucket" validate:"required,oneof=auto-course-files auto-course-cover auto-course-default auto-course-avatar" label:"桶名称"`
}
