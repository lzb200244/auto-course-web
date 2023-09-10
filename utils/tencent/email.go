package tencent

import (
	"auto-course-web/global"
	"gopkg.in/gomail.v2"
)

/*
Created by 斑斑砖 on 2023/9/10.
Description：
	腾讯相关的第三方
*/

func SendEmail(title, message string, users []string) {
	//1. 创建发件器
	mailer := gomail.NewMessage()
	//2. 设置发件人
	mailer.SetHeader("From", global.Config.Email.User)
	//3. 设置收件人
	mailer.SetHeader("To", users...)
	//4. 设置标题
	mailer.SetHeader("Subject", title)
	//5. 设置正文
	mailer.SetBody("text/html", message)

	//	配置smtp服务器信息
	d := gomail.NewDialer(
		global.Config.Email.Host,
		global.Config.Email.Port,
		global.Config.Email.User,
		global.Config.Email.Pass,
	)
	//6. 发送
	if err := d.DialAndSend(mailer); err != nil {
		panic(err)
	}
}
