package config

import (
	"auto-course-web/config/autoload"
	"auto-course-web/config/autoload/db"
)

type Configuration struct {
	//项目配置项
	Project autoload.Project `mapstructure:"project" json:"project" yaml:"project"`
	//mysql配置
	Mysql db.Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//redis配置
	Redis db.Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	//日志配置
	Log autoload.Log `mapstructure:"log" json:"log" yaml:"log"`
	//rabbitmq配置
	RabbitMQ autoload.RabbitMQ `mapstructure:"rabbitmq" json:"rabbitmq" yaml:"rabbitmq"`
	//jwt配置项
	Jwt autoload.Jwt `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	//七牛云配置
	Qiniu autoload.Qiniu `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	//邮件配置
	Email autoload.Email `mapstructure:"email" json:"email" yaml:"email"`

	MultiAvatar autoload.MultiAvatar `mapstructure:"multiavatar" json:"multiavatar" yaml:"multiavatar"`
}
