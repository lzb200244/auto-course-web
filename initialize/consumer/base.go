package consumer

/*
Created by 斑斑砖 on 2023/9/15.
Description：
*/

type Rabbitmq[T any] interface {
	Declare() error
	Consumer()
	Product(msg T)
}
