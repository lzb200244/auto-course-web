package models

import (
	"time"
)

/*
Created by 斑斑砖 on 2023/9/6.
Description：
	课程模型
*/

// 学院

type College struct {
	BaseModel
	Name string `json:"name" gorm:"type:varchar(32);not null;comment:学院名称"`
	Desc string `json:"desc" gorm:"default:'';comment:学院描述"`
}

// 课程分类

type CourseCategory struct {
	BaseModel
	Name string `json:"name" gorm:"type:varchar(32);not null;comment:分类名称"`
	Desc string `json:"desc" gorm:"default:'';comment:分类描述"`
}

// 课程

type Course struct {
	BaseModel
	Title   string `json:"title" gorm:"type:varchar(64);not null;comment:课程名称"`
	Desc    string `json:"desc" gorm:"default:'';comment:课程描述"`
	Code    string `json:"code" gorm:"type:varchar(32);unique;comment:课程编码"`
	Teacher string `json:"teacher" gorm:"type:varchar(32);default:'';comment:讲师"`
	UserID  uint   `json:"userID" gorm:"not null;comment:创建人ID"`
	User    *User  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;comment:创建人"`
	Cover   string `json:"cover" gorm:"default:'';comment:封面"`
	//	学分
	Credit uint32 `json:"credit" gorm:"default:0;comment:学分"`
	//课程分类
	CategoryID uint            `json:"categoryID" gorm:"not null;comment:分类ID"`
	Category   *CourseCategory `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;comment:课程分类"`
	//上课地点
	CollegeID uint `json:"collegeID" gorm:"not null;comment:学院ID"`
	//	上课地点
	College *College `json:"college" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;comment:上课学院"`
	Address string   `json:"address" gorm:"default:'';comment:地点"`
	//	上课时间段
	Schedule string `json:"schedule" gorm:"type:varchar(64);default:'';comment:上课时间段"`
	//	开课时间
	StartTime int64 `json:"startTime" gorm:" not null;comment:开课时间"`
	EndTime   int64 `json:"endTime" gorm:"not null;comment:结束时间"`
}

// CourseSchedule 时刻表
type CourseSchedule struct {
	BaseModel
	Duration string `json:"duration" gorm:"type:varchar(64);"`
}

// UserCourse 用户选课关系表
type UserCourse struct {
	UserID    uint      `json:"userID" gorm:"not null;uniqueIndex:user_course;comment:用户ID"`
	User      *User     `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;comment:用户"`
	CourseID  uint      `json:"courseID" gorm:"not null;uniqueIndex:user_course;comment:课程ID"`
	Course    *Course   `json:"course" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;comment:课程"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`
}
