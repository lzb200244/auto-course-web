package models

import "gorm.io/gorm"

/*
Created by 斑斑砖 on 2023/8/15.
Description：
	rbac认证模型
*/

// Role 用户角色对应关系
type Role struct {
	gorm.Model
	Name        string       `json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
	gorm.Model
	Name string `json:"name"`
}
