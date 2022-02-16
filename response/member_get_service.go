package response

import (
	"CAS/model"
	"CAS/types"
)

// 获取成员信息

type GetMemberRequest struct {
	UserID string
}

// 如果用户已删除请返回已删除状态码，不存在请返回不存在状态码

type GetMemberResponse struct {
	Code types.ErrNo
	Data types.TMember
}

func (serv *GetMemberRequest) Getmember() (res GetMemberResponse) {
	if member, err := model.GetUser(serv.UserID); err == nil {
		res.Code = types.OK
		if member.Deleted {
			res.Code = types.UserHasDeleted
		} else {
			res.Code = types.OK
		}
		res.Data.Nickname = member.Nickname
		res.Data.UserID = serv.UserID
		res.Data.Username = member.Username
		res.Data.UserType = member.Type
	} else {
		res.Code = types.UserNotExisted
	}
	return res
}
