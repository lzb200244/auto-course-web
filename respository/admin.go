package respository

import (
	"auto-course-web/global"
	"auto-course-web/models"
)

/*
Created by 斑斑砖 on 2023/9/2.
Description：
	权限的curd
*/

func AddAuth(roleID int, permission []int) error {
	var role models.Role
	if err := global.MysqlDB.First(&role, roleID).Error; err != nil {
		return err
	}
	var permissions []models.Permission
	if err := global.MysqlDB.Find(&permissions, "id IN (?)", permission).Error; err != nil {
		return err
	}
	role.Permissions = append(role.Permissions, permissions...)
	if err := global.MysqlDB.Save(&role).Error; err != nil {
		return err
	}
	return nil
}

// DeleteAuth 删除角色的权限
func DeleteAuth(roleID int, permissionID int) error {
	// 查询要删除的 Role
	var role models.Role
	if err := global.MysqlDB.First(&role, roleID).Error; err != nil {
		// TODO log 找不到角色
		return err
	}
	// 查询要删除的 Permission
	var permission models.Permission
	if err := global.MysqlDB.First(&permission, permissionID).Error; err != nil {
		// TODO log 找不到该权限
		return err
	}
	// 从 Role 的 Permissions 中移除要删除的 Permission
	if err := global.MysqlDB.Model(&role).Association("Permissions").Delete(&permission); err != nil {
		//TODO log
		return err
	}
	return nil
}
