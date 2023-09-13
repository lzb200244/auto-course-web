package global

import (
	"auto-course-web/config"
	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config   config.Configuration //全局的配置文件
	MysqlDB  *gorm.DB             // mysql实例
	Logger   *zap.Logger          // logger实例
	Redis    *redis.Client
	RabbitMQ *amqp.Channel
)
