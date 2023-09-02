package utils

import (
	"fmt"
	"testing"
)

/*
Created by 斑斑砖 on 2023/8/14.
Description：
*/
type User struct {
	UserName string `form:"username" json:"username" validate:"required" label:"用户名"`
	Age      uint8  `form:"age" json:"age" validate:"required,gt=18"`
	Password string `form:"password" json:"password" validate:"required,max=20,min=6" label:"密码"`
	Code     string `form:"code" json:"code" validate:"required,len=6" label:"验证码"`
}

func TestValidateUtil(t *testing.T) {
	user := User{
		Password: "12",
	}
	err := Validator.Validate(&user)
	if err != nil {
		fmt.Println(err.Error())
	}

}
