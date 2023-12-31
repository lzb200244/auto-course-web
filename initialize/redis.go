package initialize

import (
	"auto-course-web/global"
	bloom "github.com/RedisBloom/redisbloom-go"
	"github.com/go-redis/redis"
)

/*
Created by 斑斑砖 on 2023/9/8.
Description：
	初始化redis
*/

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,     // url
		Password: global.Config.Redis.Password, // no password set
		DB:       global.Config.Redis.DB,
	})
	if err := client.Ping().Err(); err != nil {
		panic(err)
	}
	global.Bloom = bloom.NewClient("192.168.10.131:6379", "root", nil)

	global.Redis = client
	global.Logger.Debug("redis初始化成功！")
}
