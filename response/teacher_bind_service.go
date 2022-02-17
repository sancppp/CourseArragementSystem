package response

import (
	"CAS/db/mysql"
	"CAS/db/redis"
	"CAS/model"
	"CAS/types"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// 老师绑定课程
// Method： Post
// 注：这里的 teacherID 不需要做已落库校验
// 一个老师可以绑定多个课程 , 不过，一个课程只能绑定在一个老师下面
type BindCourseRequest struct {
	CourseID  string
	TeacherID string
}

type BindCourseResponse struct {
	Code types.ErrNo
}

func (serv *BindCourseRequest) Bind() (res BindCourseResponse) {
	cour, err := model.GetCourse(serv.CourseID)
	if err != nil {
		//没这门课
		res.Code = types.CourseNotExisted
		return res
	}
	teacher, err := model.GetUser(serv.TeacherID)
	if err != nil {
		//不存在这个人
		res.Code = types.UserNotExisted
		return res
	}

	if teacher.Type != types.Teacher {
		//不是老师
		res.Code = types.UnknownError
		return res
	}

	temp := &model.Course2Teacher{
		Course:  &cour,
		Teacher: &teacher,
	}
	if err := mysql.MysqlDB.GetConn().Where("course_id = ?", cour.ID).First(&model.Course2Teacher{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		//这门课还没被绑定
		mysql.MysqlDB.GetConn().Create(&temp)

		//写入redis
		redis.Client().Set(redis.Ctx, fmt.Sprintf("courseteacher%d", temp.CourseID), temp.TeacherID, -1)
		res.Code = types.OK
		return res
	}
	res.Code = types.CourseHasBound
	return res
}
