package middleware

import (
	"auto-course-web/global/code"
	"auto-course-web/utils"
	"github.com/gin-gonic/gin"

	"time"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		cd := code.OK
		if token == "" {
			//缺失token
			utils.Fail(c, code.ERROR_TOKEN_NOT_EXIST, code.GetMsg(code.ERROR_TOKEN_NOT_EXIST), nil)
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			//token解析失败
			cd = code.ERROR_TOKEN_TYPE_WRONG
		} else if time.Now().Unix() > claims.ExpiresAt {
			//过期
			cd = code.ERROR_TOKEN_RUNTIME
		}
		if cd != code.OK {
			utils.Fail(c, cd, code.GetMsg(cd), nil)
			return
		}
		/*
			在原来旧的上下文基础上添加信息,然后使用新的上下文替换旧的上下文,本质是是在原来的上下文添加了新的信息
			c.Request = c.Request.WithContext(utils.NewContext(c.Request.Context(), &utils.UserInfo{
				ID:       claims.Id,
				UserName: claims.Username,
				Email:    claims.Email,
			}))

		*/
		utils.SetUser(c, &utils.UserInfo{
			ID:       claims.Id,
			UserName: claims.Username,
			Email:    claims.Email,
			Role:     claims.Role,
		})
		c.Next()
	}
}
