package autoload

/*
Created by 斑斑砖 on 2023/8/13.
Description：
	项目配置
*/

type Project struct {
	Name   string `mapstructure:"name" json:"name" yaml:"name"`
	Domain string `mapstructure:"domain" json:"domain" yaml:"domain"`
	Port   string `mapstructure:"port" json:"port" yaml:"port"`
	LogDir string `mapstructure:"logDir" json:"logDir" yaml:"logDir"`
	Mode   string `mapstructure:"mode" json:"mode" yaml:"mode"`
}
