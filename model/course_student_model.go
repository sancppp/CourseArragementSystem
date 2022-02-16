package model

type Course2Student struct {
	ID        uint
	Course    *Course
	CourseID  uint
	Student   *Member
	StudentID uint
}

func (Course2Student) TableName() string {
	return "course_student"
}
