package consumer

import (
	"auto-course-web/global"
	"auto-course-web/global/variable"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

/*
Created by 斑斑砖 on 2023/9/15.
	Description： 只负责生产消息
*/

var RobConsumer *Rob

type Rob struct {
	channel *amqp.Channel
}

func InitRobListener() error {

	RobConsumer = &Rob{
		channel: global.RabbitMQ,
	}
	err := RobConsumer.Declare()
	if err != nil {
		return err
	}
	return nil

}
func (rob Rob) Declare() error {
	err := rob.channel.ExchangeDeclare(variable.RobExchange, variable.RobKind,
		true, false, false, false, nil,
	)
	if err != nil {
		return err
	}
	_, err = rob.channel.QueueDeclare(variable.RobQueue, true,
		false, false, false, nil,
	)
	if err != nil {
		return err
	}
	// 将队列绑定到交换机上
	err = rob.channel.QueueBind(variable.RobQueue, variable.RobRoutingKey,
		variable.RobExchange, false, nil)
	if err != nil {
		return err
	}
	return nil
}

// Consumer 处理抢课成功的操作,记录日志,入库
func (rob Rob) Consumer() {

}
func (rob Rob) Product(msg *Rob) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = rob.channel.Publish(
		variable.RobExchange,
		variable.RobRoutingKey,
		false,
		false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		})
	if err != nil {
		fmt.Println(err)
		//TODO log
		return
	}
}
