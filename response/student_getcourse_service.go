package response

import (
	"CAS/db/redis"
	"CAS/types"
	"fmt"
)

type GetStudentCourseRequest struct {
	StudentID string
}

type GetStudentCourseResponse struct {
	Code types.ErrNo
	Data struct {
		CourseList []types.TCourse
	}
}

//从直接redis中查
func (serv *GetStudentCourseRequest) GetCourse() (res GetStudentCourseResponse) {
	if redis.Client().TTL(redis.Ctx, fmt.Sprintf("studentcourse%s", serv.StudentID)).Val().String()[:2] == "-2" {
		res.Code = types.StudentNotExisted
	} else {
		courses := redis.Client().SMembers(redis.Ctx, fmt.Sprintf("studentcourse%s", serv.StudentID)).Val()
		if len(courses) == 1 {
			res.Code = types.StudentHasNoCourse
		} else {
			for _, course := range courses {
				if course == "" {
					continue
				}
				res.Data.CourseList = append(res.Data.CourseList, types.TCourse{
					CourseID:  course,
					Name:      redis.Client().Get(redis.Ctx, fmt.Sprintf("coursename%s", course)).Val(),
					TeacherID: redis.Client().Get(redis.Ctx, fmt.Sprintf("courseteacher%s", course)).Val(),
				})
			}
		}
	}
	return res
}
