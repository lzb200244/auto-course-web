package response

/*
Created by 斑斑砖 on 2023/9/15.
Description：
*/

type ElectionsResponse struct {
	CourseID int `json:"courseID"`
	Capacity int `json:"capacity"`
}
type SelectCourseResponse struct {
	CourseResponse
	Capacity int `json:"capacity" gorm:"-"`
	Left     int `json:"left" gorm:"-"`
}
