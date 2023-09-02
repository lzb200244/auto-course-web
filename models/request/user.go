package request

import "reflect"

/*
Created by 斑斑砖 on 2023/8/14.
Description：
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
