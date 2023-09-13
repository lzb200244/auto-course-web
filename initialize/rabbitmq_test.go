package initialize

import (
	"testing"
)

/*
Created by 斑斑砖 on 2023/9/13.
Description：
	rabbitmq连接测试
*/

func TestInitRabbitMQ(t *testing.T) {
	InitConfig("../config/dev.conf.yml")
	InitLogger()
	InitRabbitMQ()

}
