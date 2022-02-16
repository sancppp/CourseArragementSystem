package controller

import (
	"CAS/response"
	"CAS/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	BaseController
}

func (StudentController) BookCourse(c *gin.Context) {
	// if res := util.CheckUserType(c, types.Student); res == -1 {
	// 	c.JSON(http.StatusOK, response.BookCourseResponse{
	// 		Code: types.LoginRequired,
	// 	})
	// 	return
	// } else if res == 0 {
	// 	//非学生
	// 	c.JSON(http.StatusOK, response.BookCourseResponse{
	// 		Code: types.PermDenied,
	// 	})
	// 	return
	// }

	serv := &response.BookCourseRequest{}
	if err := c.ShouldBind(serv); err == nil {
		res := serv.BookCourse()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BookCourseResponse{
			Code: types.ParamInvalid,
		})
	}
}

func (StudentController) GetCourse(c *gin.Context) {
	serv := &response.GetStudentCourseRequest{}
	if err := c.ShouldBind(serv); err == nil {
		res := serv.GetCourse()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.GetStudentCourseResponse{
			Code: types.ParamInvalid,
		})
	}
}
