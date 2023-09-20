package service

import (
	"auto-course-web/initialize"
	"auto-course-web/models/request"
	"fmt"
	"testing"
)

/*
Created by 斑斑砖 on 2023/9/20.
Description：

*/

func TestListMySign(t *testing.T) {
	initialize.InitConfig("../config/dev.conf.yml")
	initialize.InitLogger()
	initialize.InitRedis()
	req := request.SignList{Year: 2023, Month: 9}
	//	获取我的签到信息
	result, _ := ListMySign(1, &req)
	
	fmt.Println(result)

}
func TestSign_Do(t *testing.T) {
	initialize.InitConfig("../config/dev.conf.yml")
	initialize.InitLogger()
	initialize.InitRedis()
	//进行签到
	sign := Sign{}
	result, _ := sign.Do(1)
	fmt.Println(result)
}
func TestRegister(t *testing.T) {

}
