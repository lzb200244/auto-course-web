package service

import (
	"auto-course-web/global"
	"auto-course-web/global/auth"
	"auto-course-web/global/code"
	"auto-course-web/global/keys"
	"auto-course-web/initialize/consumer"
	"auto-course-web/models"
	"auto-course-web/models/mq"
	"auto-course-web/models/request"
	"auto-course-web/models/response"
	"auto-course-web/respository"
	"auto-course-web/utils"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"math"
	"sync"
	"time"
)

/*
Created by 斑斑砖 on 2023/8/14.
Description：
	用户
*/

// ================================================================= 用户注册

type UserRegister struct {
	data *request.Register
}

func Register(data *request.Register) (interface{}, code.Code) {
	return UserRegister{
		data: data,
	}.Do()
}
func (r UserRegister) Do() (interface{}, code.Code) {
	if _, c := r.checkCode(); c != code.OK {
		return nil, c
	}
	if _, c := r.checkExists(); c != code.OK {
		return nil, c
	}
	if _, c := r.create(); c != code.OK {
		return nil, c
	}
	return nil, code.OK
}

// 校验邮箱操作
func (r UserRegister) checkCode() (interface{}, code.Code) {
	result, _ := global.Redis.Get(keys.CodeKey + r.data.Email).Result()

	if result != r.data.Code {
		return nil, code.ERROR_VERIFICATION_CODE
	}
	return nil, code.OK
}

// 用户是否存在
func (r UserRegister) checkExists() (interface{}, code.Code) {
	//先往布隆查看，如果布隆不存在就是一定不存在，存在也有可能不存在
	exists, err := global.Bloom.Exists(keys.UserBloomKey, r.data.Username)
	if err != nil {
		return nil, code.ERROR_DB_OPE
	}
	if !exists { // 不存在就一定不存在
		return nil, code.OK
	}
	//可能存在，往库检查

	exist, err := respository.Exist(&models.User{}, "user_name=?", r.data.Username)
	if err != nil {
		fmt.Println(err)
		return nil, code.ERROR_DB_OPE
	}
	if exist {
		return nil, code.ERROR_USER_NAME_EXIST
	}
	return nil, code.OK

}

// 创建用户
func (r UserRegister) create() (interface{}, code.Code) {
	user := models.User{
		UserName: r.data.Username,
		Password: r.data.Password,
		Email:    r.data.Email,
		RoleID:   uint(auth.Student),
		Name:     r.data.Username,
		Desc:     utils.GenerateDesc(),
		Avatar:   utils.GenerateAvatar(r.data.Username),
	}
	if err := user.SetPassword(); err != nil {
		fmt.Println(err)
		return nil, code.ERROR_DB_OPE
	}
	//创建操作
	if err := respository.Create(&user); err != nil {
		fmt.Println(err)
		return nil, code.ERROR_DB_OPE
	}
	//用户名放入布隆过滤器
	if _, err := global.Bloom.Add(keys.UserBloomKey, r.data.Username); err != nil {

		return nil, code.ERROR_DB_OPE
	}
	//邮箱放入布隆过滤器
	if _, err := global.Bloom.Add(keys.EmailBloomKey, r.data.Email); err != nil {
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
	//先往bloom查看
	/*
		exists, err := global.Bloom.Exists(keys.UserBloomKey, r.Username)
		if err != nil {
			fmt.Println(err)
			return nil, code.ERROR_DB_OPE
		}
		if !exists {
			return nil, code.ERROR_USER_NOT_EXIST
		}

	*/
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

	token, err := utils.GenerateToken(userObj.ID, userObj.UserName, userObj.Email, int(userObj.Role.ID))
	if err != nil {
		//签发token失败
		return nil, code.ERROR_TOKEN_CREATE
	}
	roleObj := userObj.Role
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
		global.Logger.Warn(err.Error())
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
		[]string{auth.GetAuthorityName(auth.Auth(roleID))},
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
		//布隆查看
		exists, err := global.Bloom.Exists(keys.EmailBloomKey, u.data.Email)
		if err != nil {
			return nil, code.ERROR_DB_OPE
		}
		//存在,是否是自己的
		if exists {
			ok, _ := respository.Exist(&models.User{}, " id !=? and email=?", userID, u.data.Email)
			if ok {
				return nil, code.ERROR_EMAIL_EXIST
			}
		}

	}
	//2. 更新
	if err := respository.Updates(&models.User{}, &u.data, "id=?", userID); err != nil {
		fmt.Println(err)
		return nil, code.ERROR_UPDATE_USER
	}
	return &u.data, code.OK
}

