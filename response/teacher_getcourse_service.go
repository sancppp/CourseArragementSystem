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
	//Q: 在获取老师课程请求，老师没有课程和没有老师和没有对应状态码
	//A: 获取老师课程，没有课程，就返回空数组。 老师是否存在不需要再做校验

	// //先判断是否为未被删除的老师
	// teacher, err := model.GetUser(serv.TeacherID)
	// if err != nil {
	// 	//不存在这个人
	// 	res.Code = types.UserNotExisted
	// 	return res
	// }

	// if teacher.Type != types.Teacher {
	// 	//不是老师
	// 	res.Code = types.UnknownError
	// 	return res
	// }
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
