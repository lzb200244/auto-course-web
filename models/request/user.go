package request

import "reflect"

/*
Created by 斑斑砖 on 2023/8/14.
Description：
	进行校验参数
*/

type Register struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required,min=4,max=20" label:"密码"`
	Email    string `json:"email" validate:"required" label:"邮箱"`
}
type Login struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required,min=4,max=20" label:"密码"`
}

func (r Register) IsEmpty() bool {
	return reflect.DeepEqual(r, Register{})
}

func (r Login) IsEmpty() bool {
	return reflect.DeepEqual(r, Login{})
}

type UserInfo struct {
	ID     int    `json:"id" label:"用户ID" validate:"omitempty"`
	Name   string `json:"name" label:"昵称" validate:"omitempty"`
	Sex    int    `json:"sex" label:"性别" validate:"omitempty"`
	Email  string `json:"email" label:"邮箱" validate:"email,omitempty"`
	Desc   string `json:"desc" label:"描述" validate:"omitempty"`
	Avatar string `json:"avatar" label:"头像"  validate:"url,omitempty"`
}
