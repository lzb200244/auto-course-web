package initialize

import (
	"auto-course-web/global"
	"auto-course-web/initialize/consumer"
	"github.com/streadway/amqp"
	"time"
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

	//初始化mq
	err = consumer.InitEmailListener()

	if err != nil {
		panic("InitEmailListener初始化失败")
	}
	err = consumer.InitRobListener()
	if err != nil {
		panic("InitRobListener初始化失败")
	}
	err = consumer.InitPushListener()

	if err != nil {
		panic("InitPushListener初始化失败")
	}
	go func() {
		for true {
			consumer.PushConsumer.Product("ok")
			time.Sleep(time.Second)
		}
	}()
	//TODO 初始化队列和交换机等操作
	global.Logger.Debug("rabbitmq初始化成功！")

}
