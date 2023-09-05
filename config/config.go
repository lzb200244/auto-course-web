package config

import (
	"auto-course-web/config/autoload"
	"auto-course-web/config/autoload/db"
)

type Configuration struct {
	//项目配置项
	Project autoload.Project `mapstructure:"project" json:"project" yaml:"project"`
	Mysql   db.Mysql         `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis   db.Redis         `mapstructure:"redis" json:"redis" yaml:"redis"`
	Log     autoload.Log     `mapstructure:"log" json:"log" yaml:"log"`
	//jwt配置项
	Jwt autoload.Jwt `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	//七牛云配置
	Qiniu autoload.Qiniu `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
}
