package service

import (
	"auto-course-web/global/auth"
	"auto-course-web/global/code"
	"auto-course-web/models"
	"auto-course-web/models/request"
	"auto-course-web/models/response"
	"auto-course-web/respository"
	"auto-course-web/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"sync"
)

/*
Created by 斑斑砖 on 2023/8/14.
Description：
	用户
*/

// ================================================================= 用户注册

type UserRegister struct {
	Username string
	Password string
	Email    string
}

func NewUser(username string, password string, email string) *UserRegister {
	return &UserRegister{
		Username: username, Password: password, Email: email,
	}
}
func Register(username string, password string, email string) (interface{}, code.Code) {
	return NewUser(username, password, email).Do()
}
func (r UserRegister) Do() (interface{}, code.Code) {
	_, c := r.checkExists()
	if c != code.OK {
		return nil, c
	}
	_, c = r.create()
	if c != code.OK {
		return nil, c
	}

	return nil, code.OK
}

// 用户是否存在
func (r UserRegister) checkExists() (interface{}, code.Code) {
	_, usernameErr := respository.GetOne(&models.User{}, "user_name", r.Username)
	if usernameErr != nil {
		if errors.Is(usernameErr, gorm.ErrRecordNotFound) {
			// 用户名可用
			// 继续检查邮箱是否已存在
			_, emailErr := respository.GetOne(&models.User{}, "email", r.Email)
			if emailErr != nil {
				if errors.Is(emailErr, gorm.ErrRecordNotFound) {
					// 邮箱可用
					return nil, code.OK
				}
			}
			// 邮箱已存在
			return nil, code.ERROR_EMAIL_EXIST
		}
		// 处理其他用户名查询异常
		return nil, code.ERROR_DB_OPE
	}
	// 用户名已存在
	return nil, code.ERROR_USER_NAME_USED
}

// 创建用户
func (r UserRegister) create() (interface{}, code.Code) {
	user := models.User{UserName: r.Username, Password: r.Password, Email: r.Email}
	user.SetPassword() // 加密

	err := respository.Create(&user)

	//给用户赋予权限
	respository.AddUserAuthority(user, []int{auth.Student})
	if err != nil {
		fmt.Println(err, 2222)
		//TODO 记录日志
		return nil, code.ERROR_DB_OPE
	}
	return nil, code.OK

}

// ================================================================= 用户登录

type UserLogin struct {
	Username string
	Password string
}

func NewUserLogin(username string, password string) *UserLogin {
	return &UserLogin{Username: username, Password: password}
}
func Login(username string, password string) (interface{}, code.Code) {
	return NewUserLogin(username, password).Do()
}
func (r UserLogin) Do() (interface{}, code.Code) {
	data, c := r.checkAndSign()
	if c != code.OK {
		return nil, c
	}
	return data, code.OK
}
func (r UserLogin) checkAndSign() (interface{}, code.Code) {

	userObj, err := respository.GetUserInfo(&models.User{}, "user_name", r.Username)
	if err != nil {
		//不存在该用户
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, code.ERROR_USER_NOT_EXIST
		}
		//数据库异常错误
		return nil, code.ERROR_DB_OPE
	}
	//密码错误
	if ok := userObj.CheckPassword(r.Password); !ok {
		return nil, code.ERROR_PASSWORD_WRONG
	}
	//校验通过生成Token
	token, err := utils.GenerateToken(userObj.ID, userObj.UserName, userObj.Email, int(userObj.Roles[0].ID))
	if err != nil {
		//签发token失败
		return nil, code.ERROR_TOKEN_CREATE
	}
	roleObj := userObj.Roles[0]
	permission := respository.GetPermission(int(roleObj.ID))

	return response.NewUserResponse(
		userObj.ID,
		userObj.UserName,
		userObj.Name,
		userObj.Email,
		userObj.Desc,
		userObj.Avatar,
		token,
		userObj.Sex,
		[]string{roleObj.Name},
		permission,
	), code.OK

}

// ================================================================= 获取用户信息

type UserInfo struct {
}

func NewUserInfo() *UserInfo {
	return &UserInfo{}
}
func GetUserInfo(userID, roleID int) (interface{}, code.Code) {
	return NewUserInfo().Do(userID, roleID)
}
func (u *UserInfo) Do(userID, roleID int) (interface{}, code.Code) {
	//1. 判断用户是否存在
	userObj, c := u.GetUserObj(userID, roleID)
	if c != code.OK {

		return nil, c
	}

	return userObj, code.OK
}
func (u *UserInfo) GetUserObj(userID, roleID int) (interface{}, code.Code) {
	var wg sync.WaitGroup
	var permission []int
	var userObj models.User
	var err error
	wg.Add(2)
	//1. 获取用户信息
	go func() {
		defer wg.Done()
		userObj, err = respository.GetUserInfo(models.User{}, "id", userID)

	}()
	//2. 获取用户角色与权限
	go func() {
		defer wg.Done()
		permission = respository.GetPermission(roleID)

	}()
	wg.Wait()
	if err != nil {
		//不存在该用户
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, code.ERROR_USER_NOT_EXIST
		}
		//数据库异常错误
		return nil, code.ERROR_DB_OPE
	}
	return response.NewUserResponse(
		userObj.ID,
		userObj.UserName,
		userObj.Name,
		userObj.Email,
		userObj.Desc,
		userObj.Avatar,
		"",
		userObj.Sex,
		[]string{auth.GetAuthorityName(roleID)},
		permission,
	), code.OK

}

// ================================================================= 修改用户信息

type UserInfoUpdate struct {
	data *request.UserInfo
}

func NewUserInfoUpdate(req *request.UserInfo) *UserInfoUpdate {
	return &UserInfoUpdate{
		data: req,
	}
}
func UpdateInfo(userID uint, req *request.UserInfo) (interface{}, code.Code) {
	return NewUserInfoUpdate(req).Do(userID)
}

// Do 更新用户信息
func (u *UserInfoUpdate) Do(userID uint) (interface{}, code.Code) {
	//1. 判断新的邮箱是否存在
	if u.data.Email != "" {
		ok, _ := respository.Exist(&models.User{}, "email=? and id !=?", u.data.Email, userID)
		if ok {
			return nil, code.ERROR_EMAIL_EXIST
		}
	}
	//2. 更新
	if err := respository.Updates(&models.User{}, &u.data, "id=?", userID); err != nil {
		return nil, code.ERROR_UPDATE_USER
	}
	return &u.data, code.OK
}
