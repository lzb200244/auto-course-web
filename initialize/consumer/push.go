package consumer

import (
	"auto-course-web/global"
	"auto-course-web/global/keys"
	"auto-course-web/global/variable"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"math/rand"
	"strconv"
)

/*
Created by 斑斑砖 on 2023/9/15.
	Description： 推送到web前端,进行实现消费数据
*/

var PushConsumer *Push

type Push struct {
	channel *amqp.Channel
}

func InitPushListener() error {

	PushConsumer = &Push{
		channel: global.RabbitMQ,
	}
	err := PushConsumer.Declare()
	if err != nil {
		return err
	}
	return nil
}
func (push Push) Declare() error {
	err := push.channel.ExchangeDeclare(variable.PushExchange, variable.PushKind,
		false, false, false, false, nil,
	)
	if err != nil {
		return err
	}
	_, err = push.channel.QueueDeclare(variable.PushQueue, true,
		false, false, false, nil,
	)
	if err != nil {
		return err
	}
	// 将队列绑定到交换机上
	err = push.channel.QueueBind(variable.PushQueue, variable.PushRoutingKey,
		variable.PushExchange, false, nil)
	if err != nil {
		return err
	}
	return nil
}

// Consumer 处理抢课成功的操作,记录日志,入库
func (push Push) Consumer() {

}
func (push Push) Product(v interface{}) {
	type Push struct {
		ID    int `json:"id,omitempty"`
		Value int `json:"value,omitempty"`
	}
	var data []*Push
	msg := global.Redis.HGetAll(keys.SelectCourseKey).Val()
	for k, v := range msg {
		id, _ := strconv.Atoi(k)
		value, _ := strconv.Atoi(v)
		data = append(data, &Push{
			ID:    id,
			Value: value + rand.Intn(50),
		})
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = push.channel.Publish(
		variable.PushExchange,
		variable.PushRoutingKey,
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
