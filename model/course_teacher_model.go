package model

import "gorm.io/gorm"

type Course2Teacher struct {
	gorm.Model
	Course    *Course
	CourseID  uint
	Teacher   *Member
	TeacherID uint
}

func (Course2Teacher) TableName() string {
	return "course_teacher"
}
