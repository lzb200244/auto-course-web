package utils

import (
	"auto-course-web/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/*
Created by 斑斑砖 on 2023/8/14.
Description：
	绑定、获取参数
*/

// BindValidJson Json 绑定验证 + 合法性校验
func BindValidJson[T any](c *gin.Context) (data T, err error) {
	// Json 绑定
	if err := c.ShouldBindJSON(&data); err != nil {
		//一般都是非法请求参数了
		global.Logger.Info("BindValidJson", zap.Error(err))
	}
	// 参数合法性校验
	err = Validator.Validate(&data)
	return data, err
}
func BindValidQuery[T any](c *gin.Context) (data T, err error) {
	// Json 绑定
	if err := c.BindQuery(&data); err != nil {
		//一般都是非法请求参数了
		global.Logger.Info("BindQuery", zap.Error(err))
	}
	// 参数合法性校验
	err = Validator.Validate(&data)
	return data, err
}

//func Validate(c *gin.Context, data any) {
//	validMsg := Validator.Validate(data)
//	if validMsg != nil {
//		Fail(c, 1000, validMsg.Error(), nil)
//	}
//}
