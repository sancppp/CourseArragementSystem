package model

import (
	"CAS/db/mysql"
	"crypto/md5"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

//gorm绑定模型，同时初始化Admin
func Default() {
	migration()
	// 系统内置管理员账号
	// 账号名：JudgeAdmin 密码：JudgePassword2022
	//初始化管理员
	if err := mysql.MysqlDB.GetConn().Where("username = 'JudgeAdmin'").First(&Member{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		mysql.MysqlDB.GetConn().Create(&Member{
			Username: "JudgeAdmin",
			Nickname: "Admin",
			Password: fmt.Sprint(md5.Sum([]byte("JudgePassword2022"))),
			Type:     1,
		})
	}
}

func migration() {
	mysql.MysqlDB.GetConn().AutoMigrate(&Member{})
	mysql.MysqlDB.GetConn().AutoMigrate(&Course{})
	mysql.MysqlDB.GetConn().AutoMigrate(&Course2Student{})
	mysql.MysqlDB.GetConn().AutoMigrate(&Course2Teacher{})
}
