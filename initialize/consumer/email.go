package consumer

import (
	"auto-course-web/global"
	"auto-course-web/global/variable"
	"auto-course-web/models/mq"
	"auto-course-web/utils/tencent"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

/*
Created by 斑斑砖 on 2023/9/15.
Description：
*/

var EmailConsumer *Email

type Email struct {
	channel *amqp.Channel
}

func InitEmailListener() error {
	EmailConsumer = &Email{
		channel: global.RabbitMQ,
	}
	err := EmailConsumer.Declare()
	if err != nil {
		return err
	}

	EmailConsumer.Consumer()
	return nil

}

func (e Email) Declare() error {
	err := e.channel.ExchangeDeclare(variable.EmailExchange, amqp.ExchangeDirect,
		true, false, false, false, nil,
	)
	if err != nil {
		return err
	}
	_, err = e.channel.QueueDeclare(variable.EmailQueue, true,
		false, false, false, nil,
	)
	if err != nil {
		return err
	}
	// 将队列绑定到交换机上
	err = e.channel.QueueBind(variable.EmailQueue, variable.EmailRoutingKey,
		variable.EmailExchange, false, nil)
	if err != nil {
		return err
	}
	return nil
}
func (e Email) Consumer() {
	//接收消息
	results, err := e.channel.Consume(
		variable.EmailQueue,
		variable.EmailRoutingKey,
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
			var msg *mq.EmailReq
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				//TODO log
				fmt.Println(err)
				continue
			}
			tencent.SendEmail(
				msg.Title, msg.Message, msg.Users,
			)
			d.Ack(false)
		}
	}()

}

func (e Email) Product(msg *mq.EmailReq) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = e.channel.Publish(
		variable.EmailExchange,
		variable.EmailRoutingKey,
		false,
		false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		})
	fmt.Println(msg)
	if err != nil {
		fmt.Println(err)
		//TODO log
		return
	}
}
