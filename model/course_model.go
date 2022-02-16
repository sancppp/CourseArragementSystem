package model

import (
	"CAS/db/mysql"

	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	// ID        uint
	Subject   string `gorm:"unique"`
	Capacity  int
	RemainCap int
}

func (Course) TableName() string {
	return "course"
}

//通过主键找到一条记录
func GetCourse(ID interface{}) (Course, error) {
	var cour Course
	result := mysql.MysqlDB.GetConn().First(&cour, ID)
	return cour, result.Error
}
