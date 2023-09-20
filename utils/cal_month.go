package utils

import (
	"time"
)

/*
Created by 斑斑砖 on 2023/9/20.
Description：
*/

// CalMonth 计算当前月份的天数
func CalMonth(year int, month time.Month) int {
	// 创建一个指定月份的下一个月份的时间
	nextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	// 减去一天，获取当前月份的最后一天
	lastDay := nextMonth.Add(-time.Hour * 24)
	// 获取当前月份的天数
	daysInMonth := lastDay.Day()
	return daysInMonth
}
