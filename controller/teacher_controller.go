package controller

import (
	"CAS/response"
	"CAS/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeacherController struct {
	BaseController
}

func (TeacherController) Bind(c *gin.Context) {
	serv := &response.BindCourseRequest{}
	if err := c.ShouldBind(serv); err == nil {
		res := serv.Bind()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.BindCourseResponse{
			Code: types.ParamInvalid,
		})
	}
}

func (TeacherController) Unbind(c *gin.Context) {
	serv := &response.UnbindCourseRequest{}
	if err := c.ShouldBind(serv); err == nil {
		res := serv.Unbind()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.UnbindCourseResponse{
			Code: types.ParamInvalid,
		})
	}
}
func (TeacherController) GetCourse(c *gin.Context) {
	serv := &response.GetTeacherCourseRequest{}
	if err := c.ShouldBind(serv); err == nil {
		res := serv.GetCourse()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.GetTeacherCourseResponse{
			Code: types.ParamInvalid,
		})
	}
}
