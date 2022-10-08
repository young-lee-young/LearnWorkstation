package main

/**
反转字符串 LeetCode No344

标签：对撞指针
*/
func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	solution(s)
}

func solution(s []byte) {
	left := 0
	right := len(s) - 1

	for left < right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
}
