package char

/**
十进制A为65，a为97
*/

/**
判断是否是字母
*/
func IsAlpha(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}

/**
判断是否为大写字符
*/
func IsCap(ch byte) bool {
	return ch >= 'A' && ch <= 'Z'
}

/**
判断是否为小写字符
*/
func IsLower(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}

/**
判断是否是数字
*/
func IsDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

/**
将大写字母转换为小写字母
*/
func CapToLower(ch byte) byte {
	return ch + 32
}

/**
将小写字母转换为大小字母
*/
func LowerToCap(ch byte) byte {
	return ch - 32
}
