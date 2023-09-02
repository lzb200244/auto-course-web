package request

/*
Created by 斑斑砖 on 2023/9/2.
Description：
	进行校验参数
*/

type Auths struct {
	RoleID     int   `json:"roleID" validate:"required" label:"角色ID"`
	Permission []int `json:"permission" validate:"required" label:"权限ID"`
}

type Auth struct {
	RoleID       int `json:"roleID" validate:"required" label:"角色ID"`
	PermissionID int `json:"permissionID" validate:"required" label:"权限ID"`
}

type Permission struct {
	Name string `json:"name" validate:"required" label:"权限名称"`
}
