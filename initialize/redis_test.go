package initialize

import (
	"auto-course-web/global"
	"fmt"
	"testing"
)

/*
Created by 斑斑砖 on 2023/9/8.
Description：
*/
func TestRedisConn(t *testing.T) {
	InitConfig("../config/dev.conf.yml")
	InitLogger()
	InitRedis()
	ping := global.Redis.Ping()
	fmt.Println(ping)

}
