package middleware

import (
	"auto-course-web/global/auth"
	"auto-course-web/global/code"
	"auto-course-web/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// HasRole 是否属于角色访问的范围
func HasRole(role auth.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		userObj, _ := utils.GetUser(c)
		fmt.Println(role, userObj.Role)
		if userObj.Role < int(role) {
			utils.Fail(c, code.ERROR_PERMI_DENIED, code.GetMsg(code.ERROR_PERMI_DENIED), nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
