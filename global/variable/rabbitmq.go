package variable

import "github.com/streadway/amqp"

/*
Created by 斑斑砖 on 2023/9/15.
Description：
*/
const (
	// ==================================================================== 邮箱

	EmailExchange   = "auto-course:email_exchange"
	EmailRoutingKey = "auto-course:email_routing_key"
	EmailQueue      = "auto-course:email_queue"

	// =================================================================== 抢课

	RobExchange   = "auto-course:rob_exchange"
	RobRoutingKey = "auto-course:rob_routing_key"
	RobKind       = amqp.ExchangeDirect
	RobQueue      = "auto-course:rob_queue"

	// =================================================================== 退课

	DropExchange   = "auto-course:drop_exchange"
	DropRoutingKey = "auto-course:drop_routing_key"
	DropKind       = amqp.ExchangeDirect
	DropQueue      = "auto-course:drop_queue"

	// =================================================================== 推送 进行给前端websocket进行消费

	PushExchange   = "auto-course:push_exchange"
	PushRoutingKey = "auto-course:push_routing_key"
	PushKind       = amqp.ExchangeFanout
	PushQueue      = "auto-course:push_queue"
)
