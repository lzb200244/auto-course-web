package utils

import (
	"math/rand"
	"time"
)

/*
Created by 斑斑砖 on 2023/9/16.
Description：
	随机生成课程code
*/

func GenerateRandomCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	charset := "0123456789"
	charsetLength := len(charset)
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = charset[rand.Intn(charsetLength)]
	}
	return string(code)
}
