package models

import (
	"gorm.io/gorm"
)

/*
Created by 斑斑砖 on 2023/9/6.
Description：
	课程模型
*/

// 课程分类

type CourseCategory struct {
	gorm.Model
	Name string `json:"name" gorm:"not null;comment:分类名称"`
	Desc string `json:"desc" gorm:"default:'';comment:分类描述"`
	// 课程 一对多
	//Courses []Course `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Course struct {
	gorm.Model
	Title   string `json:"title" gorm:"not null;comment:课程名称"`
	Desc    string `json:"desc" gorm:"default:'';comment:课程描述"`
	Code    string `json:"code" gorm:"unique;comment:课程编码"`
	Teacher string `json:"teacher" gorm:"default:'';comment:讲师"`
	UserID  uint   `json:"userID" gorm:"not null;comment:创建人ID"`
	User    *User  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;comment:创建人"`
	Cover   string `json:"cover" gorm:"default:'';comment:封面"`
	//	容量
	Capacity uint32 `json:"capacity" gorm:"default:0;comment:容量"`
	//	学分
	Credit uint32 `json:"credit" gorm:"default:0;comment:学分"`
	//课程分类
	CategoryID uint            `json:"categoryID" gorm:"not null;comment:分类ID"`
	Category   *CourseCategory `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;comment:课程分类"`
	//	上课时间段
	Schedule string `json:"schedule" gorm:"default:'';comment:上课时间段"`
	//	开课时间
	StartTime int64 `json:"startTime" gorm:" not null;comment:开课时间"`
	EndTime   int64 `json:"endTime" gorm:"not null;comment:结束时间"`
}
