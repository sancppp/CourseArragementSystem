package response

import (
	"CAS/db/mysql"
	"CAS/model"
	"CAS/types"
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string
	Password string
}

// 登录成功后需要 Set-Cookie("camp-session", ${value})
// 密码错误返回密码错误状态码

type LoginResponse struct {
	Code types.ErrNo
	Data struct {
		UserID string
	}
}

func (service *LoginRequest) setSession(c *gin.Context, user *model.Member) {
	//登录成功，设置session
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	s.Set("userid", user.ID)
	s.Set("username", user.Username)
	s.Set("nickname", user.Nickname)
	s.Set("usertype", int(user.Type))
	s.Save()
}

func (serv *LoginRequest) Login(c *gin.Context) (res LoginResponse) {
	user := &model.Member{}
	mysql.MysqlDB.GetConn().Where("username = ?", serv.Username).First(&user)
	if user.Username == serv.Username && strings.Compare(user.Password, fmt.Sprint(md5.Sum([]byte(serv.Password)))) == 0 {
		//登录成功
		res.Code = types.OK
		res.Data.UserID = fmt.Sprint(user.ID)
		//设置session
		serv.setSession(c, user)

	} else {
		//一律返回密码错误
		res.Code = types.WrongPassword
		return res
	}
	return res
}
func (serv *LoginRequest) Loginbackup(c *gin.Context) (res LoginResponse) {
	user := &model.Member{}
	if err := mysql.MysqlDB.GetConn().Where("username = ?", serv.Username).First(&user).Error; err != nil {
		//用户不存在
		res.Code = types.UserNotExisted
		return res
	}
	if strings.Compare(user.Password, fmt.Sprint(md5.Sum([]byte(serv.Password)))) != 0 {
		//密码错误
		res.Code = types.WrongPassword
		return res
	}
	if user.Deleted {
		//已删除用户
		res.Code = types.UserHasDeleted
		return res
	}
	res.Code = types.OK
	res.Data.UserID = fmt.Sprint(user.ID)
	//设置session
	serv.setSession(c, user)
	return res
}
