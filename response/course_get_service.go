package response

import (
	"CAS/db/redis"
	"CAS/types"
	"fmt"
)

// 获取课程
// Method: Get
type GetCourseRequest struct {
	CourseID string
}

type GetCourseResponse struct {
	Code types.ErrNo
	Data types.TCourse
}

//直接从redis中获取
func (serv *GetCourseRequest) Getcourse() (res GetCourseResponse) {
	if redis.Client().TTL(redis.Ctx, fmt.Sprintf("course%s", serv.CourseID)).Val().String()[:2] == "-2" { // 是否存在该课程
		res.Code = types.CourseNotExisted
		return res
	}
	res.Code = types.OK
	res.Data.CourseID = serv.CourseID
	res.Data.Name = redis.Client().Get(redis.Ctx, fmt.Sprintf("coursename%s", serv.CourseID)).Val()
	res.Data.TeacherID = redis.Client().Get(redis.Ctx, fmt.Sprintf("courseteacher%s", serv.CourseID)).Val()
	return res
}

// func (serv *GetCourseRequest) Getcourse() (res GetCourseResponse) {
// 	if course, err := model.GetCourse(serv.CourseID); err == nil {
// 		res.Code = types.OK
// 		res.Data.CourseID = fmt.Sprint(course.ID)
// 		res.Data.Name = course.Subject
// 		// res.Data.TeacherID?
// 		var temp model.Course2Teacher
// 		if err:=mysql.MysqlDB.GetConn().Where("course_id = ?",res.Data.CourseID).First(&temp).Error;err==nil{
// 			//已经被绑定
// 		}
// 		// types.CourseNotBind
// 	} else {
// 		res.Code = types.CourseNotExisted
// 	}
// 	return res
// }
