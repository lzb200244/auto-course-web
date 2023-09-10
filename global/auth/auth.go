package auth

/*
Created by 斑斑砖 on 2023/8/15.
Description：
*/

type Auth int

const (
	Student Auth = iota + 1
	Teacher
	Admin
)

var authMap = map[Auth]string{
	Student: "student",
	Teacher: "teacher",
	Admin:   "admin",
}

func GetAuthorityName(c Auth) string {
	return authMap[c]
}
