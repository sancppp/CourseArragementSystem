package controller

import (
	"CAS/response"
	"CAS/types"
	"CAS/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
	BaseController
}

func (MemberController) CreateMember(c *gin.Context) {
	if res := util.CheckUserType(c, types.Admin); res == -1 {
		//未登录
		c.JSON(http.StatusOK, response.CreateMemberResponse{
			Code: types.LoginRequired,
		})
		return
	} else if res == 0 {
		//非管理员
		c.JSON(http.StatusOK, response.CreateMemberResponse{
			Code: types.PermDenied,
		})
		return
	}

	serv := &response.CreateMemberRequest{}
	if err := c.ShouldBind(serv); err == nil {
		res := serv.Create()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.CreateMemberResponse{
			Code: types.ParamInvalid,
		})
	}
}

func (MemberController) GetMember(c *gin.Context) {
	if res := util.CheckUserType(c, types.Admin); res == -1 {
		c.JSON(http.StatusOK, response.GetMemberResponse{
			Code: types.LoginRequired,
		})
		return
	}
	// else if res == 0 {
	// 	//非管理员 ????
	// 	// c.JSON(http.StatusOK, response.GetMemberResponse{
	// 	// 	Code: types.PermDenied,
	// 	// })
	// 	// return
	// }

	serv := &response.GetMemberRequest{}
	if err := c.ShouldBind(serv); err == nil {
		res := serv.Getmember()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.GetMemberResponse{
			Code: types.ParamInvalid,
		})
	}
}
func (MemberController) GetMemberList(c *gin.Context) {
	if res := util.CheckUserType(c, types.Admin); res == -1 {
		c.JSON(http.StatusOK, response.GetMemberListResponse{
			Code: types.LoginRequired,
		})
		return
	}
	// else if res == 0 {
	// 	//非管理员 ????
	// 	// c.JSON(http.StatusOK, response.GetMemberListResponse{
	// 	// 	Code: types.PermDenied,
	// 	// })
	// 	// return
	// }

	serv := &response.GetMemberListRequest{}
	if err := c.ShouldBind(serv); err == nil {
		res := serv.Getmemberlist()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.GetMemberListResponse{
			Code: types.ParamInvalid,
		})
	}
}
func (MemberController) UpdateMember(c *gin.Context) {
	//修改nickname的权限检查？
	if res := util.CheckUserType(c, types.Admin); res == -1 {
		c.JSON(http.StatusOK, response.UpdateMemberResponse{
			Code: types.LoginRequired,
		})
		return
	}
	// else if res == 0 {
	// 	//非管理员 ????
	// 	// c.JSON(http.StatusOK, response.UpdateMemberResponse{
	// 	// 	Code: types.PermDenied,
	// 	// })
	// 	// return
	// }
	serv := &response.UpdateMemberRequest{}
	if err := c.ShouldBind(serv); err == nil {
		res := serv.Updatemember()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.UpdateMemberResponse{
			Code: types.ParamInvalid,
		})
	}
}
func (MemberController) DeleteMember(c *gin.Context) {
	if res := util.CheckUserType(c, types.Admin); res == -1 {
		c.JSON(http.StatusOK, response.DeleteMemberResponse{
			Code: types.LoginRequired,
		})
		return
	} else if res == 0 {
		// 非管理员
		c.JSON(http.StatusOK, response.DeleteMemberResponse{
			Code: types.PermDenied,
		})
		return
	}

	serv := &response.DeleteMemberRequest{}
	if err := c.ShouldBind(serv); err == nil {
		res := serv.Deletemember()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, response.DeleteMemberResponse{
			Code: types.ParamInvalid,
		})
	}
}
