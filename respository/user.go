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
