package response

import (
	"CAS/db/mysql"
	"CAS/model"
	"CAS/types"
	"fmt"
)

// 批量获取成员信息

type GetMemberListRequest struct {
	Offset int
	Limit  int
}

type GetMemberListResponse struct {
	Code types.ErrNo
	Data struct {
		MemberList []types.TMember
	}
}

func (serv *GetMemberListRequest) Getmemberlist() (res GetMemberListResponse) {
	var users []model.Member
	mysql.MysqlDB.GetConn().Where("deleted = ?", 0).Limit(serv.Limit).Offset(serv.Offset).Find(&users)
	//fmt.Printf("users: %v\n", users)
	res.Code = types.OK
	for _, item := range users {
		res.Data.MemberList = append(res.Data.MemberList, types.TMember{
			UserID:   fmt.Sprint(item.ID),
			Nickname: item.Nickname,
			Username: item.Username,
			UserType: item.Type,
		})
	}
	return res
}
