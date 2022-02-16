package response

import (
	"CAS/db/mysql"
	"CAS/db/redis"
	"CAS/model"
	"CAS/types"
	"CAS/util"
	"crypto/md5"
	"fmt"
)

type CreateMemberRequest struct {
	Nickname string         `binding:"required,min=4,max=20"` // required，不小于 4 位 不超过 20 位
	Username string         `binding:"required,min=8,max=20"` // required，只支持大小写，长度不小于 8 位 不超过 20 位
	Password string         `binding:"required,min=8,max=20"` // required，同时包括大小写、数字，长度不少于 8 位 不超过 20 位
	UserType types.UserType `binding:"required"`              // required, 枚举值
}

type CreateMemberResponse struct {
	Code types.ErrNo
	Data struct {
		UserID string
	}
}

func (serv *CreateMemberRequest) Create() (res CreateMemberResponse) {
	//检查参数合法性
	if serv.UserType != 1 && serv.UserType != 2 && serv.UserType != 3 {
		res.Code = types.ParamInvalid
		return res
	}
	if !util.Checkstring(serv.Username, 8, 20) || !util.Checkstring(serv.Nickname, 4, 20) || !util.Checkstring(serv.Password, 8, 20) {
		res.Code = types.ParamInvalid
		return res
	}
	//username nickname 只包含大小写字母
	//Q: username需要验证是否只包含大小写字母吗
	//A: 是的
	//Q: Nickname 汉字的位数检查不正确,一个中文占三位？
	//A: Nickname只有大小写字母
	if !util.CheckWords(serv.Nickname) {
		res.Code = types.ParamInvalid
		return res
	}
	if !util.CheckWords(serv.Username) {
		res.Code = types.ParamInvalid
		return res
	}
	if err := util.CheckPassword(serv.Password); !err {
		res.Code = types.ParamInvalid
		return res
	}

	//判断userid是否存在
	if err := mysql.MysqlDB.GetConn().Where("username = ?", serv.Username).First(&model.Member{}).Error; err == nil {
		res.Code = types.UserHasExisted
		return
	}

	//可以向数据库写入数据了
	//mysql
	member := model.Member{
		Nickname: serv.Nickname,
		Password: fmt.Sprint(md5.Sum([]byte(serv.Password))),
		Username: serv.Username,
		Type:     serv.UserType,
	}
	mysql.MysqlDB.GetConn().Create(&member)
	fmt.Printf("\"ok\": %v\n", "ok")
	//redis
	if member.Type == types.Student {
		redis.Client().SAdd(redis.Ctx, fmt.Sprintf("studentcourse%d", member.ID), "") // 为每位同学默认加入一个空的课程。
	}

	res.Code = types.OK
	res.Data.UserID = fmt.Sprint(member.ID)
	return res
}
