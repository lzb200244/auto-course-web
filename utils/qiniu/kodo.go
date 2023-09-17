package qiniu

import (
	"auto-course-web/global"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

/*
Created by 斑斑砖 on 2023/9/3.
Description：

*/

// GetCredits  获取凭证
func GetCredits(bucket string) string {
	mac := qbox.NewMac(global.Config.Qiniu.AccessKey, global.Config.Qiniu.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
