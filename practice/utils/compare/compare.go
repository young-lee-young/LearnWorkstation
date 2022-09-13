package compare

/**
	返回小值
 */
func Min(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}

/**
	返回大值
 */
func Max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}
