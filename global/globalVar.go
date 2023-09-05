package global

import (
	"auto-course-web/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config  config.Configuration //全局的配置文件
	MysqlDB *gorm.DB             // mysql实例
	Logger  *zap.Logger          // logger实例
)
