package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestRoutes(t *testing.T) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(127.0.0.1:3306)/course?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&Router{},
	)
	//db.Create(
	//	&Router{
	//		Name:      "飒飒撒",
	//		Path:      "飒飒撒",
	//		Redirect:  "飒飒撒",
	//		Component: "飒飒撒",
	//		Parent:    2,
	//		Children:  nil,
	//		Meta: Meta{
	//			Name:       "飒飒撒",
	//			KeepAlive:   false,
	//			RequireAuth: false,
	//		},
	//	})
	d := &Router{}
	db.Preload("Children").Find(d, "id=?", 5)

	// 使用GORM的Create方法将Router对象插入数据库
}
