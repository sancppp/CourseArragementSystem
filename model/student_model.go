package model

//sample unused
type Student struct {
	Id         int    `gorm:"column:student_id" json:"student_id"` //
	Name       string `gorm:"size:30" json:"name"`                 // string默认长度为255, 使用这种tag重设
	Department string `gorm:"size:45" json:"department"`           // string默认长度为255, 使用这种tag重设
	Major      string `gorm:"size:45" json:"major"`                // string默认长度为255, 使用这种tag重设
}

func (Student) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "student"
}
