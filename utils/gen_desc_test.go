package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

/*
Created by 斑斑砖 on 2023/9/21.
Description：

*/

// GenerateDesc 随机生成个性签名
// https://v1.hitokoto.cn/?c=b
func TestGenerateDesc(t *testing.T) {
	get, err := http.Get("https://v1.hitokoto.cn/?c=b")
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
