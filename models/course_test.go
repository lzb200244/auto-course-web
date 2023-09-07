package models

import (
	"auto-course-web/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
)

/*
Created by 斑斑砖 on 2023/9/6.
Description：
*/
func TestCourse(t *testing.T) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(127.0.0.1:3306)/course?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: global.Config.Mysql.Singular, // 表明不加s
		},
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&CourseCategory{},
		&Course{},
	)

	db.Create(&CourseCategory{
		Name: "前端",
		Desc: "前端课程",
	})
	//
	db.Create(&Course{
		Title:      "前端课程",
		Desc:       "前端课程",
		UserID:     1,
		CategoryID: 1,
		Capacity:   50,
		Credit:     2,
		Code:       "FT123",
		Schedule:   "8:00-9:00",
		StartTime:  "2023-01-01",
		EndTime:    "2023-04-02",
	})
	courses := &Course{}
	db.Preload("Category").Preload("User").First(courses, 1)
	fmt.Println(courses.Category)
	fmt.Println(courses.User)

}
