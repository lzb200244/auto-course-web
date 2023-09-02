package respository

import (
	"go-template/global"
	"go-template/models"
)

/*
Created by 斑斑砖 on 2023/8/15.
Description：
	用户
*/

// AddUserAuthority 给用户赋予角色
func AddUserAuthority(user models.User, roleID []int) error {
	var roles models.Role
	global.MysqlDB.Find(&roles, "id in ?", roleID)
	user.Roles = append(user.Roles, roles)
	err := global.MysqlDB.Save(&user).Error
	return err
}

// GetUserInfo 获取用户信息、角色和权限
func GetUserInfo[T any](data T, query string, args ...any) (T, error) {
	err := global.MysqlDB.
		Preload("Roles").
		Where(query, args...).
		First(&data).Error
	if err != nil {
		return data, err
	}
	return data, err
}

// GetPermission 获取用户角色信息
func GetPermission(id int) []int {
	var permission []int
	var role models.Role
	global.MysqlDB.Preload("Permissions").Find(&role, "id=?", id)
	for _, item := range role.Permissions {
		permission = append(permission, int(item.ID))
	}
	return permission

}