func (u *UserInfoUpdate) check(userID int) (interface{}, code.Code) {
	if u.data.Email != "" {
		ok, _ := respository.Exist(&models.User{}, "email=? and id !=?", u.data.Email, userID)
		if ok {
			return nil, code.ERROR_EMAIL_EXIST
		}
	}
	return nil, code.OK
}

// ================================================================= 获取验证码

type Email struct {
	data *request.SendEmail
}

func SendEmail(data *request.SendEmail) (interface{}, code.Code) {
	return Email{
		data: data,
	}.Do()
}
func (email Email) Do() (interface{}, code.Code) {
	if _, c := email.check(); c != code.OK {
		return nil, c
	}
	if _, c := email.load2Redis(); c != code.OK {
		return nil, c
	}
	return nil, code.OK
}

func (email Email) check() (interface{}, code.Code) {
	//校验邮箱是否存在
	//1. 先在布隆中查看
	exists, err := global.Bloom.Exists(keys.EmailBloomKey, email.data.Email)
	if err != nil {
		return nil, code.ERROR_DB_OPE
	}
	if !exists {
		return nil, code.OK
	}

	exist, err := respository.Exist(&models.User{}, "email=?", email.data.Email)
	if err != nil {
		return nil, code.ERROR_DB_OPE
	}
	if exist {
		return nil, code.ERROR_EMAIL_EXIST
	}
	return nil, code.OK

}
func (email Email) load2Redis() (interface{}, code.Code) {
	randomCode := utils.GenerateRandomCode(6)
	if err := global.Redis.Set(
		keys.CodeKey+email.data.Email, randomCode, keys.CodeKeyDurationKey,
	).Err(); err != nil {
		return nil, code.ERROR_DB_OPE
	}
	msg := mq.EmailReq{
		Title:   "验证码",
		Message: "您的验证码为:" + randomCode,
		Users:   []string{email.data.Email},
	}
	go consumer.EmailConsumer.Product(&msg)
	return nil, code.OK
}

// ================================================================= 获取我的签到信息

type Sign struct {
}

func (s Sign) Do(userID int) (interface{}, code.Code) {
	today := time.Now()
	year := today.Year()
	month := today.Month()
	day := today.Day()
	key := keys.SignKey + fmt.Sprintf("%d:%d:%d", year, month, userID)
	//签到
	if err := global.Redis.SetBit(key, int64(day), 1).Err(); err != nil {
		return nil, code.ERROT_SIGN_ERROR
	}
	return nil, code.OK
}

func CreateSign(userID int) (interface{}, code.Code) {
	return Sign{}.Do(userID)
}

type MySign struct {
	data *request.SignList
}

func ListMySign(userID int, data *request.SignList) (interface{}, code.Code) {
	return MySign{
		data: data,
	}.Do(userID)
}
func (r MySign) Do(userID int) (interface{}, code.Code) {
	var signCount, signContinueMax, temp int64

	if r.data.Month == 0 && r.data.Year == 0 {
		//	本月
		now := time.Now()
		r.data.Year = now.Year()
		r.data.Month = int(now.Month())
	}
	//当前月的天数
	days := utils.CalMonth(r.data.Year, time.Month(r.data.Month))
	key := keys.SignKey + fmt.Sprintf("%d:%d:%d", r.data.Year, r.data.Month, userID)
	//签到表
	var signs = make([]int, days)
	for i := 0; i < days; i++ {
		val := global.Redis.GetBit(key, int64(i)).Val()
		if val == 1 {
			temp++
		} else {
			//统计连续最大签到次数
			signContinueMax = int64(math.Max(float64(signContinueMax), float64(temp)))
			temp = 0
		}
		signs[i] = int(val)
	}
	//签到次数
	signCount = global.Redis.BitCount(key, &redis.BitCount{Start: 0, End: -1}).Val()
	return response.NewSignResponse(
		signs, signCount, signContinueMax,
	), code.OK
}
