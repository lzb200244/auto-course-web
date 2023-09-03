package qiniu

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go-template/global"
)

/*
Created by 斑斑砖 on 2023/9/3.
Description：

*/

// GetCredits  获取凭证
func GetCredits() string {
	mac := qbox.NewMac(global.Config.Qiniu.AccessKey, global.Config.Qiniu.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope: global.Config.Qiniu.Bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
