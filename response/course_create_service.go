package response

import (
	"CAS/db/mysql"
	"CAS/model"
	"CAS/types"
	"fmt"
)

// 创建课程
// Method: Post
type CreateCourseRequest struct {
	Name string
	Cap  int
}

type CreateCourseResponse struct {
	Code types.ErrNo
	Data struct {
		CourseID string
	}
}

func (serv *CreateCourseRequest) Createcourse() (res CreateCourseResponse) {
	course := model.Course{
		Subject:   serv.Name,
		Capacity:  serv.Cap,
		RemainCap: serv.Cap,
	}
	if err := mysql.MysqlDB.GetConn().Where("subject = ?", course.Subject).First(&model.Course{}).Error; err == nil {
		//已经有该课程名
		res.Code = types.UnknownError
		return res
	}

	mysql.MysqlDB.GetConn().Create(&course)

	res.Code = types.OK
	res.Data.CourseID = fmt.Sprint(course.ID)
	return res
}
