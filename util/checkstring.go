package util

//判断字符串长度
func Checkstring(x string, _min int, _max int) bool {
	return len(x) >= _min && len(x) <= _max
}
