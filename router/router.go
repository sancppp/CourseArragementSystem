package router

import (
	"CAS/controller"
	"CAS/middleware"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")
	g.Use(middleware.LoggerToFile())
	g.Use(middleware.Session(os.Getenv("SESSION_SECRET")))

	//test
	g.Any("/auth/test", func(c *gin.Context) {
		session := sessions.Default(c)
		c.String(http.StatusOK, " %v %v %v %v", session.Get("userid"), session.Get("username"), session.Get("nickname"), session.Get("usertype"))
	})

	// 成员管理
	g.POST("/member/create", controller.MemberController{}.CreateMember)
	g.GET("/member", controller.MemberController{}.GetMember)
	g.GET("/member/list", controller.MemberController{}.GetMemberList)
	g.POST("/member/update", controller.MemberController{}.UpdateMember)
	g.POST("/member/delete", controller.MemberController{}.DeleteMember)

	// 登录

	g.POST("/auth/login", controller.AuthController{}.UserLogin)
	g.POST("/auth/logout", controller.AuthController{}.UserLogout)
	g.GET("/auth/whoami", controller.AuthController{}.Whoami)

	// 排课
	g.POST("/course/create", controller.CourseController{}.CreateCourse)
	g.GET("/course/get", controller.CourseController{}.GetCourse)

	g.POST("/teacher/bind_course", controller.TeacherController{}.Bind)
	g.POST("/teacher/unbind_course", controller.TeacherController{}.Unbind)
	g.GET("/teacher/get_course", controller.TeacherController{}.GetCourse)
	g.POST("/course/schedule", controller.CourseController{}.Hungary)

	// 抢课
	g.POST("/student/book_course", controller.StudentController{}.BookCourse)
	g.GET("/student/course", controller.StudentController{}.GetCourse)

}
