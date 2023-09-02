package initialize

import (
	"github.com/spf13/viper"
	"go-template/config"
	"go-template/global"
)

// InitConfig 初始化viper加载配置文件
func InitConfig() {
	v := viper.New()
	v.SetConfigFile("./config/dev.conf.yml")
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
