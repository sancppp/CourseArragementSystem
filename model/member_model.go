package model

import (
	"CAS/db/mysql"
	"CAS/types"
)

type Member struct {
	ID       uint
	Username string `gorm:"unique"`
	Nickname string
	Password string
	Type     types.UserType
	Deleted  bool `gorm:"default=false"`
}

func (Member) TableName() string {
	return "member"
}

//通过主键从mysql中查到一条记录
func GetUser(ID interface{}) (Member, error) {

	var user Member
	result := mysql.MysqlDB.GetConn().First(&user, ID)
	return user, result.Error
}
func InsertUser(member *Member) error {
	result := mysql.MysqlDB.GetConn().Create(member)
	return result.Error
}
