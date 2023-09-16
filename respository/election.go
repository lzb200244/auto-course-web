package respository

import (
	"auto-course-web/global"
	"auto-course-web/models/request"
	"auto-course-web/respository/scopes"
)

/*
Created by 斑斑砖 on 2023/9/14.
Description：

	选课的dao
*/

// QuerySelectCourse 查询选课的课程，带分页，根据条件，带排序，带分类，带标题
func QuerySelectCourse[T any](model any, data T, pager *request.Pages, title, query, order string, categoryID uint, args ...any) (int64, error) {
	var count int64
	sql := global.MysqlDB.Model(model).Preload("Category").Order(order)

	if categoryID != 0 {
		sql.Where("category_id=?", categoryID)
	}

	if query != "" {
		sql.Where(query, args...)
	}
	if title != "" {
		sql.Where("title like ? ", "%"+title+"%")
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
