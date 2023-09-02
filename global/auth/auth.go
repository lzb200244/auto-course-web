package auth

/*
Created by 斑斑砖 on 2023/8/15.
Description：
*/

const (
	Student int = iota + 1
	Teacher
	Admin
)

var authMap = map[int]string{
	Student: "student",
	Teacher: "teacher",
	Admin:   "admin",
}

func GetAuthorityName(c int) string {
	return authMap[c]
}
