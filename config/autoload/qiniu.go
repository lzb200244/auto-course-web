package autoload

/*
	Created by 斑斑砖 on 2023/9/3.
		Description：七牛云的配置
*/

type Qiniu struct {
	AccessKey string `mapstructure:"access-key" json:"access-key" yaml:"access-key"`
	SecretKey string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`
	Bucket    string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
}
