package keys

import "time"

/*
Created by 斑斑砖 on 2023/9/18.
Description：
*/
const (
	// ============================================================= 验证码

	CodeKey            = "code:" //=>string
	CodeKeyDurationKey = time.Second * 120

	// ============================================================= 签到

	SignKey = "sign:" //=>bitmap sign:年:月:userID

	// ============================================================= 布隆key

	UserBloomKey  = "users:bloom"
	EmailBloomKey = "email:bloom"
)
