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
	Data *request.Auths
}

func NewAuths(data *request.Auths) *Auths {
	return &Auths{
		Data: data,
	}
}
func SetAuth(res *request.Auths) (interface{}, code.Code) {
	return NewAuths(res).Do()
}
func (auths *Auths) Do() (interface{}, code.Code) {
	// 1.角色是否存在

	// 2. 给角色赋予权限
	err := respository.AddAuth(auths.Data.RoleID, auths.Data.Permission)
	if err != nil {
		//TODO	log
		return nil, code.ERROR_ADD_AUTH
	}
	return nil, code.OK
}

// ================================================================== 删除角色的某个权限

type Auth struct {
	Data *request.Auth
}

func NewAuth(data *request.Auth) *Auth {
	return &Auth{
		Data: data,
	}
}
func DelAuth(res *request.Auth) (interface{}, code.Code) {
	return NewAuth(res).Do()
}

func (auth *Auth) Do() (interface{}, code.Code) {
	// 1.角色是否存在

	// 2. 给角色赋予权限
	err := respository.DeleteAuth(auth.Data.RoleID, auth.Data.PermissionID)
	if err != nil {
		//TODO	log
		return nil, code.ERROR_DEL_AUTH
	}
	return nil, code.OK
}

// ================================================================== 创建新的权限

type Permission struct {
	Data *request.Permission
}

func NewPermission(data *request.Permission) *Permission {
	return &Permission{Data: data}
}
func CreatePermission(data *request.Permission) (interface{}, code.Code) {
	return NewPermission(data).Do()
}
func (p Permission) Do() (interface{}, code.Code) {
	err := respository.Create(&models.Permission{Name: p.Data.Name})
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
		Priority:  c.data.Property,
		Meta: models.Meta{
			Title:       c.data.Meta.Title,
			KeepAlive:   c.data.Meta.KeepAlive,
			RequireAuth: c.data.Meta.RequireAuth,
		},
	}
	if _, err := respository.Creat("router", &comp, ""); err != nil {

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
		Priority:  c.data.Property,
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
