package response

/*
Created by 斑斑砖 on 2023/8/14.
	Description：
		用户相关的响应
*/

type UserResponse struct {
	ID         uint     `json:"id"`
	UserName   string   `json:"username" `
	Name       string   `json:"name"`
	Email      string   `json:"email" `
	Sex        int      `json:"sex"`
	Desc       string   `json:"desc"`
	Avatar     string   `json:"avatar"`
	Roles      []string `json:"roles"`
	Permission []int    `json:"permission"`
	Token      string   `json:"token"`
}

func NewUserResponse(ID uint, userName, name, email, desc, avatar, token string, sex int, roleName []string, permission []int) *UserResponse {
	return &UserResponse{
		ID:         ID,
		UserName:   userName,
		Name:       name,
		Email:      email,
		Sex:        sex,
		Desc:       desc,
		Avatar:     avatar,
		Roles:      roleName,
		Permission: permission,
		Token:      token,
	}
}
