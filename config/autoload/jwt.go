package autoload

import "time"

/*
Created by 斑斑砖 on 2023/8/15.
Description：
	jwt的配置项
*/

type Jwt struct {
	Expire time.Duration `json:"expire" mapstructure:"expire"` // 过期时间
	Issuer string        `json:"issuer" mapstructure:"issuer"` // 签发者
	SECRET string        `json:"secret" mapstructure:"secret"` // 密钥
}
