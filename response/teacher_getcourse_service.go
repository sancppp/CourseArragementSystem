package response

import (
	"CAS/db/mysql"
	"CAS/model"
	"CAS/types"
	"fmt"
)

// 获取老师下所有课程
// Method：Get
type GetTeacherCourseRequest struct {
	TeacherID string
}

type GetTeacherCourseResponse struct {
	Code types.ErrNo
	Data struct {
		CourseList []*types.TCourse
	}
}

func (serv *GetTeacherCourseRequest) GetCourse() (res GetTeacherCourseResponse) {
	var users []model.Course2Teacher
	mysql.MysqlDB.GetConn().Where("teacher_id = ?", serv.TeacherID).Find(&users)
	res.Code = types.OK
	for _, item := range users {
		course, _ := model.GetCourse(item.CourseID)
		res.Data.CourseList = append(res.Data.CourseList, &types.TCourse{
			CourseID:  fmt.Sprint(item.CourseID),
			Name:      course.Subject,
			TeacherID: fmt.Sprint(item.TeacherID),
		})
	}
	return res
}
