package middleware

import (
	"github.com/gin-gonic/gin"
	"go-template/global/auth"
	"go-template/global/code"
	"go-template/utils"
)

// IsAdmin 判断是否是管理员
func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {

		userObj, _ := utils.GetUser(c)
		if userObj.Authority != auth.Admin {
			utils.Fail(c, code.ERROR_PERMI_DENIED, code.GetMsg(code.ERROR_PERMI_DENIED), nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
