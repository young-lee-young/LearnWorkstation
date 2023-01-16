package main

import "fmt"

/**
LeetCode No283 将数组0移动到末尾，其他保持相对顺序

时间复杂度：O(n)
空间复杂度：O(1)
*/
func main() {
	numArray := [...]int{4, 0, 3, 1, 0, 9, 5}
	j := 0
	for i := 0; i < len(numArray); i++ {
		if numArray[i] != 0 {
			numArray[j] = numArray[i]
			if i != j {
				numArray[i] = 0
			}
			j++
		}
	}
	fmt.Println(numArray)
}
