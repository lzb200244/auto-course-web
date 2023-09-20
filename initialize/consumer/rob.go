package consumer

import (
	"auto-course-web/global"
	"auto-course-web/global/variable"
	"auto-course-web/models"
	"auto-course-web/models/mq"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
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
	RobConsumer.Consumer()
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
	//接收消息
	results, err := rob.channel.Consume(
		variable.RobQueue,
		variable.RobRoutingKey,
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
			var userCourse = models.UserCourse{
				UserID:    uint(msg.UserID),
				CourseID:  uint(msg.CourseID),
				CreatedAt: time.Now(),
			}
			// ============================================================= 开启事物
			err = global.MysqlDB.Transaction(func(tx *gorm.DB) error {
				// Set the isolation level to Serializable
				tx = tx.Set("gorm:query_option", "SERIALIZABLE")

				// Use SELECT ... FOR UPDATE to lock the rows
				err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).FirstOrCreate(&userCourse).Error
				if err != nil {
					if errors.Is(err, gorm.ErrDuplicatedKey) {
						return nil
					}
					return err
				}
				return nil
			})
			if err != nil {
				continue
			}
			d.Ack(false)
		}
	}()
}
func (rob Rob) Product(msg *mq.CourseReq) {
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
