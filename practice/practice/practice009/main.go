package main

import "fmt"

/**
	括号生成 LeetCode No22

	标签：递归
 */
func main() {
	result := solution(3)
	fmt.Println(result)
}

func solution(count int) []string {
	result := make([]string, 0)
	if count == 0 {
		return result
	}
	if count == 1 {
		result = append(result, "()")
		return result
	}
	bracketsSlice := solution(count - 1)
	fmt.Println(bracketsSlice)
	hashMap := make(map[string]string)
	for _, value := range bracketsSlice {
		// 循环字符串
		for i := 0; i < len(value); i ++ {
			newValue := value[:i]  + "()" + value[i:]
			if _, ok := hashMap[newValue]; !ok {
				result = append(result, newValue)
				hashMap[newValue] = newValue
			}
		}
	}
	fmt.Println(result)
	return result
}
