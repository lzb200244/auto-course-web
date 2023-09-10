package keys

import "time"

/*
Created by 斑斑砖 on 2023/9/8.
Description：
 课程的Key
*/

const (
	// ====================================== 预发布阶段

	IsPreLoadedKey       = "course:open"    //是否开启预发布通道
	PreLoadedDurationKey = time.Second * 60 //key过期时间，也就是预发布时间
	// PreLoadCourseKey 课程预加载所需的Key，id为key容量为value
	PreLoadCourseKey = "course:preload" // hash
	// PreLoadCourseListKey 已经发布的列表
	PreLoadCourseListKey = "course:preload:set" //set

)
