package initialize

import (
	"auto-course-web/global"
	"github.com/streadway/amqp"
)

/*
Created by 斑斑砖 on 2023/9/13.
Description：
	rabbitmq初始化
*/

func InitRabbitMQ() {
	conn, err := amqp.Dial(global.Config.RabbitMQ.Dsn())
	if err != nil {
		panic("连接rabbitmq失败")
	}
	// 创建一个通道
	ch, err := conn.Channel()
	if err != nil {
		panic("创建通道失败")
	}
	global.RabbitMQ = ch

	//TODO 初始化队列和交换机等操作
	global.Logger.Debug("rabbitmq初始化成功！")

}
