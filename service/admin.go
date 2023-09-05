package service

import (
	"auto-course-web/global"
	"auto-course-web/global/code"
	"auto-course-web/models"
	"auto-course-web/models/request"
	"auto-course-web/respository"
)

/*
Created by 斑斑砖 on 2023/9/2.
Description：

	权限的服务层,赋予权限和删除权限等操作
*/
// ================================================================== 给角色新增权限

type Auths struct {
	RoleID     int   `json:"roleID" `
	Permission []int `json:"permission" `
}

func NewAuths(roleID int, permission []int) *Auths {
	return &Auths{RoleID: roleID, Permission: permission}
}
func SetAuth(roleID int, permission []int) (interface{}, code.Code) {
	return NewAuths(roleID, permission).Do()
}
func (a *Auths) Do() (interface{}, code.Code) {
	// 1.角色是否存在

	// 2. 给角色赋予权限
	err := respository.AddAuth(a.RoleID, a.Permission)
	if err != nil {
		//TODO	log
		return nil, code.ERROR_ADD_AUTH
	}
	return nil, code.OK
}

// ================================================================== 删除角色的某个权限

type Auth struct {
	RoleID       int `json:"roleID" `
	PermissionID int `json:"permissionID"`
}

func NewAuth(roleID int, permissionID int) *Auth {
	return &Auth{RoleID: roleID, PermissionID: permissionID}
}
func DelAuth(roleID int, permissionID int) (interface{}, code.Code) {
	return NewAuth(roleID, permissionID).Do()
}

func (a *Auth) Do() (interface{}, code.Code) {
	// 1.角色是否存在

	// 2. 给角色赋予权限
	err := respository.DeleteAuth(a.RoleID, a.PermissionID)
	if err != nil {
		//TODO	log
		return nil, code.ERROR_DEL_AUTH
	}
	return nil, code.OK
}

// ================================================================== 创建新的权限

type Permission struct {
	Name string `json:"name" `
}

func NewPermission(name string) *Permission {
	return &Permission{Name: name}
}
func CreatePermission(name string) (interface{}, code.Code) {
	return NewPermission(name).Do(name)
}
func (p Permission) Do(name string) (interface{}, code.Code) {
	err := respository.Create(&models.Permission{Name: name})
	if err != nil {
		return nil, code.ERROR_CREATE_PERMISSION
	}
	return nil, code.OK
}

// ================================================================== 创建新的页面

type Component struct {
	data *request.Component
}

func NewComponent(data *request.Component) *Component {
	return &Component{
		data: data,
	}
}
func CreatePage(data *request.Component) (interface{}, code.Code) {
	return NewComponent(data).Do()
}
func (c Component) Do() (interface{}, code.Code) {
	comp := models.Router{
		Name:      c.data.Name,
		Component: c.data.Component,
		Path:      c.data.Path,
		Redirect:  c.data.Redirect,
		Parent:    c.data.Parent,
		Limit:     c.data.Limit,
		Meta: models.Meta{
			Title:       c.data.Meta.Title,
			KeepAlive:   c.data.Meta.KeepAlive,
			RequireAuth: c.data.Meta.RequireAuth,
		},
	}
	if err := respository.Create(&comp); err != nil {

		return nil, code.ERROR_DB_OPE
	}
	return nil, code.OK

}

type UpdateComponent struct {
	data *request.Component
}

func UpdatePage(data *request.Component) (interface{}, code.Code) {
	return NewComponent(data).Do()
}
func (c UpdateComponent) Do() (interface{}, code.Code) {
	comp := models.Router{
		Name:      c.data.Name,
		Component: c.data.Component,
		Path:      c.data.Path,
		Redirect:  c.data.Redirect,
		Parent:    c.data.Parent,
		Disable:   c.data.Disable,
		Limit:     c.data.Limit,
		Meta: models.Meta{
			Title:       c.data.Meta.Title,
			KeepAlive:   c.data.Meta.KeepAlive,
			RequireAuth: c.data.Meta.RequireAuth,
		},
	}
	if err := global.MysqlDB.Updates(&comp).Error; err != nil {
		return nil, code.ERROR_DB_OPE
	}

	return nil, code.OK

}
