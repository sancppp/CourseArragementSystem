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

// 老师解绑课程
// Method： Post
type UnbindCourseRequest struct {
	CourseID  string
	TeacherID string
}

type UnbindCourseResponse struct {
	Code types.ErrNo
}

func (serv *UnbindCourseRequest) Unbind() (res UnbindCourseResponse) {
	_, err := model.GetCourse(serv.CourseID)
	if err != nil {
		//没这门课
		res.Code = types.CourseNotExisted
		return res
	}
	temp := &model.Course2Teacher{}
	if err := mysql.MysqlDB.GetConn().Where("course_id = ? and teacher_id = ?", serv.CourseID, serv.TeacherID).First(temp).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		//todo 找不到这条记录
		res.Code = types.CourseNotBind
		return res
	}
	//删除MySQL中的记录
	mysql.MysqlDB.GetConn().Delete(&temp)
	//删除redis中的记录
	redis.Client().Del(redis.Ctx, fmt.Sprintf("courseteacher%d", temp.CourseID))
	res.Code = types.OK
	return res
}
