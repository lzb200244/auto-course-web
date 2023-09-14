package keys

import "time"

/*
Created by 斑斑砖 on 2023/9/8.
Description：
 课程的Key
*/

const (
	// ====================================== 预发布阶段

	IsPreLoadedKey         = "preselection:status" //是否处于预发布阶段
	PreSelectedDurationKey = time.Second * 60      //key过期时间，也就是预发布时间
	// PreLoadCourseKey 课程预加载所需的Key，id为key容量为value
	PreLoadCourseKey = "course:preload" // hash
	// PreLoadCourseListKey 已经发布的列表
	PreLoadCourseListKey = "course:preload:set" //set

	// ====================================== 选课阶段

	IsSelectCourseKey       = "selection:status"  //是否开启选课阶段
	SelectCourseDurationKey = time.Second * 120   //key过期时间
	SelectCourseKey         = "course:select"     //hash
	SelectCourseListKey     = "course:select:set" //set

)
