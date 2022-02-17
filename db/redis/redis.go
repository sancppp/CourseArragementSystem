package redis

import (
	"CAS/db/mysql"
	"CAS/model"
	"CAS/types"
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type connect struct {
	client *redis.Client
}

var once = sync.Once{}

var _connect *connect

var Ctx = context.Background()

func Default() {

	cxt, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	conf := &redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "bytedancecamp",
		DB:       0,
	}

	c := redis.NewClient(conf)

	if re, err := c.Ping(cxt).Result(); err != nil {
		panic(err)
	} else {
		fmt.Printf("re: %v\n", re)
	}

	c.FlushDB(Ctx)

	var students []model.Member // 为每一位同学开一个course集合 key = studentcourse%d，并进行课程同步。
	mysql.MysqlDB.GetConn().Where("deleted = ? and type = ?", 0, types.Student).Find(&students)
	for _, student := range students {
		c.SAdd(Ctx, fmt.Sprintf("studentcourse%d", student.ID), "") // 为每位同学的set中默认加入一个空的课程。
	}

	var studentcourse []model.Course2Student
	mysql.MysqlDB.GetConn().Find(&studentcourse)
	for _, it := range studentcourse {
		c.SAdd(Ctx, fmt.Sprintf("studentcourse%d", it.StudentID), fmt.Sprintf("%d", it.CourseID))
	}
	// //初始化课程，只有绑定了老师的课程才是有效课程
	// var coursesteacher []model.Course2Teacher
	// mysql.MysqlDB.GetConn().Find(&coursesteacher)
	// for _, course := range coursesteacher {
	// 	tmp, _ := model.GetCourse(course.CourseID)
	// 	c.Set(Ctx, fmt.Sprintf("coursename%d", course.CourseID), tmp.Subject, redis.KeepTTL)
	// 	c.Set(Ctx, fmt.Sprintf("courseteacher%d", course.CourseID), course.TeacherID, redis.KeepTTL)
	// 	c.Set(Ctx, fmt.Sprintf("course%d", course.CourseID), tmp.RemainCap, redis.KeepTTL)
	// }

	//所有课程都有效，能被抢
	var courses []model.Course
	var c2t model.Course2Teacher
	mysql.MysqlDB.GetConn().Find(&courses)
	for _, course := range courses {
		c.Set(Ctx, fmt.Sprintf("coursename%d", course.ID), course.Subject, redis.KeepTTL)
		c.Set(Ctx, fmt.Sprintf("course%d", course.ID), course.RemainCap, redis.KeepTTL)
		if err := mysql.MysqlDB.GetConn().Where("course_id = ?", course.ID).First(&c2t).Error; err == nil {
			c.Set(Ctx, fmt.Sprintf("courseteacher%d", course.ID), c2t.TeacherID, redis.KeepTTL)

		}
	}
	go BookCourse()
	_connect = &connect{
		client: c,
	}
}

func Client() *redis.Client {

	if _connect == nil {

		once.Do(func() {

			Default()
		})

	}

	return _connect.client

}

type Pair struct {
	StudentID string
	CourseID  string
}

var BookCourseInfo = make(chan Pair, 10000)

//同步到mysql中
func BookCourse() {
	for {
		info := <-BookCourseInfo
		// studencourse表中加入该课程
		studentid, _ := strconv.Atoi(info.StudentID)
		courseid, _ := strconv.Atoi(info.CourseID)
		studentcourse := model.Course2Student{
			StudentID: uint(studentid),
			CourseID:  uint(courseid),
		}
		mysql.MysqlDB.GetConn().Create(&studentcourse)
		// course 表中将该课程remaincap - 1
		mysql.MysqlDB.GetConn().Model(&model.Course{}).Where("id = ?", info.CourseID).Update("remain_cap", gorm.Expr("remain_cap - ?", 1))
	}
}
