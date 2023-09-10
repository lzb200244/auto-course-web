package utils

import (
	"github.com/gin-gonic/gin"
)

const UserKey = "user"

type UserInfo struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
}

/*
func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息错误")
	}
	return user, nil
}

func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo)

	return u, ok
}

*/

// SetUser 封装了User供给ctx上下链
func SetUser(ctx *gin.Context, u *UserInfo) {
	ctx.Set(UserKey, u)
}
func GetUser(ctx *gin.Context) (*UserInfo, bool) {
	user, exists := ctx.Get(UserKey)
	if exists {
		return user.(*UserInfo), true
	}
	return nil, false
}
