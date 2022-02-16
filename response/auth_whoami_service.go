package response

import (
	"CAS/types"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// WhoAmI 接口，用来测试是否登录成功，只有此接口需要带上 Cookie

type WhoAmIRequest struct {
}

// 用户未登录请返回用户未登录状态码

type WhoAmIResponse struct {
	Code types.ErrNo
	Data types.TMember
}

func (WhoAmIRequest) WhoAmI(c *gin.Context) (res WhoAmIResponse) {
	session := sessions.Default(c)
	tmp := session.Get("userid")
	if tmp != nil {
		//当前用户已经登录
		res.Code = types.OK
		res.Data.UserID = fmt.Sprint(tmp)
		res.Data.Nickname = fmt.Sprint(session.Get("nickname"))
		res.Data.Username = fmt.Sprint(session.Get("username"))
		res.Data.UserType = types.UserType(session.Get("usertype").(int))
	} else {
		//未登录
		res.Code = types.LoginRequired
	}
	return res
}
