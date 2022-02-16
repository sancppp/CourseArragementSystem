package response

import (
	"CAS/db/mysql"
	"CAS/model"
	"CAS/types"
	"CAS/util"
	"errors"

	"gorm.io/gorm"
)

// 更新成员信息

type UpdateMemberRequest struct {
	UserID   string
	Nickname string `binding:"required,min=4,max=20"`
}

type UpdateMemberResponse struct {
	Code types.ErrNo
}

func (serv *UpdateMemberRequest) Updatemember() (res UpdateMemberResponse) {
	if !util.Checkstring(serv.Nickname, 4, 20) {
		//新的Nickname长度不合法
		res.Code = types.ParamInvalid
		return res
	}

	member, err := model.GetUser(serv.UserID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//用户不存在
		res.Code = types.UserNotExisted
		return res
	}

	if member.Deleted {
		//用户已经被删除
		res.Code = types.UserHasDeleted
		return res
	}
	if err := mysql.MysqlDB.GetConn().Model(&model.Member{}).Where("id = ?", serv.UserID).Update("nickname", serv.Nickname).Error; err == nil {
		res.Code = types.OK
		return res
	}
	res.Code = types.UnknownError
	return res
}
