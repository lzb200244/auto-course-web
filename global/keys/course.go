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

	SelectCourseMax         = 5
	IsSelectCourseKey       = "selection:status"  //是否开启选课阶段
	SelectCourseDurationKey = time.Second * 120   //key过期时间
	SelectCourseKey         = "course:select"     //hash
	SelectCourseListKey     = "course:select:set" //set

	// ====================================== 用户的Key

	UserSelectedCourseListKey = "user:selected:" //set

	//	Lua脚本
	Lua2CreateCourse = `
	if tonumber(redis.call('hget', KEYS[1], KEYS[2])) > 0 then
		-- 进行对应课程-1操作
    	redis.call('hincrby', KEYS[1], KEYS[2], -1)
        -- 进行添加到用户已选集合
    	redis.call('sadd', KEYS[3] .. ARGV[1], KEYS[2])
    	return 1 
	else
		return 0 
	end
	`
	Lua2CancelCourse = `
	redis.call('hincrby', KEYS[1], KEYS[2], 1)
	redis.call('srem', KEYS[3] .. ARGV[1], KEYS[2])
	return 1 
	`
	//key 1 用户的key，key2,课程的key，key3课程id的key
	LuaScript2SelectCourse = `
		-- 判断用户是否选过该课程  
		if redis.call('sismember', KEYS[1], KEYS[3]) == 0 then  
			-- 选课操作  
			local count = tonumber(redis.call('hget', KEYS[2], KEYS[3]))  
		  
			if count and count > 0 then  
				-- 进行对应课程-1操作  
				redis.call('hincrby', KEYS[2], KEYS[3], -1)  
				-- 进行添加到用户已选集合  
				redis.call('sadd', KEYS[1], KEYS[3])  
				return {1, count - 1} -- 0 选课成功  
			else  
				return {-1, 0} -- 1 课程抢完了  
			end  
		else  
			-- 退课操作  
			local count = tonumber(redis.call('hget', KEYS[2], KEYS[3]))  
			redis.call('hincrby', KEYS[2], KEYS[3], 1)  
			redis.call('srem', KEYS[1], KEYS[3])  
			return {2, count + 1} -- 2 退课成功  
		end
		`
)
