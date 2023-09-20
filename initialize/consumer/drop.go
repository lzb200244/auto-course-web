package consumer

import (
	"auto-course-web/global"
	"auto-course-web/global/variable"
	"auto-course-web/models"
	"auto-course-web/models/mq"
	"auto-course-web/respository"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

/*
Created by 斑斑砖 on 2023/9/15.
	Description：
*/

var DropConsumer *Drop

type Drop struct {
	channel *amqp.Channel
}

func InitDropListener() error {

	DropConsumer = &Drop{
		channel: global.RabbitMQ,
	}
	err := DropConsumer.Declare()
	if err != nil {
		return err
	}
	DropConsumer.Consumer()
	return nil

}
func (drop Drop) Declare() error {
	err := drop.channel.ExchangeDeclare(variable.DropExchange, variable.RobKind,
		true, false, false, false, nil,
	)
	if err != nil {
		return err
	}
	_, err = drop.channel.QueueDeclare(variable.DropQueue, true,
		false, false, false, nil,
	)
	if err != nil {
		return err
	}
	// 将队列绑定到交换机上
	err = drop.channel.QueueBind(variable.DropQueue, variable.DropRoutingKey,
		variable.DropExchange, false, nil)
	if err != nil {
		return err
	}
	return nil
}

// Consumer 处理抢课成功的操作,记录日志,入库
func (drop Drop) Consumer() {
	//接收消息
	results, err := drop.channel.Consume(
		variable.DropQueue,
		variable.DropRoutingKey,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//启用协程处理消息
	go func() {
		for d := range results {
			//消息逻辑处理，可以自行设计逻辑
			var msg *mq.CourseReq
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				//TODO log
				fmt.Println(err)
				continue
			}
			//进行退课

			err = respository.Delete(
				models.UserCourse{},
				"user_id = ? and course_id = ?",
				msg.UserID, msg.CourseID)
			if err != nil {
				fmt.Println(err)
				return
			}

			d.Ack(false)
		}
	}()
}
func (drop Drop) Product(msg *mq.CourseReq) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = drop.channel.Publish(
		variable.DropExchange,
		variable.DropRoutingKey,
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
