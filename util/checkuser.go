package util

import (
	"CAS/types"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//检查用户类型
func CheckUserType(c *gin.Context, target types.UserType) int {
	session := sessions.Default(c)
	tmp := session.Get("usertype")
	if tmp == nil {
		//没登录
		return -1
	}
	if types.UserType(tmp.(int)) != target {
		//不是
		return 0
	}
	return 1
}
