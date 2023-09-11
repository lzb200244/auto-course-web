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
	Name        string       `json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
	gorm.Model
	Name string `json:"name"`
}

type Meta struct {
	Title       string `json:"title" gorm:"default:''" `
	KeepAlive   bool   `json:"keepAlive" gorm:"default:false"`
	RequireAuth bool   `json:"requireAuth" gorm:"default:true"`
}
type Router struct {
	gorm.Model
	Name      string    `json:"name" gorm:"not null;comment:标题"`
	Path      string    `json:"path"  gorm:"not null;comment:路由" `
	Redirect  string    `json:"redirect"  gorm:"default:'';comment:重定向(针对父路由)"`
	Component string    `json:"component" gorm:"default:'';comment:路由标识/组件的位置"`
	Meta      Meta      `json:"meta" gorm:"type:json;comment:附加属性"`
	Role      uint8     `json:"role" gorm:"default:1;comment:权限控制"`
	Parent    uint      `json:"parent" gorm:"default:1;comment:父级路由ID"`
	Disable   bool      `json:"disable" gorm:"default:false;comment:是否禁用"`
	Children  []*Router `json:"children" gorm:"foreignkey:parent;association_foreignkey:id;" `
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
