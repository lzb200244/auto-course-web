package service

import (
	"auto-course-web/global"
	"auto-course-web/global/keys"
	"auto-course-web/initialize"
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

/*
Created by 斑斑砖 on 2023/9/15.
Description：
*/
func TestCreateLua(t *testing.T) {
	initialize.InitConfig("../config/dev.conf.yml")
	initialize.InitLogger()
	initialize.InitRedis()
	//redis.NewScript()

	script := redis.NewScript(keys.Lua2CreateCourse)
	val, err := script.Run(global.Redis, []string{
		keys.SelectCourseKey,
		"1",
		keys.UserSelectedCourseListKey,
	}, []string{
		"1",
	}).Result()
	v, _ := val.(int64)

	fmt.Println(v, err)
}
