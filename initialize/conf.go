package initialize

import (
	"auto-course-web/config"
	"auto-course-web/global"
	"github.com/spf13/viper"
)

// InitConfig 初始化viper加载配置文件
func InitConfig(path string) {

	v := viper.New()
	v.SetConfigType("yaml")
	//v.SetConfigName("dev.conf") // 设置配置文件名
	//v.AddConfigPath("./config")
	if path != "" {
		v.SetConfigFile(path)
	} else {
		v.SetConfigFile("./config/dev.conf.yml")
	}

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.Configuration{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	// 传递给全局变量
	global.Config = serverConfig

}
