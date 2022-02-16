package util

//判断字符串长度
func Checkstring(x string, _min int, _max int) bool {
	return len(x) >= _min && len(x) <= _max
}

//检查字符串是否只含大小写字母
func CheckWords(x string) bool {
	for _, v := range x {
		if !(('a' <= v && v <= 'z') || ('A' <= v && v <= 'Z')) {
			return false
		}
	}
	return true
}
