package main

import "fmt"

/**
	反转字符串的元音字符 LeetCode No345

	标签：双指针
 */
func main() {
	vowelMap := map[string]string{"a": "a", "e": "e", "i": "i", "o": "o", "u": "u"}

	testString := "leeyoung"
	testStringRune := []rune(testString)

	i := 0
	j := len(testString) - 1
	for i < j {
		characterHead := string(testStringRune[i])
		characterTail := string(testStringRune[j])

		if vowelMap[characterHead] == "" {
			if vowelMap[characterTail] == "" {
				i ++
				j --
			} else {
				i --
			}
		} else {
			if vowelMap[characterTail] == "" {
				j --
			} else {
				characterTemp := testStringRune[i]
				testStringRune[i] = testStringRune[j]
				testStringRune[j] = characterTemp
				i ++
				j --
			}
		}
	}
	fmt.Println(testString)
	result := string(testStringRune)
	fmt.Println(result)
}
