package response

import (
	"CAS/db/mysql"
	"CAS/db/redis"
	"CAS/model"
	"CAS/types"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// 删除成员信息
// 成员删除后，该成员不能够被登录且不应该不可见，ID 不可复用

type DeleteMemberRequest struct {
	UserID string
}

type DeleteMemberResponse struct {
	Code types.ErrNo
}

func (serv *DeleteMemberRequest) Deletemember() (res DeleteMemberResponse) {
	// fmt.Printf("serv: %v\n", serv)
	member, err := model.GetUser(serv.UserID)
	// fmt.Printf("err: %v\n", err)
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
	if err := mysql.MysqlDB.GetConn().Model(&model.Member{}).Where("id = ?", serv.UserID).Update("deleted", 1).Error; err == nil {
		if member.Type == types.Student {
			//从redis中删除该学生
			redis.Client().Del(redis.Ctx, fmt.Sprintf("studentcourse%d", member.ID))
		}
		res.Code = types.OK
		return res
	}
	res.Code = types.UnknownError
	return res
}
