package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*
Created by 斑斑砖 on 2023/8/13.
Description：
	用户模型
*/

const (
	PassWordCost = 12 //密码加密难度
)

type User struct {
	gorm.Model
	UserName string `json:"username" gorm:"type:varchar(64);not null;index;comment:用户名称;"`
	Name     string `json:"name" gorm:"type:varchar(32);default:'';comment:姓名"`
	Password string `json:"password" gorm:"not null;comment:用户密码"`
	Email    string `json:"email" gorm:"not null;unique;default:'';comment:邮箱"`
	Avatar   string `json:"avatar" gorm:"default:'';comment:头像"`
	Sex      int    `json:"sex" gorm:"default:0;comment:性别"`
	Desc     string `json:"desc" gorm:"default:'';comment:描述"`
	Role     *Role  `gorm:"foreignKey:RoleID;"`
	RoleID   uint   `json:"role_id" gorm:"not null;comment:用户角色ID"`
}

func (user *User) SetPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
