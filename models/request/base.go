package request

/*
Created by 斑斑砖 on 2023/9/6.
Description：

*/

type Pages struct {
	Page int `json:"page"   label:"页码" validate:"min=0" form:"page" url:"page"`
	Size int `json:"size"    label:"每页数量" validate:"min=0,max=20" form:"size" url:"size"`
}
