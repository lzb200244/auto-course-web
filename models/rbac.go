package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

/*
Created by 斑斑砖 on 2023/8/15.
Description：
	rbac认证模型
*/

// Role 用户角色对应关系
type Role struct {
	gorm.Model
	Name        string       `json:"name" gorm:"type:varchar(32);"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(32)"`
}

type Meta struct {
	Title       string `json:"title" gorm:"default:''" `
	KeepAlive   bool   `json:"keepAlive" gorm:"default:false"`
	RequireAuth bool   `json:"requireAuth" gorm:"default:true"`
	Icon        string `json:"icon" gorm:"default:''"`
}
type Router struct {
	gorm.Model
	//gorm:"type:varchar(32);"
	Name      string  `json:"name" gorm:"type:varchar(32);not null;comment:标题"`
	Path      string  `json:"path"  gorm:"type:varchar(64);not null;comment:路由" `
	Redirect  string  `json:"redirect"  gorm:"type:varchar(64);default:'';comment:重定向(针对父路由)"`
	Component string  `json:"component" gorm:"type:varchar(64);default:'';comment:路由标识/组件的位置"`
	Meta      Meta    `json:"meta" gorm:"type:json;comment:附加属性"`
	Parent    uint    `json:"parent" gorm:"default:1;comment:父级路由ID"`
	Disable   bool    `json:"disable" gorm:"default:false;comment:是否禁用"`
	Role      []*Role `json:"role" gorm:"many2many:role_routers;comment:访问路由的权限"`

	Children []*Router `json:"children" gorm:"foreignkey:parent;association_foreignkey:id;" `
}

func (i *Meta) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	var meta Meta
	err := json.Unmarshal(bytes, &meta)
	if err != nil {
		return err
	}
	*i = meta
	return nil
}

// Value  实现 driver.Valuer 接口，Value 返回 json value
func (i Meta) Value() (driver.Value, error) {
	return json.Marshal(&i)
}
