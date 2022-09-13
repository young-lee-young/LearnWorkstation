package main

import "fmt"

/**
	可以删除一个字符，判断回文字符串 LeetCode No680

	标签：双指针
 */
func main() {
	//testString := "leeyoung"
	//testString := "leeyounggnuoyeel"
	//testString := "leeyounggnuoyel"
	//testString := "leeyounggnuoyl"
	testString := "leeyoungzgnuoyeel"
	//testString := "leeyoungzgnuoyel"

	i := 0
	j := len(testString) - 1
	for i < j {
		if testString[i] == testString[j] {
			i ++
			j --
			continue
		}

		stringDeleteLeft := testString[i + 1 : j + 1]
		stringDeleteRight := testString[i : j]
		leftResult := isPalindrome(stringDeleteLeft)
		rightResult := isPalindrome(stringDeleteRight)
		if leftResult || rightResult {
			fmt.Println(true)
			return
		} else {
			return
		}
	}
	fmt.Println(true)
}

func isPalindrome(testString string) bool {
	i := 0
	j := len(testString) - 1

	for i < j {
		if testString[i] == testString[j] {
			i ++
			j --
		} else {
			return false
		}
	}
	return true
}
