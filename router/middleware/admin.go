package middleware

import (
	"auto-course-web/global/auth"
	"auto-course-web/global/code"
	"auto-course-web/utils"
	"github.com/gin-gonic/gin"
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
