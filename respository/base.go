package respository

import (
	"go-template/global"
)

/*
Created by 斑斑砖 on 2023/8/14.
Description：
	通用 CRUD
*/

// Create 创建数据(可以创建[单条]数据, 也可[批量]创建)
func Create[T any](data *T) error {
	err := global.MysqlDB.Create(&data).Error

	return err
}

// GetOne [单条]数据查询
func GetOne[T any](data T, query string, args ...any) (T, error) {
	err := global.MysqlDB.Preload("Roles").Where(query, args...).First(&data).Error
	if err != nil {
		return data, err
	}
	return data, err
}

// Update [单行]更新: 传入对应结构体[传递主键用] 和 带有对应更新字段值的[结构体]，结构体不更新零值
func Update[T any](data *T, slt ...string) {
	if len(slt) > 0 {
		global.MysqlDB.Model(&data).Select(slt).Updates(&data)
		return
	}
	err := global.MysqlDB.Model(&data).Updates(&data).Error
	if err != nil {
		panic(err)
	}
}

// UpdatesMap [批量]更新: map 的字段就是要更新的字段 (map 可以更新零值), 通过条件可以实现[单行]更新
func UpdatesMap[T any](data *T, maps map[string]any, query string, args ...any) {
	err := global.MysqlDB.Model(&data).Where(query, args...).Updates(maps).Error
	if err != nil {
		panic(err)
	}
}

// Updates [批量]更新: 结构体的属性就是要更新的字段 (结构体不更新零值), 通过条件可以实现[单行]更新
func Updates[T any](data *T, query string, args ...any) {
	err := global.MysqlDB.Model(&data).Where(query, args...).Updates(&data).Error
	if err != nil {
		panic(err)
	}
}

// List 数据列表
func List[T any](data T, slt, order, query string, args ...any) T {
	global.MysqlDB.Model(&data).Select(slt).Order(order)
	if query != "" {
		global.MysqlDB = global.MysqlDB.Where(query, args...)
	}
	if err := global.MysqlDB.Find(&data).Error; err != nil {
		panic(err)
	}
	return data
}

// Delete [批量]删除数据, 通过条件控制可以删除单条数据
func Delete[T any](data T, query string, args ...any) {
	err := global.MysqlDB.Where(query, args...).Delete(&data).Error
	if err != nil {
		panic(err)
	}
}
