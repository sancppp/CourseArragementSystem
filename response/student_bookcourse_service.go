package response

import (
	"CAS/db/redis"
	"CAS/model"
	"CAS/types"
	"fmt"
)

type BookCourseRequest struct {
	StudentID string
	CourseID  string
}

// 课程已满返回 CourseNotAvailable

type BookCourseResponse struct {
	Code types.ErrNo
}

//todo：如果有时间的话，需要重写
func (serv *BookCourseRequest) BookCourse() (res BookCourseResponse) {
	//从mysql中判断user是否被删除
	if tmp, err := model.GetUser(serv.StudentID); err == nil {
		if tmp.Deleted {
			res.Code = types.UserHasDeleted
			return res
		}
	}
	if redis.Client().TTL(redis.Ctx, fmt.Sprintf("studentcourse%s", serv.StudentID)).Val().String()[:2] == "-2" { // 是否存在该学生
		res.Code = types.StudentNotExisted
	} else if redis.Client().TTL(redis.Ctx, fmt.Sprintf("course%s", serv.CourseID)).Val().String()[:2] == "-2" { // 是否存在该课程
		res.Code = types.CourseNotExisted
	} else if redis.Client().SIsMember(redis.Ctx, fmt.Sprintf("studentcourse%s", serv.StudentID), serv.CourseID).Val() { // 学生是否已经选择该课程
		res.Code = types.StudentHasCourse
	} else { // 课程是否有容量
		remain := redis.Client().Decr(redis.Ctx, fmt.Sprintf("course%s", serv.CourseID)).Val()
		if remain < 0 {
			redis.Client().Incr(redis.Ctx, fmt.Sprintf("course%s", serv.CourseID))
			res.Code = types.CourseNotAvailable
		} else {
			redis.Client().SAdd(redis.Ctx, fmt.Sprintf("studentcourse%s", serv.StudentID), serv.CourseID)
			redis.BookCourseInfo <- redis.Pair{
				StudentID: serv.StudentID,
				CourseID:  serv.CourseID,
			}
			res.Code = types.OK
		}
	}
	return res
}
