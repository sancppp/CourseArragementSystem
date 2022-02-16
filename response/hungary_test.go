package response

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tmp := ScheduleCourseRequest{}
	input := map[string][]string{
		"a": {"1", "3"},
		"b": {"1"},
		"c": {"1", "2"},
		"d": {"3", "4"},
	}
	tmp.TeacherCourseRelationShip = input
	res := tmp.Hungary()
	fmt.Printf("res: %v\n", res)
}
