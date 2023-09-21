package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*
Created by 斑斑砖 on 2023/9/21.
Description：
*/
type Hitokoto struct {
	Id         int         `json:"id"`
	Uuid       string      `json:"uuid"`
	Hitokoto   string      `json:"hitokoto"`
	Type       string      `json:"type"`
	From       string      `json:"from"`
	FromWho    interface{} `json:"from_who"`
	Creator    string      `json:"creator"`
	CreatorUid int         `json:"creator_uid"`
	Reviewer   int         `json:"reviewer"`
	CommitFrom string      `json:"commit_from"`
	CreatedAt  string      `json:"created_at"`
	Length     int         `json:"length"`
}

// GenerateDesc 随机生成个性签名
// https://v1.hitokoto.cn/?c=b
func GenerateDesc() string {
	const base = "个性签名"
	res, err := http.Get("https://v1.hitokoto.cn/?c=b")
	if err != nil {
		return base
	}
	if res.StatusCode != 200 {
		return base
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return base
	}
	var hitokoto *Hitokoto
	if err := json.Unmarshal(body, &hitokoto); err != nil {
		return base
	}
	return hitokoto.Hitokoto
}
