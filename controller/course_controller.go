package controller

import (
	"CAS/response"
	"CAS/types"
	"CAS/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	BaseController
}

func (CourseController) CreateCourse(c *gin.Context) {
	// if res := util.CheckUserType(c, types.Admin); res == -1 {
	// 	c.JSON(http.StatusOK, response.CreateCourseResponse{
	// 		Code: types.LoginRequired,
	// 	})
	// 	return
	// } else if res == 0 {
	// 	// 非管理员
	// 	c.JSON(http.StatusOK, response.CreateCourseResponse{
	// 		Code: types.PermDenied,
	// 	})
	// 	return
	// }

	serv := &response.CreateCourseRequest{}

	if err := c.ShouldBind(serv); err == nil {
		// fmt.Printf("serv: %v\n", serv)
		res := serv.Createcourse()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.CreateCourseResponse{
			Code: types.ParamInvalid,
		})
	}
}
func (CourseController) GetCourse(c *gin.Context) {
	if res := util.CheckUserType(c, types.Admin); res == -1 {
		c.JSON(http.StatusOK, response.GetCourseResponse{
			Code: types.LoginRequired,
		})
		return
	}
	// else if res == 0 {
	// 	//非管理员 ????
	// 	// c.JSON(http.StatusOK, response.GetCourseResponse{
	// 	// 	Code: types.PermDenied,
	// 	// })
	// 	// return
	// }

	serv := &response.GetCourseRequest{}
	if err := c.ShouldBind(serv); err == nil {
		res := serv.Getcourse()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.GetCourseResponse{
			Code: types.ParamInvalid,
		})
	}
}

func (CourseController) Hungary(c *gin.Context) {
	serv := &response.ScheduleCourseRequest{}
	if err := c.ShouldBind(serv); err == nil {
		res := serv.Hungary()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.ScheduleCourseResponse{
			Code: types.ParamInvalid,
		})
	}
}
