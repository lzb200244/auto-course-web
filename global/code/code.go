package code

/*
Created by 斑斑砖 on 2023/8/14.
Description：
	// 错误码汇总
*/

type Code int

const (
	OK   Code = 0
	FAIL Code = 500

	// code= 9000... 通用错误
	ERROR_REQUEST_PARAM Code = 9001
	ERROR_REQUEST_PAGE  Code = 9002
	ERROR_INVALID_PARAM Code = 9003
	ERROR_DB_OPE        Code = 9004

	// code= 1000... 用户模块的错误
	ERROR_USER_NAME_USED    Code = 1001
	ERROR_PASSWORD_WRONG    Code = 1002
	ERROR_USER_NOT_EXIST    Code = 1003
	ERROR_USER_NO_RIGHT     Code = 1009
	ERROR_OLD_PASSWORD      Code = 1010
	ERROR_EMAIL_SEND        Code = 1011
	ERROR_EMAIL_HAS_SEND    Code = 1012
	ERROR_VERIFICATION_CODE Code = 1013

	// code = 1200.. 鉴权相关错误
	ERROR_TOKEN_NOT_EXIST  Code = 1201
	ERROR_TOKEN_RUNTIME    Code = 1202
	ERROR_TOKEN_WRONG      Code = 1203
	ERROR_TOKEN_TYPE_WRONG Code = 1204
	ERROR_TOKEN_CREATE     Code = 1205
	ERROR_PERMI_DENIED     Code = 1206
	FORCE_OFFLINE          Code = 1207
	LOGOUT                 Code = 1208

	// code=6000... 权限模块的错误
	ERROR_ROLE_NAME_EXIST Code = 60010
	ERROR_ROLE_NOT_EXIST  Code = 60011

	// code=7000 ... 页面模块的错误
)

var codeMsg = map[Code]string{
	OK:   "OK",
	FAIL: "FAIL",

	// ###################################################################### 通用错误
	ERROR_REQUEST_PARAM: "请求参数格式错误",
	ERROR_REQUEST_PAGE:  "分页参数错误",
	ERROR_INVALID_PARAM: "不合法的请求参数",
	ERROR_DB_OPE:        "数据库操作异常",

	// ###################################################################### 登录与注册
	ERROR_USER_NAME_USED:    "用户名已存在",
	ERROR_USER_NOT_EXIST:    "该用户不存在",
	ERROR_PASSWORD_WRONG:    "密码错误",
	ERROR_USER_NO_RIGHT:     "该用户无权限",
	ERROR_OLD_PASSWORD:      "旧密码不正确",
	ERROR_EMAIL_SEND:        "邮件发送失败",
	ERROR_EMAIL_HAS_SEND:    "已朝该邮箱发送验证码（有效期 15 分钟），请检查回收站",
	ERROR_VERIFICATION_CODE: "验证码错误",

	// ###################################################################### 权限
	ERROR_TOKEN_NOT_EXIST:  "TOKEN 不存在，请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN 已过期，请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN 不正确，请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN 格式错误，请重新登陆",
	ERROR_TOKEN_CREATE:     "TOKEN 生成失败",
	ERROR_PERMI_DENIED:     "权限不足",
	FORCE_OFFLINE:          "您已被强制下线",
	LOGOUT:                 "您已退出登录",

	ERROR_ROLE_NAME_EXIST: "该角色名已经存在",
	ERROR_ROLE_NOT_EXIST:  "该角色不存在",
}

func GetMsg(code Code) string {
	return codeMsg[code]
}
