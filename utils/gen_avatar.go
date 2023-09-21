package utils

import "auto-course-web/global"

/*
Created by 斑斑砖 on 2023/9/21.
Description：
	生成头像
*/

// GenerateAvatar 生成头像
// https://api.multiavatar.com/2.png?apikey=dHmzRGLe5hFUsY
func GenerateAvatar(key string) string {
	return global.Config.MultiAvatar.Url + key + "?" + global.Config.MultiAvatar.Secret
}
