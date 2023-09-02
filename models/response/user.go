package response

/*
Created by 斑斑砖 on 2023/8/14.
Description：

*/

type UserResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"username" `
	Email    string `json:"email" `
	Token    string `json:"token"`
}

func NewUserResponse(ID uint, userName string, email string, token string) *UserResponse {
	return &UserResponse{ID: ID, UserName: userName, Email: email, Token: token}
}
