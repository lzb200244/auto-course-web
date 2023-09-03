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
	Name   string `json:"name" label:"昵称"`
	Sex    int    `json:"sex" label:"性别"`
	Email  string `json:"email" label:"邮箱"`
	Desc   string `json:"desc" label:"描述"`
	Avatar string `json:"avatar" label:"头像"`
}
