package response

import (
	"CAS/types"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 登出

type LogoutRequest struct{}

// 登出成功需要删除 Cookie

type LogoutResponse struct {
	Code types.ErrNo
}

func (service *LogoutRequest) Logout(c *gin.Context) (res LogoutResponse) {
	//todo:重复登出
	session := sessions.Default(c)
	if tmp := session.Get("userid"); tmp == nil {
		res.Code = types.LoginRequired
		return res
	}
	//清除sessions
	session.Clear()
	session.Save()
	res.Code = types.OK
	return res
}
