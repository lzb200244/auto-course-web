package autoload

/*
Created by 斑斑砖 on 2023/9/10.
Description：

*/

type Email struct {
	User string `mapstructure:"user" json:"user" yaml:"user"`
	Name string `mapstructure:"name" json:"name" yaml:"name"`
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Pass string `mapstructure:"pass" json:"pass" yaml:"pswd"`
}
