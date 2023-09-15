package mq

/*
Created by 斑斑砖 on 2023/9/15.
Description：

*/
//title, message string, users []string

type EmailReq struct {
	Title   string
	Message string
	Users   []string
}
