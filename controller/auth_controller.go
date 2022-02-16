package controller

import (
	"CAS/response"
	"CAS/types"
	"CAS/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	BaseController
}

func (AuthController) UserLogin(c *gin.Context) {
	serv := &response.LoginRequest{}
	if err := c.ShouldBind(serv); err == nil && util.Checkstring(serv.Username, 8, 20) {
		res := serv.Login(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.LoginResponse{
			Code: types.ParamInvalid,
		})
	}
}

func (AuthController) UserLogout(c *gin.Context) {
	serv := &response.LogoutRequest{}
	if err := c.ShouldBind(serv); err == nil {
		res := serv.Logout(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.LogoutResponse{
			Code: types.ParamInvalid,
		})
	}
}

func (AuthController) Whoami(c *gin.Context) {
	serv := &response.WhoAmIRequest{}
	if err := c.ShouldBind(serv); err == nil {
		res := serv.WhoAmI(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.WhoAmIResponse{
			Code: types.ParamInvalid,
		})
	}
}
