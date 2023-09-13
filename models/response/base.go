package response

/*
Created by 斑斑砖 on 2023/9/8.
Description：
*/

type List struct {
	Results interface{} `json:"results"`
	Count   int64       `json:"count"`
}
