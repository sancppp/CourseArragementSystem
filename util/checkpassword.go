package util

func CheckPassword(s string) bool {
	//同时包含大小写数字
	var a, b, c bool
	for _, item := range s {
		if item >= 'A' && item <= 'Z' {
			a = true
		}
		if item >= 'a' && item <= 'z' {
			b = true
		}
		if item >= '0' && item <= '9' {
			c = true
		}
	}
	return a && b && c
}
