package autoload

/*
Created by 斑斑砖 on 2023/9/21.
Description：

*/

type MultiAvatar struct {
	Secret string `json:"secret" mapstructure:"secret"` // 密钥
	Url    string `json:"url" mapstructure:"url"`       // 图片地址
}
