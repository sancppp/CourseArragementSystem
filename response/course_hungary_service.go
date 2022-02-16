package response

import (
	"CAS/types"
)

// 排课求解器，使老师绑定课程的最优解， 老师有且只能绑定一个课程
// Method： Post
type ScheduleCourseRequest struct {
	TeacherCourseRelationShip map[string][]string // key 为 teacherID , val 为老师期望绑定的课程 courseID 数组
	//如何传参还没弄清楚，写了test todo
}

type ScheduleCourseResponse struct {
	Code types.ErrNo
	Data map[string]string // key 为 teacherID , val 为老师最终绑定的课程 courseID
	//teacher -> course
}

var st map[string]bool
var res map[string]string
var rres map[string]string

func find(serv *ScheduleCourseRequest, x string) bool {
	tmp := serv.TeacherCourseRelationShip[x]
	for _, v := range tmp {
		if !st[v] {
			st[v] = true
			if xx, err := res[v]; !err || find(serv, xx) {
				res[v] = x
				rres[x] = v
				return true
			}
		}
	}
	return false
}

func (serv *ScheduleCourseRequest) Hungary() ScheduleCourseResponse {
	st = make(map[string]bool)
	res = make(map[string]string)
	rres = make(map[string]string)
	for _, values := range serv.TeacherCourseRelationShip {
		for _, v := range values {
			st[v] = false
		}
	}
	for key := range serv.TeacherCourseRelationShip {
		//fmt.Printf("key: %v\n", key)
		for s := range st {
			st[s] = false
		}
		find(serv, key)
		//fmt.Printf("rres: %v\n", rres)
	}
	resp := ScheduleCourseResponse{
		Code: types.OK,
		Data: rres,
	}
	return resp
}
