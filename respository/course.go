package respository

import (
	"auto-course-web/global"
	"auto-course-web/models/request"
	"auto-course-web/respository/scopes"
)

/*
Created by 斑斑砖 on 2023/9/16.
Description：
*/

func QueryCourseList[T any](model any, data T, pager *request.Pages, order, query string, args ...any) (int64, error) {
	var count int64
	sql := global.MysqlDB.Model(model).Preload("Category").Preload("College").Order(order)
	if query != "" {
		sql.Where(query, args...)
	}
	err := sql.Count(&count).Error
	if err != nil {
		return 0, err
	}
	if pager != nil {
		sql = sql.Scopes(
			scopes.Paginate(pager.Page, pager.Size),
		)
	}
	err = sql.Find(&data).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
